package controllers

import (
	"net/http"
	"time"

	"github.com/farisraii/go-restAPI-echo/helpers"
	"github.com/farisraii/go-restAPI-echo/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Username is required"})

	}

	if password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Password is required"})

	}

	res, err := models.CheckLogin(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	// return c.String(http.StatusOK, "Successfully login")

	claims := &jwtCustomClaims{
		username,
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})

}

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}
