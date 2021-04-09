package controllers

import (
	"fmt"
	"strconv"

	"github.com/LarryKapija/shoppinglist_api/models"
	"github.com/LarryKapija/shoppinglist_api/utils"
	"github.com/gin-gonic/gin"
)

func PostItems(c *gin.Context) {
	defer utils.Recover(c)
	body := c.Request.Body
	listId, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var item models.Item
	if err := utils.ReadFromBody(body, &item); err != nil {
		fmt.Println(err.Error())
		c.JSON(utils.BadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	_, err = findItem(listId, item.Name, false)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(utils.NotFound, gin.H{"message": err.Error()})
		return
	}
	if _, ok := models.ShoppingLists[listId].Items[item.Name]; ok {
		c.JSON(utils.Conflict, gin.H{"message": "Already Exists"})
		return
	}
	item.State = models.Pending
	models.ShoppingLists[listId].Items[item.Name] = item
	c.JSON(utils.Created, item)

}

func GetItems(c *gin.Context) {
	defer utils.Recover(c)
	itemName := c.Param("name")
	listId, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		fmt.Println(err)
		return
	}

	item, err := findItem(listId, itemName, true)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(utils.NotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(utils.Ok, gin.H{
		"name":     item.Name,
		"quantity": item.Quantity,
		"state":    item.State,
	})
}
func PutItems(c *gin.Context) {
	defer utils.Recover(c)
	name := c.Param("name")
	listId, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		fmt.Println(err)
		return
	}
	body := c.Request.Body
	var item models.Item
	if err := utils.ReadFromBody(body, &item); err != nil {
		fmt.Println(err.Error())
		c.JSON(utils.BadRequest, gin.H{"message": "Bad Request"})
	}
	_, err = findItem(listId, name, true)
	if err != nil {
		c.JSON(utils.NotFound, gin.H{"message": err.Error()})
		return
	}
	item.Name = name
	models.ShoppingLists[listId].Items[name] = item
	c.JSON(utils.Ok, item)
}

func DeleteItems(c *gin.Context) {
	defer utils.Recover(c)
	name := c.Param("name")
	listIdString := c.Param("listId")

	listId, err := strconv.Atoi(listIdString)

	if err != nil {
		fmt.Println(err)
		return
	}

	item, err := findItem(listId, name, true)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(utils.NotFound, gin.H{"message": err.Error()})
	}
	//Lazy remove
	item.State = models.Discarded
	models.ShoppingLists[listId].Items[name] = item

	//Remove
	// delete(models.ShoppingLists[listId].Items, item.Name)
	c.JSON(utils.Ok, item)
}

func findItem(listId int, name string, isnotFoundError bool) (models.Item, error) {
	_, ok := models.ShoppingLists[listId]

	if !ok {
		return models.Item{}, fmt.Errorf("List Not Found")
	}

	item, ok2 := models.ShoppingLists[listId].Items[name]
	if !ok2 && isnotFoundError {
		return models.Item{}, fmt.Errorf("Item Not Found")
	}
	return item, nil
}
