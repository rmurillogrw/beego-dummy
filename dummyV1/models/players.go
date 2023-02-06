package models

import (
	"encoding/csv"
	"errors"
	"os"
)

// Player model
type Player struct {
	PlayerId string `json:"playerId"`
	Player   string `json:"player"`
	Country  string `json:"country"`
	Currency string `json:"currency"`
}

func GetPlayer(f *os.File) ([]Player, error) {
	var players []Player

	// 2. Read CSV file using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return players, errors.New("internal server error")
	}

	// 3. Assign successive lines of raw CSV data to fields of the created structs
	players = createShoppingList(data)

	// 4. Convert an array of structs to JSON using marshaling functions from the encoding/json package
	//jsonData, err := json.MarshalIndent(shoppingList, "", "  ")
	if err != nil {
		return players, errors.New("internal server error")
	}

	return players, nil
}

// receive a slice of slices of type string and return a player for every row
func createShoppingList(data [][]string) []Player {
	var playerList []Player
	for i, line := range data {
		if i > 0 { // omit header line
			var player Player
			player.PlayerId = line[0]
			player.Player = line[1]
			player.Country = line[2]
			player.Currency = line[3]
			playerList = append(playerList, player)
		}
	}
	return playerList
}
