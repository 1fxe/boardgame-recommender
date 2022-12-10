package shared

import "log"

type Characteristic struct {
	Categories []Data `json:"categories,omitempty"`
	Mechanisms []Data `json:"mechanisms,omitempty"`
}

type Data struct {
	Name        string `json:"dataName"`
	Description string `json:"dataDescription"`
}

type Range struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type BoardGame struct {
	ID             int
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	YearReleased   int            `json:"yearReleased"`
	NoPlayers      Range          `json:"noPlayers"`
	PlayTime       Range          `json:"playTime"`
	MinAge         int            `json:"age"`
	Characteristic Characteristic `json:"characteristic"`
}

type Recommendation struct {
	BoardGame BoardGame `json:"boardGame"`
	Score     float64   `json:"score"`
}

type User struct {
	ID      int
	Ratings map[int]int
	// TODO: add more fields
	// Favourite Characteristics
	// Favourite Board Games
	// Play Time
	// Age
	// Number of players?
}

func DataEquals(a, b Data) bool {
	if a.Name == b.Name && a.Description == b.Description {
		return true
	}

	return false
}

func PrintUserData(user User, boardGames []BoardGame) {
	log.Println("User ID:", user.ID)
	log.Println("User Ratings:")
	for key, value := range user.Ratings {
		log.Printf("%s : %d", boardGames[key].Name, value)
	}
	log.Println()
}

func PrettyPrintRecommendations(recommendations []Recommendation) {
	for _, recommendation := range recommendations {
		log.Println(recommendation.BoardGame.Name, recommendation.Score)
	}
}
