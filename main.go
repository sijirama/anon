package main

import (
	"anon/config"
	"anon/utils"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	//NOTE: instantiate echo
	e := echo.New()

	//NOTE: set up the server and routes
	utils.SetUpServer(e)

	//NOTE: port stuff
	PORT := fmt.Sprintf(":%v", config.Get(config.PORT))

	//NOTE: start and log error if error
	e.Logger.Fatal(e.Start(PORT))
}
