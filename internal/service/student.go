package service

import (
"Aws-pgx/internal/repository"
"Aws-pgx/model"
)

type StudentService struct {
	repo repository.Student
}

func NewStudentService(repo repository.Student) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) CreateStudent(student model.Student) (int, error) {
	return s.repo.CreateStudent(student)
}
