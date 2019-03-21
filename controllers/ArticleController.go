package controllers

import (
	"crudv1/db"
	"crudv1/models"
	"net/http"

	"github.com/labstack/echo"
)

// Index
func Index(c echo.Context) error {
	// var data map[string]interface{}
	return c.Render(http.StatusOK, "index", map[string]interface{}{})
}

//get user
func GetUser(c echo.Context) error {
	db := db.DbManager()
	user := []models.Article{}
	db.Find(&user)
	return c.JSON(http.StatusOK, user)
}

// store
func CreateArticle(c echo.Context) error {
	db := db.DbManager()
	article := new(models.Article)
	article.Title = c.FormValue("title")
	article.Content = c.FormValue("content")

	db.Create(&article)
	return c.Render(http.StatusOK, "index", map[string]interface{}{
		"article": article,
	})

}

// update

// delete

// search
