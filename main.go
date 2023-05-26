package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"server/controller"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getBatteries", controller.AllBatteries).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Connected to port 3001")
	log.Fatal(http.ListenAndServe(":3001", router))
}
