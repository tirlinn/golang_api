package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

const (
	classTable = "classes"
	studentTable = "students"
	classStudentTable = "class_students"
)

type Config struct {
	Host string
	Port string
	Username string
	Password string
	DBName string
}

func NewPostgresDB(cfg Config) (*pgxpool.Pool, error) {

	dbUrl := fmt.Sprintf("%s://%s:%s@%s:%s/%s", cfg.DBName, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	dbPool, err := pgxpool.Connect(context.Background(), dbUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	err = dbPool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return dbPool, nil
}
