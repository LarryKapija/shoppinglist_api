package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/LarryKapija/shoppinglist_api/models"
	"github.com/gin-gonic/gin"
)

var Etags map[string]string = make(map[string]string)

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
func EvaluatePreconditions(path string, value string, method string) bool {
	fmt.Println(path, value, method)
	e := Etags[path]
	fmt.Println(e)
	if strings.Compare(e, value) == 0 && (method == "GET" || method == "HEAD") {
		return false
	}
	return true
}
