package shared

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	db *sql.DB
}

func (database *Database) Connect(connectionURL string) {
	db, err := sql.Open("postgres", connectionURL)
	if err != nil {
		log.Panicln("Error connecting to database: ", err)
	}

	database.db = db
}

func (database *Database) Close() {
	err := database.db.Close()
	if err != nil {
		log.Println("Error closing database connection: ", err)
	}
}

func (database *Database) CreateCharacteristicTable() {
	sqlStatement := `
		CREATE TABLE IF NOT EXISTS characteristics (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL
-- 			description TEXT NOT NULL
		)`
	_, err := database.db.Exec(sqlStatement)
	if err != nil {
		log.Println("Error creating characteristics table: ", err)
	}
}

func (database *Database) CreateBoardGameTable() {
	createBoardGameTableStatement := `
		CREATE TABLE IF NOT EXISTS board_games (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			year_released INT NOT NULL,
			min_players INT NOT NULL,
			max_players INT NOT NULL,
			min_play_time INT NOT NULL,
			max_play_time INT NOT NULL,
			min_age INT NOT NULL,
			categories INT[] NOT NULL,
			mechanisms INT[] NOT NULL
		)`
	_, err := database.db.Exec(createBoardGameTableStatement)
	if err != nil {
		log.Println("Error creating board_games table: ", err)
	}
}

func (database *Database) AddBoardGame(boardGame BoardGame) {
	var stmt, err = database.db.Prepare("INSERT INTO board_games (name, description, year_released, min_players, max_players, min_play_time, max_play_time, min_age, categories, mechanisms) " +
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Error preparing statement: ", err)
	}

	_, err = stmt.Exec(boardGame.Name, boardGame.Description, boardGame.YearReleased, boardGame.NoPlayers.Min,
		boardGame.NoPlayers.Max, boardGame.PlayTime.Min, boardGame.PlayTime.Max,
		boardGame.MinAge, boardGame.Characteristic.Categories, boardGame.Characteristic.Mechanisms,
	)
	if err != nil {
		log.Fatal("Error inserting board game: ", err)
	}

	log.Println("Board game inserted", boardGame.Name)
}

func (database *Database) GetBoardGame(id int) BoardGame {
	var boardGame BoardGame
	stmt, err := database.db.Prepare("SELECT * FROM album WHERE id = ?")
	if err != nil {
		log.Fatal("Error preparing statement: ", err)
	}

	row := stmt.QueryRow(id)

	err = row.Scan(&boardGame.ID, &boardGame.Name, &boardGame.Description,
		&boardGame.YearReleased, &boardGame.NoPlayers.Min, &boardGame.NoPlayers.Max,
		&boardGame.PlayTime.Min, &boardGame.PlayTime.Max, &boardGame.MinAge,
	)
	if err != nil {
		log.Fatalln("Error getting board game: ", err)
		return BoardGame{}
	}

	return boardGame
}
