package controllers

import (
	"crudv1/db"
	"crudv1/models"
	"net/http"

	"github.com/labstack/echo"
)

// Index
func Index(c echo.Context) error {
	db := db.DbManager()
	article := []models.Article{}
	db.Table("articles").Select("id,title,content").Scan(&article)
	// 	return c.Render(http.StatusOK, "index", map[string]interface{}{
	// 		"article": article,
	// 	})
	return c.JSON(http.StatusOK, article)
}

// store
func CreateArticle(c echo.Context) error {
	db := db.DbManager()
	article := new(models.Article)
	article.Title = c.FormValue("title")
	article.Content = c.FormValue("content")

	db.Create(&article)
	// return c.Redirect(http.StatusMovedPermanently, "/index")
	return c.JSON(http.StatusCreated, article)

}

// Edit
func Edit(c echo.Context) error {
	db := db.DbManager()
	article := []models.Article{}
	id := c.Param("id")
	db.Where("id = ?", id).Find(&article)
	// return c.Render(http.StatusOK, "edit", map[string]interface{}{
	// 	"articles": article,
	// 	"id":       id,
	// })
	return c.JSON(http.StatusOK, article)
}

// Update
func Update(c echo.Context) error {
	db := db.DbManager()
	article := []models.Article{}
	id := c.Param("id")
	title := c.FormValue("title")
	content := c.FormValue("content")
	db.Model(&article).Where("id= ?", id).Update(map[string]interface{}{"title": title, "content": content})
	// return c.Redirect(http.StatusMovedPermanently, "/index")
	return c.JSON(http.StatusOK, article)

}

// Delete
func Delete(c echo.Context) error {
	db := db.DbManager()
	article := []models.Article{}
	id := c.Param("id")
	db.Unscoped().Delete(&article, "id = ?", id)
	return c.JSON(http.StatusOK, article)
}

//searching article
func GetUser(c echo.Context) error {
	db := db.DbManager()
	article := []models.Article{}
	db.Where("id = ?", "1").Find(&article)
	return c.JSON(http.StatusOK, article)
}
