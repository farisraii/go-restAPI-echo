package routes

import (
	"net/http"

	"github.com/farisraii/go-restAPI-echo/controllers"
	"github.com/farisraii/go-restAPI-echo/middleware"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Faris!")
	})

	e.GET("/karyawan", controllers.FetchAllData, middleware.IsAuthenticated)
	e.POST("/karyawan", controllers.StoreKaryawan, middleware.IsAuthenticated)
	e.PUT("/karyawan", controllers.UpdateKaryawan, middleware.IsAuthenticated)
	e.DELETE("/karyawan", controllers.DeleteKaryawan, middleware.IsAuthenticated)

	e.POST("/create/users", controllers.StoreUser)
	e.POST("/login", controllers.CheckLogin)
	// e.GET("/generate-hash/:password", controllers.GenerateHashPassword)

	return e

}
