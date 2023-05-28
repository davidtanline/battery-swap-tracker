package config

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func Connect() *sql.DB {
	// access db pass from secret.json
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := ioutil.ReadFile(currentDir + "/config.json")
	if err != nil {
		log.Fatal("Error reading configuration file:", err)
	}

	var config map[string]interface{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Error parsing configuration file:", err)
	}

	dbPass, ok := config["db_password"].(string)
	if !ok {
		log.Fatal("DB password not found or invalid type")
	}

	dbDriver := "mysql"
	dbUser := "root"
	dbName := "battery_swap_tracker"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}
