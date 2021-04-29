package models

import (
	"encoding/json"
	"fmt"
)

type Doge struct {
	Data struct {
		Id struct {
			Quote struct {
				USD struct {
					Price              float64 `json:"price"`
					Volume_24h         int64   `json:"volume_24h"`
					Percent_change_1h  float64 `json:"Percent_change_1h"`
					Percent_change_24h float64 `json:"Percent_change_24h"`
					Percent_change_7d  float64 `json:"Percent_change_7d"`
					Percent_change_30d float64 `json:"Percent_change_30d"`
					Percent_change_60d float64 `json:"Percent_change_60d"`
					Percent_change_90d float64 `json:"percent_change_90d"`
					Market_cap         float64 `json:"market_cap"`
					Last_updated       string  `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"74"`
	} `json:"data"`
}

func BuildDogeStruct(body string) Doge {
	var doge Doge
	err := json.Unmarshal([]byte(body), &doge)

	if err != nil {
		fmt.Println(err)
	}

	return doge
}
