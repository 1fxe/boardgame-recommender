package main

import (
	"encoding/json"
	"github.com/1fxe/board-game-recommender-system/recommenders"
	"github.com/1fxe/board-game-recommender-system/shared"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := shared.Database{}
	// TODO Database connection
	db.Connect(os.Getenv("POSTGRES_CONNECTION_URL"))
	db.CreateBoardGameTable()
	db.CreateCharacteristicTable()
	db.Close()

	rand.Seed(time.Now().UnixNano())
	file, err := os.ReadFile("data/test_game_data.json")
	if err != nil {
		log.Panicln("Error reading file", err)
	}

	var boardGames []shared.BoardGame
	err = json.Unmarshal(file, &boardGames)
	if err != nil {
		log.Panicln("Error unmarshalling json", err)
	}

	// Give each Game a fake ID
	for i := range boardGames {
		boardGames[i].ID = i
	}

	var users []shared.User
	file, err = os.ReadFile("data/test_user_data.json")
	if err != nil {
		log.Panicln("Error reading file", err)
	}
	err = json.Unmarshal(file, &users)

	contentBasedRecommender := recommenders.ContentBasedRecommender{}
	contentBasedRecommender.PopulateBoardGames(boardGames)
	recommendations := contentBasedRecommender.RecommendBoardGames(boardGames[0], 0.4, shared.Characteristic{
		Categories: []shared.Data{
			{
				Name:        "Wargame",
				Description: "are games that depict military actions. \u003cem\u003eWargames\u003c",
			},
		},
	})

	// shared.PrettyPrintRecommendations(recommendations)

	collaborativeRecommender := recommenders.CollaborativeRecommender{}
	collaborativeRecommender.PopulateBoardGames(boardGames)
	collaborativeRecommender.PopulateUsers(users)
	theUser := rand.Intn(len(users))

	shared.PrintUserData(users[theUser], boardGames)

	recommendations = collaborativeRecommender.RecommendBoardGames(theUser)

	shared.PrettyPrintRecommendations(recommendations)
}
