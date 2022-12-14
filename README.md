# BoardGame Recommender

## Content-based Filtering
Content-based filtering is a method of recommending items to users based on the characteristics of the items themselves. This is in contrast to other methods of recommendation, such as collaborative filtering, which use the actions of other users to make recommendations.

In content-based filtering, a system first gathers information about the items that are being recommended. This can include things like the category or genre of the item, the words used in the item's description, and any other relevant characteristics. The system then uses this information to identify items that are similar to each other, and recommends those items to users who have expressed an interest in similar items in the past.

One advantage of content-based filtering is that it does not require a lot of data about the preferences of individual users. Instead, it can make recommendations based on the inherent characteristics of the items themselves, which makes it possible to make recommendations even for users who have not provided much information about their preferences. Additionally, because the recommendations are based on the characteristics of the items themselves, they can be personalized to the individual user, taking into account their unique preferences and interests.

```go
type Characteristic struct {
    Categories []Data `json:"categories,omitempty"`
    Mechanisms []Data `json:"mechanisms,omitempty"`
}

type BoardGame struct {
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	YearReleased   int            `json:"yearReleased"`
	NoPlayers      Range          `json:"noPlayers"`
	PlayTime       Range          `json:"playTime"`
	MinAge         int            `json:"age"`
	Characteristic Characteristic `json:"characteristic"`
}
```

Our BoardGame struct looks like this, we can use the Characteristics fields to compare board games, using the categories and mechanisms fields.


## Collaborative Filtering

Collaborative filtering is a method of recommending items to users based on the actions and preferences of other users. This is in contrast to other methods of recommendation, such as content-based filtering, which use the characteristics of the items themselves to make recommendations.

In collaborative filtering, a system first gathers information about the actions of other users. This can include things like what items they have viewed or purchased, what ratings they have given to items, and any other relevant actions. The system then uses this information to identify users who have similar actions and preferences, and recommends items to a given user based on what similar users have liked or purchased in the past.

One advantage of collaborative filtering is that it can make recommendations even for users who have not provided much information about their preferences. By using the actions of other users, the system can infer what a given user might like, even if they have not explicitly stated their preferences. Additionally, because the recommendations are based on the actions of other users, they can be highly personalized, taking into account the unique preferences and interests of each individual user.

### Cosine Similarity
Cosine similarity is a measure of similarity between two non-zero vectors of an inner product space that measures the cosine of the angle between them. It is defined as the dot product of the two vectors divided by the product of their magnitudes.

Cosine similarity is often used in natural language processing and information retrieval to compare the similarity of two documents or sets of words. It is particularly useful because it is not affected by the magnitude of the vectors, only by their direction. This means that it can be used to compare vectors that have different lengths or magnitudes, as long as they are in the same inner product space.

Essentially the lower the angle between two vectors, the higher the cosine similarity between them. 


![img.png](.github/img.png)

[Source](https://www.oreilly.com/library/view/statistics-for-machine/9781788295758/eb9cd609-e44a-40a2-9c3a-f16fc4f5289a.xhtml)
