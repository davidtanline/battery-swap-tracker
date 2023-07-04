package main

import (
	"server/database"
	"server/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
	// router := mux.NewRouter()
	// router.HandleFunc("/getBatteries", controller.GetAllBatteries).Methods("GET")
	// http.Handle("/", router)
	// fmt.Println("Connected to port 3001")
	// log.Fatal(http.ListenAndServe(":3001", router))
}
