package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

var items []Item
var nextID int = 1

func main() {
	r := gin.Default()

	r.POST("/create", createHandler)
	r.GET("/read", readHandler)
	r.PUT("/update", updateHandler)
	r.DELETE("/delete", deleteHandler)

	r.Run(":2566") // start the server
}

func createHandler(c *gin.Context) {
	var newItem Item
	if err := c.BindJSON(&newItem); err != nil {
		fmt.Println("Error during binding:", err) // Log the error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	newItem.ID = nextID
	nextID++
	items = append(items, newItem)
	c.JSON(http.StatusOK, newItem)
}

func readHandler(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

func updateHandler(c *gin.Context) {
	var updatedItem Item
	if err := c.BindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	for i, item := range items {
		if item.ID == updatedItem.ID {
			items[i] = updatedItem
			c.JSON(http.StatusOK, updatedItem)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

func deleteHandler(c *gin.Context) {
	idStr := c.DefaultQuery("id", "")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}
