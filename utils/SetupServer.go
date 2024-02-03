package utils

import (
	"anon/controllers"
	"anon/db"
	"anon/public"
	"anon/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetUpServer(e *echo.Echo) {

	//INFO: initialize database
	db.DatabseInit()

	//INFO: template stuff
	e.Static("/dist", "dist")

	//INFO: templ renderer
	//template.NewTemplateRenderer(e, "public/*.html")
	template.NewTemplateRenderer(e)

	//INFO: run index html
	e.GET("/", func(c echo.Context) error {
		component := public.Index("sijibomi")
		return template.AssertRender(c, http.StatusOK, component)
	})

	e.GET("/rooms", controllers.GetAllRooms)

}
