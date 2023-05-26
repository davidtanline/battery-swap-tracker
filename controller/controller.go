package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"server/model"

	"server/config"
)

func AllBatteries(w http.ResponseWriter, r *http.Request) {
	var battery model.Battery
	var response model.Response
	var arrBattery []model.Battery

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Battery")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&battery.ID, &battery.CurrentHealth, &battery.MaximumCapacity)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrBattery = append(arrBattery, battery)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrBattery

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
