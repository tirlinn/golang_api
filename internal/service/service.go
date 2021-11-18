package service

import (
	"Aws-pgx/internal/repository"
	"Aws-pgx/model"
)

type Class interface {
	CreateClass(class model.Class) (int, error)
}

type Student interface {

}

type Service struct {
	Class
	Student
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Class: NewClassService(repos.Class),
	}
}