package respFormat

import (
	"github.com/gin-gonic/gin"
)

func Get200(c *gin.Context, data any) {
	// Get data successful
	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "get data Successful",
		"data":   data,
	})
	return
}

func Update200(c *gin.Context) {
	// update successful
	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "update data Successful",
		"data":   nil,
	})
	return
}

func Delete200(c *gin.Context) {
	// delete successful
	c.JSON(200, gin.H{
		"status": 0,
		"msg":    "delete data Successful",
		"data":   nil,
	})
	return
}

func Create201(c *gin.Context) {
	c.JSON(201, gin.H{
		"status": 0,
		"data":   nil,
		"msg":    "create data Successful",
	})
}

func Failed200(c *gin.Context, e any) {
	c.JSON(200, gin.H{
		"status": 1,
		"msg":    e,
		"data":   nil,
	})
	return
}
