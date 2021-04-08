package main

import (
	"log"

	"github.com/LarryKapija/shoppinglist_api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//==-===SHOPPINGLIST===-==\\
	//==> Create
	r.POST("/LIST", controllers.PostList)
	//==> Read
	r.GET("/LIST/:listId", controllers.GetList)
	//==> Update
	r.PUT("/LIST/:listId", controllers.PutList)
	//==> Delete
	r.DELETE("/LIST/:listId", controllers.DeleteList)
	//=======================\\
	//====-===ITEMS===-=====\\
	//==> Create
	r.POST("/LIST/:listId/ITEM", controllers.PostItems)
	//==> Read

	r.GET("/LIST/:listId/ITEM/:name", controllers.GetItems)
	//==> Update
	r.PUT("/LIST/:listId/ITEM/:name", controllers.PutItems)
	//==> Delete
	r.DELETE("/LIST/:listId/ITEM/:name", controllers.DeleteItems)
	//=======================\\

	log.Fatal(r.Run())

}
