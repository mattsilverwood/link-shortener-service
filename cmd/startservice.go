package cmd

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func StartService() error {
	server := echo.New()

	server.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	server.GET("/:slug", func(c echo.Context) error {
		return c.String(http.StatusMovedPermanently, "Got "+c.Param("slug"))
	})

	return server.Start(":3000")
}
