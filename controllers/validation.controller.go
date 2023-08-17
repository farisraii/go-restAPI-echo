package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Customer struct {
	Nama   string `validate:"required"`
	Email  string `validate:"required,email"`
	Alamat string `validate:"required"`
	Umur   int    `validate:"gte=17"`
}

func TestStructValidation(c echo.Context) error {
	v := validator.New()

	cust := Customer{
		Nama:   "Faris",
		Email:  "farisraii96@gmail.com",
		Alamat: "Dago",
		Umur:   26,
	}

	err := v.Struct(cust)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully"})
}
