package main

import (
	"log"

	"github.com/LarryKapija/shoppinglist_api/controllers"
	"github.com/LarryKapija/shoppinglist_api/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middlewares.CacheControlHandler())
	//==-===SHOPPINGLIST===-==\\
	//==> Create
	r.POST("/LIST", controllers.PostList)
	//==> Read
	r.GET("/LIST", controllers.GetLists)
	r.HEAD("/LIST", controllers.GetLists)
	r.GET("/LIST/:listId", controllers.GetList)
	r.HEAD("/LIST/:listId", controllers.GetList)
	//==> Update
	r.PUT("/LIST/:listId", controllers.PutList)
	//==> Delete
	r.DELETE("/LIST/:listId", controllers.DeleteList)
	//=======================\\
	//====-===ITEMS===-=====\\
	//==> Create
	r.POST("/LIST/:listId/ITEM", controllers.PostItems)
	//==> Read
	r.GET("/LIST/:listId/ITEM", controllers.GetItems)
	r.HEAD("/LIST/:listId/ITEM", controllers.GetItems)
	r.GET("/LIST/:listId/ITEM/:id", controllers.GetItem)
	r.HEAD("/LIST/:listId/ITEM/:id", controllers.GetItem)
	//==> Update
	r.PUT("/LIST/:listId/ITEM/:id", controllers.PutItems)
	//==> Delete
	r.DELETE("/LIST/:listId/ITEM/:id", controllers.DeleteItems)
	//=======================\\

	log.Fatal(r.Run())

}
