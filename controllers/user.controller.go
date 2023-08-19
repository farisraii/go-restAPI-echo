package controllers

import (
	"net/http"

	"github.com/farisraii/go-restAPI-echo/helpers"
	"github.com/farisraii/go-restAPI-echo/models"
	"github.com/labstack/echo/v4"
)

func StoreUser(c echo.Context) error {
	username := c.FormValue("username")
	pwd := c.FormValue("password")

	if username == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Username is required"})

	}

	if pwd == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Password is required"})

	}

	pwdHashed, _ := helpers.HashPassword(pwd)

	result, err := models.StoreUser(username, pwdHashed)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)

}
