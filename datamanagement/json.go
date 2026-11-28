package datamanagement

import (
	"encoding/json/v2"
	"fmt"
	"os"
)

type Item struct {
	ItemID       int    `json:"item_id"`
	ItemName     string `json:"item_name"`
	ItemCategory []string `json:"item_category"`
	BotEmojiID   string `json:"bot_emoji_id"`
}

type ItemCatalog struct {
	Items []Item `json:"items"`
}

func ParseIndexJson() ItemCatalog {
	dat, err := os.ReadFile("./items.json")
	if err != nil {
		fmt.Println(err)
	}

	var items ItemCatalog
	err = json.Unmarshal(dat, &items)
	if err != nil {
		fmt.Println(err)
	}

	return items
}
