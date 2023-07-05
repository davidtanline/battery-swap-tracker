package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	// access db pass from secret.json
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	file, err := ioutil.ReadFile(currentDir + "/config.json")
	if err != nil {
		log.Fatal("Error reading configuration file:", err)
	}
	var connection map[string]interface{}
	err = json.Unmarshal(file, &connection)
	if err != nil {
		log.Fatal("Error parsing configuration file:", err)
	}
	dbPass, ok := connection["db_password"].(string)
	if !ok {
		log.Fatal("DB password not found or invalid type")
	}

	// dbDriver := "mysql"
	dbUser := "root"
	dbName := "battery_swap_tracker"

	// db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	db, err := gorm.Open(mysql.Open(dbUser+":"+dbPass+"@/"+dbName), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	// db.AutoMigrate(&models.Account{})

	return db
}
