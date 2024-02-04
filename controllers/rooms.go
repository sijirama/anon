package controllers

import (
	"anon/db"
	"anon/models"
	"anon/public"
	"anon/template"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllRooms(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func CreateRoom(c echo.Context) error {
	component := public.AddRoom()
	return template.AssertRender(c, http.StatusOK, component)
}

func AddRoom(e echo.Context) error {
	room := new(models.Room)
	if err := e.Bind(room); err != nil {
		return err
	}
	fmt.Println(room)
	models.InsertRoom(db.DatabaseClient, room)
	return e.String(http.StatusOK, "Room has been created")
}

// func GetRoom(e echo.Context) error    {}
// func DeleteRoom(e echo.Context) error {} //NOTE: users can use passwords to delete a room
