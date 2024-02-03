package utils

import (
	"anon/controllers"
	"anon/db"
	"anon/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetUpServer(e *echo.Echo) {

	//INFO: initialize database
	db.DatabseInit()

	//INFO: template stuff
	e.Static("/dist", "dist")
	template.NewTemplateRenderer(e, "public/*.html")

	//INFO: run index html
	e.GET("/", func(c echo.Context) error {
		//return c.String(http.StatusOK, "Hello, World!")
		res := map[string]interface{}{
			"Name":  "Oluwsasijibomi",
			"Phone": "8888888",
			"Email": "skyscraper@gmail.com",
		}
		return c.Render(http.StatusOK, "index", res)
	})

	e.GET("/rooms", controllers.GetAllRooms)

}
