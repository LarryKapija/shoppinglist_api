package utils

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/LarryKapija/shoppinglist_api/models"
)

func FindItem(listId int, id int, isnotFoundError bool) (models.Item, error) {
	_, ok := models.ShoppingLists[listId]

	if !ok {
		return models.Item{}, fmt.Errorf("List Not Found")
	}

	item, ok2 := models.ShoppingLists[listId].Items[id]
	if !ok2 && isnotFoundError {
		return models.Item{}, fmt.Errorf("Item Not Found")
	}
	return item, nil
}

func NameExists(listId int, name string) bool {
	list := models.ShoppingLists[listId]

	for _, item := range list.Items {
		if item.Name == name {
			return true
		}
	}
	return false
}

func GetLastId(listId int) int {
	list := models.ShoppingLists[listId]
	var id int = 0
	for _, item := range list.Items {
		if item.Id > id {
			id = item.Id
		}
	}
	return id
}

func itemFilterBy(list []models.Item, f func(models.Item) bool) []models.Item {
	response := make([]models.Item, 0)
	for _, v := range list {
		if f(v) {
			response = append(response, v)
		}
	}
	return response
}

func ItemFilterBy(queries url.Values, list []models.Item) []models.Item {
	for k, v := range queries {
		val := strings.Join(v, "")
		switch k {
		case "state":
			list = itemFilterBy(list, func(item models.Item) bool {
				state := models.ToId(val)
				return item.State == state
			})
		case "name":
			list = itemFilterBy(list, func(item models.Item) bool {
				return strings.Contains(item.Name, val)
			})
		case "quantity":
			list = itemFilterBy(list, func(item models.Item) bool {
				value := StringToFloat(val)
				if value == -1 {
					return false
				}
				return item.Quantity == value
			})
		}
	}
	return list
}
