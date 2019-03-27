package route

import (
	"crudv1/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// exported Init
func Init() *echo.Echo {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
	}))
	e.GET("/index", controllers.Index)
	e.GET("/findArticle", controllers.GetUser)
	e.POST("/create", controllers.CreateArticle)
	e.PUT("/update/:id", controllers.Update)
	e.DELETE("/delete/:id", controllers.Delete)
	e.GET("/edit/:id", controllers.Edit)

	return e
}
