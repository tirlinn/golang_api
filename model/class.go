package model

type Class struct {
	Id int `json:"id" form:"class_name"`
	ClassName string `json:"class_name" form:"class_name"`
}

type ClassStudent struct {
	ClassId int `json:"class_id"`
	StudentId int `json:"student_id"`
}