package cotent

import (
	"encoding/json"
	"fmt"
	"github.com/1fxe/board-game-recommender-system/internal"
	"log"
	"math/rand"
	"os"
	"time"
)

type Recommendation struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type BoardGameRecommender struct {
	BoardGames    []internal.BoardGame
	BoardGamesMap map[string]internal.BoardGame
}

func (recommender *BoardGameRecommender) PopulateBoardGames(boardGames []internal.BoardGame) {
	recommender.BoardGames = boardGames
	recommender.BoardGamesMap = make(map[string]internal.BoardGame)
	for _, game := range boardGames {
		recommender.BoardGamesMap[game.Name] = game
	}
}

// RecommendItems Recommendation algorithm for board games using Characteristics
// Checks if categories or mechanisms match
func (recommender *BoardGameRecommender) RecommendItems(boardGame internal.BoardGame) []Recommendation {
	var recommendations []Recommendation
	giveGameToCompare := recommender.BoardGamesMap[boardGame.Name]

	for _, game := range recommender.BoardGames {
		if game.Name == boardGame.Name {
			continue
		}

		commonCharacteristics := 0
		gameToCompare := recommender.BoardGamesMap[game.Name]
		for _, characteristic := range giveGameToCompare.Characteristic.Categories {
			for _, otherCharacteristic := range gameToCompare.Characteristic.Categories {
				if characteristic == otherCharacteristic {
					commonCharacteristics++
				}
			}
		}

		for _, characteristic := range giveGameToCompare.Characteristic.Mechanisms {
			for _, otherCharacteristic := range gameToCompare.Characteristic.Mechanisms {
				if characteristic == otherCharacteristic {
					commonCharacteristics++
				}
			}
		}

		if commonCharacteristics >= 3 {
			recommendations = append(recommendations, Recommendation{Name: game.Name, Score: commonCharacteristics})
		}
	}
	return recommendations
}

func main() {
	file, err := os.ReadFile("./data/test_data.json")
	if err != nil {
		log.Panicln("Error reading file", err)
	}

	var boardGames []internal.BoardGame
	err = json.Unmarshal(file, &boardGames)
	if err != nil {
		log.Panicln("Error unmarshalling board games", err)
	}

	recommender := &BoardGameRecommender{}
	recommender.PopulateBoardGames(boardGames)

	rand.Seed(time.Now().Unix())
	boardGame := boardGames[rand.Intn(len(boardGames))]

	recommendations := recommender.RecommendItems(boardGame)

	fmt.Println("Recommendations for", boardGame.Name)
	fmt.Println("Name | Score")
	for _, recommendation := range recommendations {
		fmt.Printf("%s | %d\n", recommendation.Name, recommendation.Score)
	}
}
