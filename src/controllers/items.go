package controllers

import (
	"fmt"
	"net/http"

	"github.com/LarryKapija/shoppinglist_api/models"
	"github.com/LarryKapija/shoppinglist_api/utils"
	"github.com/gin-gonic/gin"
)

func PostItems(c *gin.Context) {
	defer utils.Recover(c)
	body := c.Request.Body
	listId := utils.StringToInt(c.Param("listId"))

	var item models.Item
	if err := utils.ReadFromBody(body, &item); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	_, err := utils.FindItem(listId, item.Id, false)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	if utils.NameExists(listId, item.Name) {
		c.JSON(http.StatusConflict, gin.H{"message": "Already Exists"})
		return
	}
	id := utils.GetLastId(listId)
	item.Id = id + 1
	if item.State == 0 {
		item.State = models.Pending
	}
	models.ShoppingLists[listId].Items[item.Id] = item
	c.JSON(http.StatusCreated, item)

}

func GetItems(c *gin.Context) {
	defer utils.Recover(c)
	queries := c.Request.URL.Query()
	listId := utils.StringToInt(c.Param("listId"))

	shoppingList, ok := models.ShoppingLists[listId]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "List Not Found"})
		return
	}
	list := shoppingList.ItemsToList()
	list = utils.ItemFilterBy(queries, list)
	c.JSON(http.StatusOK, list)
}

func GetItem(c *gin.Context) {
	defer utils.Recover(c)
	id := utils.StringToInt(c.Param("id"))
	listId := utils.StringToInt(c.Param("listId"))

	item, err := utils.FindItem(listId, id, true)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"name":     item.Name,
		"quantity": item.Quantity,
		"state":    item.State,
	})
}
func PutItems(c *gin.Context) {
	defer utils.Recover(c)
	id := utils.StringToInt(c.Param("id"))
	listId := utils.StringToInt(c.Param("listId"))
	body := c.Request.Body
	var item models.Item
	if err := utils.ReadFromBody(body, &item); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
	}
	_, err := utils.FindItem(listId, id, true)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	item.Id = id
	models.ShoppingLists[listId].Items[id] = item
	c.JSON(http.StatusOK, item)
}

func PatchItems(c *gin.Context) {
	defer utils.Recover(c)
	id := utils.StringToInt(c.Param("id"))
	listId := utils.StringToInt(c.Param("listId"))
	body := c.Request.Body
	var item map[string]interface{}
	if err := utils.ReadFromBody(body, &item); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
	}
	itemToUpdate, err := utils.FindItem(listId, id, true)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	itemToUpdate.Update(item)
	c.JSON(http.StatusOK, itemToUpdate)
}

func DeleteItems(c *gin.Context) {
	defer utils.Recover(c)
	id := utils.StringToInt(c.Param("id"))
	listId := utils.StringToInt(c.Param("listId"))

	item, err := utils.FindItem(listId, id, true)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}
	//Lazy remove
	item.State = models.Discarded
	models.ShoppingLists[listId].Items[id] = item

	//Remove
	// delete(models.ShoppingLists[listId].Items, item.Name)
	c.JSON(http.StatusOK, item)
}
