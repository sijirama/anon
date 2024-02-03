package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllRooms(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// func AddRoom(e echo.Context) error    {}
// func GetRoom(e echo.Context) error    {}
// func DeleteRoom(e echo.Context) error {} //NOTE: users can use passwords to delete a room
