package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	article struct {
		ID      int    `json:"id"`
		Content string `json:"content"`
	}
)

var (
	articles = map[int]*article{}
	seq      = 1
)

func createArticle(c echo.Context) error {
	a := &article{
		ID: seq,
	}
	if err := c.Bind(a); err != nil {
		return err
	}
	articles[a.ID] = a
	seq++
	return c.JSON(http.StatusCreated, a)
}

func getArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, articles[id])
}

func updateArticle(c echo.Context) error {
	a := new(article)
	if err := c.Bind(a); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	articles[id].Content = a.Content
	return c.JSON(http.StatusOK, articles[id])
}

func deleteArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(articles, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {

	// create instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/articles", createArticle)
	e.GET("/articles/:id", getArticle)
	e.PUT("/articles/:id", updateArticle)
	e.DELETE("/articles/:id", deleteArticle)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
