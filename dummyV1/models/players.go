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

	//Read CSV file+
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return players, errors.New("internal server error")
	}

	//  Assign successive lines of raw CSV data to fields of the created structs
	players = createShoppingList(data)

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
