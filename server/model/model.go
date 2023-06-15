package model

type Battery struct {
	ID              int `form:"id" json:"id"`
	CurrentHealth   int `form:"current_health" json:"current_health"`
	MaximumCapacity int `form:"maximum_capacity" json:"maximum_capacity"`
}

type Car struct {
	ID        int `form:"id" json:"id"`
	BatteryID int `form:"battery_id" json:"battery_id"`
}

type Station struct {
	Name        string `form:"name" json:"name"`
	Address     int    `form:"address" json:"address"`
	InOperation bool   `form:"in_operation" json:"in_operation"`
}

type Account struct {
	ID          int    `form:"id" json:"id"`
	FirstName   string `form:"first_name" json:"first_name"`
	LastName    string `form:"last_name" json:"last_name"`
	Email       string `form:"email" json:"email"`
	Pass        string `form:"pass" json:"pass"`
	AccountType string `form:"account_type" json:"account_type"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    interface{}
}
