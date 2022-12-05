package main

import (
	"log"
	"math"
	"math/rand"
	"sort"
	"time"
)

// BoardGame represents a board game
// TODO Add other fields to the BoardGame struct
type BoardGame struct {
	ID   int
	Name string
}

// User represents a user and their ratings for different board games
type User struct {
	ID      int
	Ratings map[int]int // map of game name to rating
}

type Recommendation struct {
	Game  BoardGame
	Score float64
}

// BoardGameRecommender is a collaborative filtering-based board game recommender
type BoardGameRecommender struct {
	BoardGames map[int]BoardGame // map of game ID to game
	Users      map[int]User      // map of user ID to user
}

// AddGame adds a new game to the Recommender
func (recommender *BoardGameRecommender) AddGame(id int, name string) {
	recommender.BoardGames[id] = BoardGame{
		ID:   id,
		Name: name,
	}
}

// AddUser adds a new user to the Recommender
func (recommender *BoardGameRecommender) AddUser(id int, ratings map[int]int) {
	recommender.Users[id] = User{
		ID:      id,
		Ratings: ratings,
	}
}

// Recommendations calculates and returns recommendations for a given user
func (recommender *BoardGameRecommender) Recommendations(userID int) []Recommendation {
	// Calculate similarity between the target user and all other users
	similarities := make(map[int]float64)
	user := recommender.Users[userID]
	for otherID, otherUser := range recommender.Users {
		if otherID == userID {
			continue
		}

		// Calculate cosine similarity between the two users
		similarity := cosineSimilarity(user.Ratings, otherUser.Ratings)
		similarities[otherID] = similarity
	}

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

	// Recommend games that the most similar users have rated highly
	var recommendations []Recommendation
	for _, id := range sortedIDs {
		otherUser := recommender.Users[id]
		for gameID, rating := range otherUser.Ratings {
			// If the User has not previously rated this game, and the other user rated it highly, recommend it
			if user.Ratings[gameID] == 0 && rating > 3 {
				// TODO ignore duplicate recommendations
				recommendations = append(recommendations, Recommendation{
					recommender.BoardGames[gameID],
					similarities[id],
				})
			}
		}
	}

	return recommendations
}

// cosineSimilarity calculates the cosine similarity between two maps of ratings
func cosineSimilarity(userRatings, otherUserRatings map[int]int) float64 {
	var dotProduct, magnitude1, magnitude2 float64

	// Calculate the dot product
	for gameID := range userRatings {
		rating1 := userRatings[gameID]
		rating2 := otherUserRatings[gameID]

		dotProduct += float64(rating1 * rating2)
		magnitude1 += float64(rating1 * rating1)
		magnitude2 += float64(rating2 * rating2)
	}

	magnitude1 = math.Sqrt(magnitude1)
	magnitude2 = math.Sqrt(magnitude2)

	return dotProduct / (magnitude1 * magnitude2)
}

func main() {
	// Create a new recommender
	recommender := &BoardGameRecommender{
		BoardGames: make(map[int]BoardGame),
		Users:      make(map[int]User),
	}

	// Add some games
	recommender.AddGame(0, "Settlers of Catan")
	recommender.AddGame(1, "Agricola")
	recommender.AddGame(2, "Pandemic")
	recommender.AddGame(3, "Agricola")
	recommender.AddGame(4, "Power Grid")
	recommender.AddGame(5, "Dominion")
	recommender.AddGame(6, "Ticket to Ride")
	recommender.AddGame(7, "Carcassonne")
	recommender.AddGame(8, "Terra Mystica")
	recommender.AddGame(9, "7 Wonders")
	recommender.AddGame(10, "Puerto Rico")
	recommender.AddGame(11, "Through the Ages: A New Story of Civilization")
	recommender.AddGame(12, "Scrabble")
	recommender.AddGame(13, "Twilight Struggle")
	recommender.AddGame(14, "Stone Age")
	recommender.AddGame(15, "Small World")
	recommender.AddGame(16, "Splendor")
	recommender.AddGame(17, "Catan: Cities and Knights")
	recommender.AddGame(18, "A Feast for Odin")
	recommender.AddGame(19, "Castles of Burgundy")

	rand.Seed(time.Now().UnixNano())
	numberOfBoardGames := len(recommender.BoardGames)
	// Add some users
	for user := 0; user < 200; user++ {
		ratings := make(map[int]int)
		for gameId := 0; gameId < numberOfBoardGames; gameId++ {
			ratings[gameId] = rand.Intn(5) + 1

			// Test Data for user 23
			if user == 23 && rand.Int()%2 == 0 {
				ratings[gameId] = 0
			}
		}
		recommender.AddUser(user, ratings)
	}

	recommendedGames := recommender.Recommendations(23)

	log.Println("Recommended games:")
	for _, recommendation := range recommendedGames[:5] {
		log.Println(recommendation.Game.Name, recommendation.Score)
	}
}
