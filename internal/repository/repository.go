package repository

import (
	"Aws-pgx/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Class interface {
	CreateClass(class model.Class) (int, error)
	GetAllClasses() ([]model.Class, error)
}

type Student interface {
	CreateStudent(student model.Student) (int, error)
}

type Repository struct {
	Class
	Student
}

func NewRepository(dbPool *pgxpool.Pool) *Repository {
	return &Repository{
		Class: NewClassPostgres(dbPool),
		Student: NewStudentPostgres(dbPool),
	}
}

//func createConnection() *pgxpool.Pool {
//	databaseUrl := "postgres://postgres_go:golang2021@database-2.c7sjcxek3mor.eu-central-1.rds.amazonaws.com:5432/postgres"
//	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)
//
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Failed connection")
//		os.Exit(1)
//	}
//
//	fmt.Fprintf(os.Stdout, "Connection done\n")
//
//	return dbPool
//}
//
//func InsertStudentQuery(s *model.Student) (string, error) {
//	dbPool := createConnection()
//	defer dbPool.Close()
//
//	id, err := dbPool.Exec(context.Background(), "insert into students (first_name, last_name, gender, class, status) values ($1, $2, $3, $4, $5) returning id", s.FirstName, s.LastName, s.Gender, s.Class, s.Status);
//
//	if  err != nil {
//		fmt.Println("Unable to insert due to: ", err)
//		return "", err
//	}
//	fmt.Println("Insert successful")
//
//	return string(id), nil
//}
//
//func SelectStudentQuery() (pgx.Rows, error) {
//	dbPool := createConnection()
//	defer dbPool.Close()
//
//	rows, err := dbPool.Query(context.Background(), "Select * from students")
//	if err != nil {
//		log.Fatal("error while executing query")
//		return pgx.Rows{} , err
//	}
//	return rows, err
//}