package service

import (
	"Aws-pgx/internal/repository"
	"Aws-pgx/model"
)

type Class interface {
	CreateClass(class model.Class) (int, error)
	GetAllClasses() ([]model.Class, error)
}

type Student interface {
	CreateStudent(student model.Student) (int, error)
}

type Service struct {
	Class
	Student
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Class: NewClassService(repos.Class),
		Student: NewStudentService(repos.Student),
	}
}