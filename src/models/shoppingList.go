package models

import "time"

type ShoppingList struct {
	Id          int          `json:"id"`
	Date        time.Time    `json:"date"`
	Description string       `json:"description"`
	Items       map[int]Item `json:"items"`
}

func (l *ShoppingList) ItemsToList() []Item {
	list := make([]Item, 0)
	for _, item := range l.Items {
		list = append(list, item)
	}
	return list
}

var ShoppingLists = make(map[int]ShoppingList)
