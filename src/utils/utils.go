package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Recover(c *gin.Context) {
	if r := recover(); r != nil {
		fmt.Println(r)
		c.JSON(InternalServerError, gin.H{
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
