package models

import (
	"net/http"

	"github.com/farisraii/go-restAPI-echo/db"

	"github.com/go-playground/validator/v10"
)

type Karyawan struct {
	Id            int    `json:"id"`
	Nama          string `json:"nama" validate:"required"`
	Jenis_Kelamin string `json:"jenis_kelamin" validate:"required"`
	Usia          int    `json:"usia" validate:"required"`
	Telepon       string `json:"telephone" validate:"required"`
	Alamat        string `json:"alamat" validate:"required"`
}

func FetchAllData() (Response, error) {
	var obj Karyawan
	var arrobj []Karyawan
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM karyawans"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Jenis_Kelamin, &obj.Usia, &obj.Telepon, &obj.Alamat)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)

	}

	res.Status = http.StatusOK
	res.Message = "Successfully"
	res.Data = arrobj

	return res, nil

}

func StoreKaryawan(nama string, jeniskelamin string, usia int, telepon string, alamat string) (Response, error) {
	var res Response

	v := validator.New()

	kar := Karyawan{
		Nama:          nama,
		Jenis_Kelamin: jeniskelamin,
		Usia:          usia,
		Telepon:       telepon,
		Alamat:        alamat,
	}

	err := v.Struct(kar)

	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT karyawans (nama, jenis_kelamin, usia, telepon, alamat) VALUES (? , ? , ? , ? , ?)"

	stnt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stnt.Exec(nama, jeniskelamin, usia, telepon, alamat)

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

func UpdateKaryawan(id int, nama string, jeniskelamin string, usia int, telepon string, alamat string) (Response, error) {
	var res Response

	v := validator.New()

	kar := Karyawan{
		Nama:          nama,
		Jenis_Kelamin: jeniskelamin,
		Usia:          usia,
		Telepon:       telepon,
		Alamat:        alamat,
	}

	err := v.Struct(kar)

	if err != nil {
		return res, err
	}

	existingData, err := fetchDataByID(id)
	if err != nil {
		return res, err
	}

	if isDataEqual(existingData, kar) {
		res.Status = http.StatusOK
		res.Message = "Data tidak berubah"
		res.Data = map[string]int64{
			"rows_affected": 0,
		}
		return res, nil
	}

	con := db.CreateCon()

	sqlStatement := "UPDATE karyawans SET nama = ?, jenis_kelamin = ?, usia = ?, telepon = ?, alamat = ?  WHERE id = ?"

	stnt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stnt.Exec(nama, jeniskelamin, usia, telepon, alamat, id)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	if rowsAffected == 0 {

		res.Status = http.StatusNotFound
		res.Message = "Id not found"
		res.Data = map[string]int64{
			"rows_affected": rowsAffected,
		}

		return res, nil
	}

	res.Status = http.StatusOK
	res.Message = "Successfully Updated"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil

}

func DeleteKaryawan(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM karyawans WHERE id = ?"

	stnt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stnt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Successfully Deleted"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil

}

func isDataEqual(existingData Karyawan, newData Karyawan) bool {
	return existingData.Nama == newData.Nama &&
		existingData.Jenis_Kelamin == newData.Jenis_Kelamin &&
		existingData.Usia == newData.Usia &&
		existingData.Telepon == newData.Telepon &&
		existingData.Alamat == newData.Alamat
}

func fetchDataByID(id int) (Karyawan, error) {
	var kar Karyawan

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM karyawans WHERE id = ?"

	err := con.QueryRow(sqlStatement, id).Scan(
		&kar.Id, &kar.Nama, &kar.Jenis_Kelamin, &kar.Usia, &kar.Telepon, &kar.Alamat,
	)

	if err != nil {
		return kar, err
	}

	return kar, nil
}
