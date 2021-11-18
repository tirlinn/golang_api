package model

type Student struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Gender    string `json:"gender" form:"gender"`
	Class     string `json:"class" form:"class"`
	Status    bool   `json:"status" form:"status"`
}