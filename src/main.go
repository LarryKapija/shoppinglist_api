package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// State = Pendiente compra | Comprado | Descartado
type State int

const (
	Pending State = iota + 1
	Bought
	Discarded
)

var autoIncrement = 0

type Item struct {
	Name     string  `json:"name"`
	Quantity float32 `json:"quantity"`
	State    State   `json:"state"`
}

type ShoppingList struct {
	Id          int             `json:"id"`
	Date        time.Time       `json:"date"`
	Description string          `json:"description"`
	Items       map[string]Item `json:"items"`
}

var shoppingLists = make(map[int]ShoppingList)

func main() {
	r := gin.Default()

	//==-===SHOPPINGLIST===-==\\
	//==> Create
	r.POST("/LIST", PostList)
	//==> Read
	r.GET("/LIST/:listId", GetList)
	//==> Update
	r.PUT("/LIST/:listId", PutList)
	//==> Delete
	r.DELETE("/LIST/:listId", DeleteList)
	//=======================\\
	//====-===ITEMS===-=====\\
	//==> Create
	r.POST("/LIST/:listId/ITEM", PostItems)
	//==> Read

	r.GET("/LIST/:listId/ITEM/:name", GetItems)
	//==> Update
	r.PUT("/LIST/:listId/ITEM/:name", PutItems)
	//==> Delete
	r.DELETE("/LIST/:listId/ITEM/:name", DeleteItems)
	//=======================\\

	r.Run()

}

//==========ITEMS===========\\

func PostItems(c *gin.Context) {
	defer Recover(c)
	body := c.Request.Body
	listId, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var item Item
	if err := json.Unmarshal(value, &item); err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	_, ok := shoppingLists[listId]
	if !ok {
		c.JSON(404, gin.H{
			"message": "Not Found",
		})
	}
	shoppingLists[listId].Items[item.Name] = item
	c.JSON(200, item)

}

func GetItems(c *gin.Context) {
	defer Recover(c)
	listIdString := c.Param("listId")
	itemName := c.Param("name")

	listId, err := strconv.Atoi(listIdString)

	if err != nil {
		fmt.Println(err)
		return
	}

	list, ok := shoppingLists[listId]

	if !ok {
		fmt.Println("Not exists")
		c.JSON(404, gin.H{
			"message": "not found",
		})
	} else {
		item := list.Items[itemName]
		c.JSON(200, gin.H{
			"name":     item.Name,
			"quantity": item.Quantity,
			"state":    item.State,
		})
	}
}
func PutItems(c *gin.Context) {
	defer Recover(c)
	name := c.Param("name")
	listId, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		fmt.Println(err)
		return
	}
	body := c.Request.Body
	var item Item
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := json.Unmarshal(value, &item); err != nil {
		fmt.Println(err)
		return
	}

	_, ok := shoppingLists[listId]

	if !ok {
		c.JSON(404, gin.H{
			"message": "not found",
		})
		return
	}

	_, ok2 := shoppingLists[listId].Items[name]
	if !ok2 {
		c.JSON(404, gin.H{
			"message": "not found",
		})
		return
	}
	shoppingLists[listId].Items[name] = item

	c.JSON(200, item)
}

func DeleteItems(c *gin.Context) {
	defer Recover(c)
	name := c.Param("name")
	listIdString := c.Param("listId")

	listId, err := strconv.Atoi(listIdString)

	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO check bad request

	item, ok := shoppingLists[listId].Items[name]

	if !ok {
		c.JSON(404, gin.H{
			"message": "not found",
		})
	}

	c.JSON(200, item)
}

//=========SHOPPINGLIST======\\
func PostList(c *gin.Context) {
	defer Recover(c)
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var list ShoppingList
	if err := json.Unmarshal(value, &list); err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}
	autoIncrement++
	list.Id = autoIncrement
	list.Date = time.Now()
	list.Items = make(map[string]Item)
	shoppingLists[list.Id] = list

	c.JSON(201, list)
}

func GetList(c *gin.Context) {
	defer Recover(c)
	idString := c.Param("listId")
	id, err := strconv.Atoi(idString)

	if err != nil {
		fmt.Println(err)
		return
	}
	list, ok := shoppingLists[id]

	if !ok {
		fmt.Println("Not exists")
		c.JSON(404, gin.H{
			"message": "not found",
		})
	} else {
		c.JSON(200, gin.H{
			"id":          list.Id,
			"date":        list.Date,
			"description": list.Description,
			"items":       list.Items,
		})
	}

}
func PutList(c *gin.Context) {
	defer Recover(c)
	id, err := strconv.Atoi(c.Param("listId"))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var list ShoppingList
	if err := json.Unmarshal(value, &list); err != nil {
		fmt.Println(err.Error())
		return
	}
	if _, ok := shoppingLists[id]; !ok {
		fmt.Println(err.Error())
		c.JSON(404, gin.H{
			"message": "Not Found",
		})
		return
	}
	list.Id = id
	list.Items = shoppingLists[id].Items
	shoppingLists[id] = list
	c.JSON(200, list)

}

func DeleteList(c *gin.Context) {
	defer Recover(c)
	id, err := strconv.Atoi(c.Param("listId"))
	// TODO check bad request
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	shopl, ok := shoppingLists[id]
	if !ok {
		c.JSON(404, gin.H{
			"message": "not found",
		})
	}

	delete(shoppingLists, id)

	c.JSON(200, shopl)

}

func Recover(c *gin.Context) {
	if r := recover(); r != nil {
		fmt.Println(r)
		c.JSON(500, gin.H{
			"message": r,
		})
	}
}
