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

func (r *ClassPostgres) GetAllClasses() ([]model.Class, error) {
	var list []model.Class

	query := fmt.Sprintf("SELECT * FROM %s", classTable)

	rows, err := r.dbPool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var class model.Class

		err = rows.Scan(&class.Id, &class.ClassName)
		if err != nil {
			return nil, err
		}

		list = append(list, class)
	}

	return list, err
}