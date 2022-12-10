package recommenders

import (
	"math"
	"sort"

	"github.com/1fxe/board-game-recommender-system/shared"
)

// CollaborativeRecommender is a collaborative filtering-based board game recommender
type CollaborativeRecommender struct {
	BoardGames map[int]shared.BoardGame // map of game ID to game
	Users      map[int]shared.User      // map of user ID to user
}

// PopulateBoardGames populates the recommender's BoardGames map
func (recommender *CollaborativeRecommender) PopulateBoardGames(boardGames []shared.BoardGame) {
	recommender.BoardGames = make(map[int]shared.BoardGame)
	for i, game := range boardGames {
		recommender.BoardGames[i] = game
	}
}

// PopulateUsers populates the recommender's Users map
func (recommender *CollaborativeRecommender) PopulateUsers(users []shared.User) {
	recommender.Users = make(map[int]shared.User)
	for i, user := range users {
		recommender.Users[i] = user
	}
}

// RecommendBoardGames calculates and returns recommendations for a given user
func (recommender *CollaborativeRecommender) RecommendBoardGames(userID int) []shared.Recommendation {
	// Calculate similarity between the target user and all other users
	similarities := make(map[int]float64)
	user := recommender.Users[userID]

	for otherID, otherUser := range recommender.Users {
		if user.ID == otherUser.ID {
			continue
		}

		// Calculate cosine similarity between the two users
		similarity := cosineSimilarity(user.Ratings, otherUser.Ratings)
		similarities[otherID] = similarity
	}

	sortedIDs := recommender.gatherSimilarUsers(similarities)

	// Recommend games that the most similar users have rated highly
	var recommendations []shared.Recommendation
	for _, id := range sortedIDs {
		otherUser := recommender.Users[id]
		for gameID, rating := range otherUser.Ratings {
			// If the User has not previously rated this game, and the other user rated it highly, recommend it
			if user.Ratings[gameID] == 0 && rating > 3 {
				// TODO ignore duplicate recommendations
				recommendations = append(recommendations, shared.Recommendation{
					BoardGame: recommender.BoardGames[gameID],
					Score:     similarities[id],
				})
			}
		}
	}

	sortRecommendations(recommendations)

	return recommendations
}

func (recommender *CollaborativeRecommender) gatherSimilarUsers(similarities map[int]float64) []int {
	// Sort the other users by similarity to the target user
	var sortedIDs []int
	for id := range similarities {
		sortedIDs = append(sortedIDs, id)
	}

	sort.Slice(sortedIDs, func(i, j int) bool {
		return similarities[sortedIDs[i]] > similarities[sortedIDs[j]]
	})

	// Get the top 25 most similar users
	sortedIDs = sortedIDs[:25]
	return sortedIDs
}

// cosineSimilarity calculates the cosine similarity between two maps of ratings
func cosineSimilarity(userRatings, otherUserRatings map[int]int) float64 {
	var dotProduct, magnitude1, magnitude2 float64

	for gameID := range userRatings {
		rating1 := userRatings[gameID]
		rating2 := otherUserRatings[gameID]

		dotProduct += float64(rating1 * rating2)
		magnitude1 += float64(rating1 * rating1)
		magnitude2 += float64(rating2 * rating2)
	}

	magnitude1 = math.Sqrt(magnitude1)
	magnitude2 = math.Sqrt(magnitude2)

	cosineSimilarity := dotProduct / (magnitude1 * magnitude2)

	if math.IsNaN(cosineSimilarity) {
		return 0
	} else {
		return cosineSimilarity
	}
}

func sortRecommendations(recommendations []shared.Recommendation) {
	sort.Slice(recommendations, func(i, j int) bool {
		return recommendations[i].Score > recommendations[j].Score
	})
}
