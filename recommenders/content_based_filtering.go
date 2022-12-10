package recommenders

import (
	"github.com/1fxe/board-game-recommender-system/shared"
)

type ContentBasedRecommender struct {
	BoardGames    []shared.BoardGame
	BoardGamesMap map[int]shared.BoardGame
}

// PopulateBoardGames Populates the board games map
func (recommender *ContentBasedRecommender) PopulateBoardGames(boardGames []shared.BoardGame) {
	recommender.BoardGames = boardGames
	recommender.BoardGamesMap = make(map[int]shared.BoardGame)
	for i, game := range boardGames {
		recommender.BoardGamesMap[i] = game
	}
}

// RecommendBoardGames Recommendation algorithm for board games using Characteristics
// Checks if categories or mechanisms match, we also account for favourable characteristics and give them a higher score
// TODO favourable should be a map
func (recommender *ContentBasedRecommender) RecommendBoardGames(boardGame shared.BoardGame, weight float64, favourable shared.Characteristic) []shared.Recommendation {
	var recommendations []shared.Recommendation
	giveGameToCompare := recommender.BoardGamesMap[boardGame.ID]

	// TODO this can probably be improved lol
	for _, game := range recommender.BoardGames {
		if game.ID == boardGame.ID {
			continue
		}

		commonCharacteristics := 0.0
		gameToCompare := recommender.BoardGamesMap[game.ID]
		for _, characteristic := range giveGameToCompare.Characteristic.Categories {
			for _, otherCharacteristic := range gameToCompare.Characteristic.Categories {
				if shared.DataEquals(characteristic, otherCharacteristic) {
					commonCharacteristics += 1
				}

				for _, favCategory := range favourable.Categories {
					if shared.DataEquals(favCategory, otherCharacteristic) {
						commonCharacteristics += weight
					}
				}
			}
		}

		for _, characteristic := range giveGameToCompare.Characteristic.Mechanisms {
			for _, otherCharacteristic := range gameToCompare.Characteristic.Mechanisms {
				if shared.DataEquals(characteristic, otherCharacteristic) {
					commonCharacteristics += 1
				}

				for _, favMechanism := range favourable.Mechanisms {
					if shared.DataEquals(favMechanism, otherCharacteristic) {
						commonCharacteristics += weight
					}
				}
			}
		}

		if commonCharacteristics >= 3 {
			recommendations = append(recommendations, shared.Recommendation{BoardGame: game, Score: commonCharacteristics})
		}
	}
	return recommendations
}
