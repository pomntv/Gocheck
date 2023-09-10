package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

var items []Item
var nextID int = 1

func main() {
	e := echo.New()

	e.POST("/create", createHandler)
	e.GET("/read", readHandler)
	e.PUT("/update", updateHandler)
	e.DELETE("/delete", deleteHandler)

	e.Logger.Fatal(e.Start(":2566"))
}

func createHandler(c echo.Context) error {
	var newItem Item
	if err := json.NewDecoder(c.Request().Body).Decode(&newItem); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON in request body")
	}

	newItem.ID = nextID
	nextID++
	items = append(items, newItem)

	return c.JSON(http.StatusOK, newItem)
}

func readHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, items)
}

func updateHandler(c echo.Context) error {
	var updatedItem Item
	if err := json.NewDecoder(c.Request().Body).Decode(&updatedItem); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON in request body")
	}

	for i, item := range items {
		if item.ID == updatedItem.ID {
			items[i] = updatedItem
			return c.JSON(http.StatusOK, updatedItem)
		}
	}

	return c.JSON(http.StatusNotFound, "Item not found")
}

func deleteHandler(c echo.Context) error {
	idStr := c.QueryParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			return c.JSON(http.StatusOK, item)
		}
	}

	return c.JSON(http.StatusNotFound, "Item not found")
}
