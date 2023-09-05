package http

import (
	"github.com/OTeeEnabor/echo_learn/controller/context/pages"
	"github.com/labstack/echo/v4"
)

func IndexRouter(app *echo.Echo) {
	app.GET("/", pages.IndexContext)
}