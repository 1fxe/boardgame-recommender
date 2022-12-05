package internal

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
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	YearReleased   int            `json:"yearReleased"`
	NoPlayers      Range          `json:"noPlayers"`
	PlayTime       Range          `json:"playTime"`
	MinAge         int            `json:"age"`
	Characteristic Characteristic `json:"characteristic,omitempty"`
}
