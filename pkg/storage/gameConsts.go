package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"farming_game/pkg/state"
)

func LoadShop() state.Shop {
	data, err := os.ReadFile("assets/shop/tractors.json")
	if err != nil {
		log.Print("ERROR: FAILED TO READ SHOP STATE", err)
		return state.Shop{}
	}
	var shopState state.Shop
	err = json.Unmarshal(data, &shopState)
	if err != nil {
		log.Print("ERROR: FAILED TO UNMARSHAL SHOP STATE", err)
		return state.Shop{}
	}
	fmt.Print(shopState)
	return shopState
}