package models

type Battery struct {
	ID              int `form:"ID" json:"id"`
	CurrentHealth   int `form:"CurrentHealth" json:"current_health"`
	MaximumCapacity int `form:"MaximumCapacity" json:"maximum_capacity"`
}

type Car struct {
	ID        int `form:"ID" json:"id"`
	BatteryID int `form:"BatteryID" json:"battery_id"`
}

type Station struct {
	Name        string `form:"Name" json:"name"`
	Address     int    `form:"Address" json:"address"`
	InOperation bool   `form:"InOperation" json:"in_operation"`
}

type Account struct {
	ID          int    `form:"ID" json:"id" gorm:"column:ID"`
	FirstName   string `form:"FirstName" json:"first_name" gorm:"column:FirstName"`
	LastName    string `form:"LastName" json:"last_name" gorm:"column:LastName"`
	Email       string `form:"Email" json:"email" gorm:"column:Email"`
	Pass        []byte `form:"Pass" json:"-" gorm:"column:Pass"`
	AccountType string `form:"AccountType" json:"account_type" gorm:"column:AccountType"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    interface{}
}
