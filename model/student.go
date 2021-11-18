package model

type Student struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Class     string `json:"class"`
	Status    bool   `json:"status"`
}