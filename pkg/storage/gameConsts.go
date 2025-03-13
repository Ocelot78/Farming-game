package storage

import (
	"encoding/json"
	"log"
	"os"

	"farming_game/pkg/state"
)

func LoadTractors() state.Shop {
    data, err := os.ReadFile("assets/shop/tractors.json")
    if err != nil {
        log.Print("ERROR: FAILED TO READ SHOP STATE", err)
        return state.Shop{}
    }
    
    var tractorMap map[string]state.Tractor
    err = json.Unmarshal(data, &tractorMap)
    if err != nil {
        log.Print("ERROR: FAILED TO UNMARSHAL SHOP STATE", err)
        return state.Shop{}
    }
    
    var shop state.Shop
    for _, tractor := range tractorMap {
        shop.Tractors = append(shop.Tractors, tractor)
    }
    
    return shop
}