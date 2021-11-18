package service

import (
	"Aws-pgx/internal/repository"
	"Aws-pgx/model"
)

type ClassService struct {
	repo repository.Class
}

func NewClassService(repo repository.Class) *ClassService {
	return &ClassService{repo: repo}
}

func (s *ClassService) CreateClass(class model.Class) (int, error) {
	return s.repo.CreateClass(class)
}
