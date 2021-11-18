package repository

import (
	"Aws-pgx/model"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ClassPostgres struct {
	dbPool *pgxpool.Pool
}

func NewClassPostgres(dbPool *pgxpool.Pool) *ClassPostgres {
	return &ClassPostgres{
		dbPool: dbPool,
	}
}

func (r *ClassPostgres) CreateClass(class model.Class) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (class_name) VALUES ($1) RETURNING id", classTable)
	row := r.dbPool.QueryRow(context.Background(), query, class.ClassName)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
