package controllers

import (
	"fmt"
	"strconv"
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
		c.JSON(utils.BadRequest, gin.H{"message": "Bad Request"})
		return
	}
	if ok := descriptionExists(list.Description, 0); ok {
		c.JSON(utils.Conflict, gin.H{"message": "Already Exists"})
		return
	}
	autoIncrement++
	list.Id = autoIncrement
	list.Date = time.Now()
	list.Items = make(map[string]models.Item)
	models.ShoppingLists[list.Id] = list

	c.JSON(utils.Created, list)
}

func GetLists(c *gin.Context) {
	defer utils.Recover(c)
	list := utils.ToList(models.ShoppingLists)
	c.JSON(utils.Ok, list)
}

func GetList(c *gin.Context) {
	defer utils.Recover(c)
	id, err := strconv.Atoi(c.Param("listId"))

	if err != nil {
		fmt.Println(err)
		return
	}
	list, ok := models.ShoppingLists[id]

	if !ok {
		fmt.Println("Not exists")
		c.JSON(utils.NotFound, gin.H{"message": "not found"})
	} else {
		c.JSON(utils.Ok, list)
	}
}

func PutList(c *gin.Context) {
	defer utils.Recover(c)
	id, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(utils.BadRequest, gin.H{"message": "Bad Request"})
		return
	}
	body := c.Request.Body
	var list models.ShoppingList
	if err := utils.ReadFromBody(body, &list); err != nil {
		fmt.Println(err.Error())
		c.JSON(utils.BadRequest, gin.H{"message": "Bad Request"})
		return
	}
	if _, ok := models.ShoppingLists[id]; !ok {
		c.JSON(utils.NotFound, gin.H{"message": "Not Found"})
		return
	}
	if ok := descriptionExists(list.Description, id); ok {
		c.JSON(utils.Conflict, gin.H{"message": "Description Already Exists"})
		return
	}
	list.Id = id
	list.Items = models.ShoppingLists[id].Items
	models.ShoppingLists[id] = list
	c.JSON(utils.Ok, list)
}

func DeleteList(c *gin.Context) {
	defer utils.Recover(c)
	id, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(utils.BadRequest, gin.H{"message": "Bad Request"})
		return
	}

	shopl, ok := models.ShoppingLists[id]
	if !ok {
		c.JSON(utils.NotFound, gin.H{"message": "not found"})
		return
	}

	delete(models.ShoppingLists, id)

	c.JSON(utils.Ok, shopl)

}

func descriptionExists(description string, key int) bool {
	for k, list := range models.ShoppingLists {
		if list.Description == description && k != key {
			return true
		}
	}
	return false
}
