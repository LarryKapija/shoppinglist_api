package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/LarryKapija/shoppinglist_api/models"
	"github.com/gin-gonic/gin"
)

func Recover(c *gin.Context) {
	if r := recover(); r != nil {
		fmt.Println(r)
		c.JSON(InternalServerError, gin.H{
			"message": r,
		})
	}
}

func ReadFromBody(body io.ReadCloser, val interface{}) error {
	value, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(value, &val); err != nil {
		return err
	}
	return nil
}

func ToList(values map[int]models.ShoppingList) []models.ShoppingList {
	list := make([]models.ShoppingList, 0)
	for _, value := range values {
		list = append(list, value)
	}
	return list
}
