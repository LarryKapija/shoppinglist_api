package models

import "time"

type ShoppingList struct {
	Id          int             `json:"id"`
	Date        time.Time       `json:"date"`
	Description string          `json:"description"`
	Items       map[string]Item `json:"items"`
}

var ShoppingLists = make(map[int]ShoppingList)
