package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/LarryKapija/shoppinglist_api/models"
	"github.com/gin-gonic/gin"
	"github.com/hhsnopek/etag"
)

var Etags map[string]string = make(map[string]string)

func Recover(c *gin.Context) {
	if r := recover(); r != nil {
		fmt.Println(r)
		c.JSON(http.StatusInternalServerError, gin.H{
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

func GenerateEtag(path string) string {
	e := ""
	if strings.Compare(path, "/LIST") == 0 {

		value := models.ShoppingLists
		val := marshallValue(value)
		e = etag.Generate(val, false)

	} else if strings.Contains(path, "/ITEM/") {

		index := strings.Index(path, "/ITEM/")
		listId := StringToInt(path[5:index])
		id := StringToInt(path[index+6:])
		value := models.ShoppingLists[listId].Items[id]
		val := marshallValue(value)
		e = etag.Generate(val, false)

	} else if strings.Contains(path, "/ITEM") {

		index := strings.Index(path, "/ITEM")
		listId := StringToInt(path[5:index])
		value := models.ShoppingLists[listId].Items
		val := marshallValue(value)
		e = etag.Generate(val, false)

	} else {

		listId := StringToInt(path[5:])
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

func StringToInt(value string) int {
	num, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	return num
}

func StringToFloat(value string) float32 {
	val, err := strconv.ParseFloat(value, 32)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	return float32(val)
}

func StringToTime(value string) time.Time {
	val, err := time.Parse(time.RFC3339Nano, value)
	if err != nil {
		fmt.Println(err.Error())
		return time.Time{}
	}
	return val
}
