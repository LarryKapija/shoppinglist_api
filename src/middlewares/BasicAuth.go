package middlewares

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/LarryKapija/shoppinglist_api/models"
	"github.com/LarryKapija/shoppinglist_api/utils"
	"github.com/gin-gonic/gin"
)

var Accounts gin.Accounts = make(gin.Accounts)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		p, ok := Accounts[username]
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if p != password {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
func Signup(c *gin.Context) {

	defer utils.Recover(c)
	body := c.Request.Body

	var user models.User

	if err := utils.ReadFromBody(body, &user); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	Accounts[user.Username] = user.Password
	token := createToken(user.Username, user.Password)
	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"token":    token,
	})
}

func createToken(username string, password string) string {
	auth := fmt.Sprintf("%s:%s", username, password)
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func Logout(c *gin.Context) {
	defer utils.Recover(c)
	body := c.Request.Body
	var user models.User
	if err := utils.ReadFromBody(body, &user); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	delete(Accounts, user.Username)
}
