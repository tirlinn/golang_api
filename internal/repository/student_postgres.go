package repository

import (
	"Aws-pgx/model"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type StudentPostgres struct {
	dbPool *pgxpool.Pool
}

func NewStudentPostgres(dbPool *pgxpool.Pool) *StudentPostgres {
	return &StudentPostgres{
		dbPool: dbPool,
	}
}

func (r *StudentPostgres) CreateStudent(student model.Student) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (first_name, last_name, gender, class, status) VALUES ($1, $2, $3, $4, $5) RETURNING id", studentTable)

	row := r.dbPool.QueryRow(context.Background(), query, student.FirstName, student.LastName, student.Gender, student.Class, student.Status)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}