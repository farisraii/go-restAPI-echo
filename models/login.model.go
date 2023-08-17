package models

import (
	"database/sql"
	"fmt"

	"github.com/farisraii/go-restAPI-echo/db"
	"github.com/farisraii/go-restAPI-echo/helpers"
)

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and password mismatch")
		return false, err
	}

	return true, nil

}
