package main

import (
	"log"

	"github.com/LarryKapija/shoppinglist_api/controllers"
	"github.com/LarryKapija/shoppinglist_api/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	auth := r.Group("/", middlewares.Authorize())
	auth.Use(middlewares.CacheControlHandler())
	//==-===SHOPPINGLIST===-==\\
	//==> Create
	auth.POST("/LIST", controllers.PostList)
	//==> Read
	auth.GET("/LIST", controllers.GetLists)
	auth.HEAD("/LIST", controllers.GetLists)
	auth.GET("/LIST/:listId", controllers.GetList)
	auth.HEAD("/LIST/:listId", controllers.GetList)
	//==> Update
	auth.PUT("/LIST/:listId", controllers.PutList)
	auth.PATCH("/LIST/:listId", controllers.PutList)
	//==> Delete
	auth.DELETE("/LIST/:listId", controllers.DeleteList)
	//=======================\\
	//====-===ITEMS===-=====\\
	//==> Create
	auth.POST("/LIST/:listId/ITEM", controllers.PostItems)
	//==> Read
	auth.GET("/LIST/:listId/ITEM", controllers.GetItems)
	auth.HEAD("/LIST/:listId/ITEM", controllers.GetItems)
	auth.GET("/LIST/:listId/ITEM/:id", controllers.GetItem)
	auth.HEAD("/LIST/:listId/ITEM/:id", controllers.GetItem)
	//==> Update
	auth.PUT("/LIST/:listId/ITEM/:id", controllers.PutItems)
	auth.PATCH("/LIST/:listId/ITEM/:id", controllers.PatchItems)
	//==> Delete
	auth.DELETE("/LIST/:listId/ITEM/:id", controllers.DeleteItems)
	//=======================\\
	r.POST("/SIGNIN", middlewares.Signup)
	r.POST("/LOGOUT", middlewares.Logout)
	log.Fatal(r.Run())

}
