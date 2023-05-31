package model

type Battery struct {
	ID              int `form:"id" json:"id"`
	CurrentHealth   int `form:"current_health" json:"current_health"`
	MaximumCapacity int `form:"maximum_capacity" json:"maximum_capacity"`
}

type Car struct {
	ID        int    `form:"id" json:"id"`
	Type      string `form:"type" json:"type"`
	Year      int    `form:"year" json:"year"`
	BatteryID int    `form:"battery_id" json:"battery_id"`
}

type Location struct {
	ID      int    `form:"id" json:"id"`
	Address string `form:"address" json:"address"`
	City    string `form:"city" json:"city"`
	State   string `form:"state" json:"state"`
	ZipCode int    `form:"zip_code" json:"zip_code"`
}

type Station struct {
	Name              string `form:"name" json:"name"`
	LocationID        int    `form:"location_id" json:"location_id"`
	NumberOfBatteries int    `form:"number_of_batteries" json:"number_of_batteries"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    interface{}
}
