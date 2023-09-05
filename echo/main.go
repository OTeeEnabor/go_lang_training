package main

import (
	"net/http"

	"github.com/OTeeEnabor/echo_learn/server"
	"github.com/labstack/echo/v4"
)

func main() {
	//  initialize echo app
	app := echo.New()
	// define URL
	app.GET("/",func(context echo.Context) error {
			return context.String(http.StatusOK,
			"Hello world Echo")})
	server.SetServer(app)
	//  start the server
	app.Start(":3000")
	// fmt.Println("LEarning with golang")
}