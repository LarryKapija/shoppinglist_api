package utils

import (
	"net/url"
	"strings"
	"time"

	"github.com/LarryKapija/shoppinglist_api/models"
)

func ToList(values map[int]models.ShoppingList) []models.ShoppingList {
	list := make([]models.ShoppingList, 0)
	for _, value := range values {
		list = append(list, value)
	}
	return list
}

func DescriptionExists(description string, key int) bool {
	for k, list := range models.ShoppingLists {
		if list.Description == description && k != key {
			return true
		}
	}
	return false
}

func listFilterBy(list []models.ShoppingList, f func(models.ShoppingList) bool) []models.ShoppingList {
	response := make([]models.ShoppingList, 0)
	for _, v := range list {
		if f(v) {
			response = append(response, v)
		}
	}
	return response
}

func ListFilterBy(queries url.Values, list []models.ShoppingList) []models.ShoppingList {
	for k, v := range queries {
		val := strings.Join(v, "")
		switch k {
		case "description":
			list = listFilterBy(list, func(shopl models.ShoppingList) bool {
				return strings.Contains(shopl.Description, val)
			})
		case "date":
			list = listFilterBy(list, func(shopl models.ShoppingList) bool {
				date := StringToTime(val)
				return shopl.Date.Truncate(24 * time.Hour).Equal(date.Truncate(24 * time.Hour))
			})
		case "date>x":
			list = listFilterBy(list, func(shopl models.ShoppingList) bool {
				date := StringToTime(val)
				return shopl.Date.Truncate(24 * time.Hour).After(date.Truncate(24 * time.Hour))
			})
		case "date<x":
			list = listFilterBy(list, func(shopl models.ShoppingList) bool {
				date := StringToTime(val)
				return shopl.Date.Truncate(24 * time.Hour).Before(date.Truncate(24 * time.Hour))
			})
		}
	}
	return list
}
