package models

import (
	"net/http"

	"github.com/farisraii/go-restAPI-echo/db"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func StoreUser(username string, password string) (Response, error) {
	var res Response

	v := validator.New()

	user := User{
		Username: username,
		Password: password,
	}

	err := v.Struct(user)

	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT users (username, password) VALUES (? , ?)"

	stnt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stnt.Exec(username, password)

	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Successfully inserted"
	res.Data = map[string]int64{
		"last_insert_id": lastInsertedId,
	}

	return res, nil
}
