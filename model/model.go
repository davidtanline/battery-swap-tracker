package model

type Battery struct {
	ID              int `form:"id" json:"id"`
	CurrentHealth   int `form:"current_health" json:"current_health"`
	MaximumCapacity int `form:"maximum_capacity" json:"maximum_capacity"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Battery
}
