package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/LarryKapija/shoppinglist_api/models"
	"github.com/LarryKapija/shoppinglist_api/utils"
	"github.com/gin-gonic/gin"
)

var autoIncrement int = 0

func PostList(c *gin.Context) {
	defer utils.Recover(c)
	body := c.Request.Body
	var list models.ShoppingList
	if err := utils.ReadFromBody(body, &list); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}
	if ok := utils.DescriptionExists(list.Description, 0); ok {
		c.JSON(http.StatusConflict, gin.H{"message": "Already Exists"})
		return
	}
	autoIncrement++
	list.Id = autoIncrement
	list.Date = time.Now()
	if list.Items == nil {
		list.Items = make(map[int]models.Item)
	}
	models.ShoppingLists[list.Id] = list

	c.JSON(http.StatusCreated, list)
}

func GetLists(c *gin.Context) {
	defer utils.Recover(c)
	queryStrings := c.Request.URL.Query()
	list := utils.ToList(models.ShoppingLists)
	list = utils.ListFilterBy(queryStrings, list)
	c.JSON(http.StatusOK, list)
}

func GetList(c *gin.Context) {
	defer utils.Recover(c)
	id := utils.StringToInt(c.Param("listId"))
	list, ok := models.ShoppingLists[id]

	if !ok {
		fmt.Println("Not exists")
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		c.JSON(http.StatusOK, list)
	}
}

func PutList(c *gin.Context) {
	defer utils.Recover(c)
	id := utils.StringToInt(c.Param("listId"))

	body := c.Request.Body
	var list models.ShoppingList
	if err := utils.ReadFromBody(body, &list); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}
	if _, ok := models.ShoppingLists[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	if ok := utils.DescriptionExists(list.Description, id); ok {
		c.JSON(http.StatusConflict, gin.H{"message": "Description Already Exists"})
		return
	}
	list.Id = id
	if list.Items == nil {
		list.Items = models.ShoppingLists[id].Items
	}
	models.ShoppingLists[id] = list
	c.JSON(http.StatusOK, list)
}

func DeleteList(c *gin.Context) {
	defer utils.Recover(c)
	id := utils.StringToInt(c.Param("listId"))

	shopl, ok := models.ShoppingLists[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}

	delete(models.ShoppingLists, id)

	c.JSON(http.StatusOK, shopl)

}
