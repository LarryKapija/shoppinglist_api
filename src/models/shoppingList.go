package models

import (
	"fmt"
	"time"
)

type ShoppingList struct {
	Id          int          `json:"id"`
	Date        time.Time    `json:"date"`
	Description string       `json:"description"`
	Items       map[int]Item `json:"items"`
}

func (l *ShoppingList) Update(list map[string]interface{}) {
	for k, v := range list {
		switch k {
		case "date":
			date, err := time.Parse(time.RFC3339Nano, v.(string))
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			l.Date = date
		case "description":
			l.Description = v.(string)
		case "items":
			items := v.(map[int]interface{})
			for x, y := range l.Items {
				y.Update(items[x].(map[string]interface{}))
			}
		}
	}
}

func (l *ShoppingList) ItemsToList() []Item {
	list := make([]Item, 0)
	for _, item := range l.Items {
		list = append(list, item)
	}
	return list
}

var ShoppingLists = make(map[int]ShoppingList)
