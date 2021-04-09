package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/LarryKapija/shoppinglist_api/models"
	"github.com/gin-gonic/gin"
	"github.com/hhsnopek/etag"
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
	if (method == "GET" || method == "HEAD") && value != "" {
		return strings.Contains(e, value)
	} else if method == "PUT" && value != "" {
		return !strings.Contains(e, value)
	}
	return false
}

func VersioningHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "public")
		c.Header("Etag", generateEtag(c.Request.URL.Path))
		if match := c.Request.Header.Get("If-None-Match"); match != "" {
			if strings.Contains(Etags[c.Request.URL.Path], match) {
				c.AbortWithStatus(http.StatusNotModified)
				return
			}
		}
		if match := c.Request.Header.Get("If-Match"); match != "" {
			if !strings.Contains(Etags[c.Request.URL.Path], match) {
				c.AbortWithStatus(http.StatusConflict)
				return
			}
		}

		c.Next()
	}
}

func generateEtag(path string) string {
	e := ""
	if strings.Compare(path, "/LIST") == 0 {
		value := models.ShoppingLists
		val := marshallValue(value)
		e = etag.Generate(val, false)
	} else if strings.Contains(path, "/ITEM/") {
		index := strings.Index(path, "/ITEM/")
		listId, err := strconv.Atoi(path[5:index])
		if err != nil {
			fmt.Println(err.Error())
			return ""
		}
		name := path[index+6:]
		value := models.ShoppingLists[listId].Items[name]
		val := marshallValue(value)
		e = etag.Generate(val, false)
	} else {
		listId, err := strconv.Atoi(path[5:])
		if err != nil {
			fmt.Println(err.Error())
			return ""
		}
		value := models.ShoppingLists[listId]
		val := marshallValue(value)
		e = etag.Generate(val, false)
	}
	Etags[path] = e
	return e
}

func marshallValue(value interface{}) []byte {
	val, err := json.Marshal(value)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return val
}
