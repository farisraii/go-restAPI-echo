package controllers

import (
	"net/http"
	"strconv"

	"github.com/farisraii/go-restAPI-echo/models"

	"github.com/labstack/echo/v4"
)

func FetchAllData(c echo.Context) error {
	result, err := models.FetchAllData()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func StoreKaryawan(c echo.Context) error {
	nama := c.FormValue("nama")
	jeniskelamin := c.FormValue("jenis_kelamin")
	usia := c.FormValue("usia")
	telepon := c.FormValue("telepon")
	alamat := c.FormValue("alamat")

	conv_usia, _ := strconv.Atoi(usia)

	result, err := models.StoreKaryawan(nama, jeniskelamin, conv_usia, telepon, alamat)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)

}

func UpdateKaryawan(c echo.Context) error {
	id := c.FormValue("id")
	nama := c.FormValue("nama")
	jeniskelamin := c.FormValue("jenis_kelamin")
	usia := c.FormValue("usia")
	telepon := c.FormValue("telepon")
	alamat := c.FormValue("alamat")

	conv_usia, _ := strconv.Atoi(usia)
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error)
	}

	result, err := models.UpdateKaryawan(conv_id, nama, jeniskelamin, conv_usia, telepon, alamat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteKaryawan(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error)
	}
	result, err := models.DeleteKaryawan(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error)
	}
	return c.JSON(http.StatusOK, result)
}
