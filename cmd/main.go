package main

import (
	lesson "Aws-pgx"
	"Aws-pgx/internal/handler"
	"Aws-pgx/internal/repository"
	"Aws-pgx/internal/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error initializing env variables: %s\n", err.Error())
	}

	dbPool, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: viper.GetString("db.name"),
	})

	defer dbPool.Close()

	repos := repository.NewRepository(dbPool)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)


	if err != nil {
		log.Fatalf("Failed to initialize db: %s\n", err.Error())
	}

	srv := lesson.NewServer()

	handlers.InitRoutes(srv.EchoServer)

	if err := srv.Run(viper.GetString("server.port")); err != nil {
		log.Fatalf("Error running server %s", err);
	}

	os.Exit(0)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

//CreateClassTableQuery(dbPool)
//CreateStudentsTableQuery(dbPool)
//CreateClassesStudentsTableQuery(dbPool)
//CreateTimetableTableQuery(dbPool)
//
//func CreateClassTableQuery(dbPool *pgxpool.Pool) {
//	var createQuery string = "create table if not exists class (" +
//		"id SERIAL," +
//		"class_name VARCHAR," +
//		"primary key(id)" +
//		")"
//
//	_, err := dbPool.Query(context.Background(), createQuery)
//	if err != nil {
//		log.Fatal("error while executing query")
//	}
//	fmt.Fprintf(os.Stdout, "Class query done\n")
//}
//
//func CreateStudentsTableQuery(dbPool *pgxpool.Pool) {
//	var createQuery string = "create table if not exists students (" +
//		"id SERIAL," +
//		"first_name VARCHAR," +
//		"last_name VARCHAR," +
//		"gender VARCHAR," +
//		"class VARCHAR," +
//		"status BOOL," +
//		"primary key(id)" +
//		")"
//
//	_, err := dbPool.Query(context.Background(), createQuery)
//	if err != nil {
//		log.Fatal("error while executing query")
//	}
//	fmt.Fprintf(os.Stdout, "Students query done\n")
//}
//
//func CreateClassesStudentsTableQuery(dbPool *pgxpool.Pool) {
//	var createQuery string = "create table if not exists class_students (" +
//		"class_id INTEGER," +
//		"student_id INTEGER," +
//		"primary key (class_id, student_id)," +
//		"foreign key (class_id) references classes (id)," +
//		"foreign key (student_id) references students (id)" +
//		")"
//
//	_, err := dbPool.Query(context.Background(), createQuery)
//	if err != nil {
//		log.Fatal("error while executing query")
//	}
//	fmt.Fprintf(os.Stdout, "Classes_students query done\n")
//}
//
//func CreateSubjectTableQuery(dbPool *pgxpool.Pool) {
//	var createQuery string = "create table if not exists subjects (" +
//		"id SERIAL," +
//		"name VARCHAR," +
//		"primary key(id)" +
//		")"
//
//	_, err := dbPool.Query(context.Background(), createQuery)
//	if err != nil {
//		log.Fatal("error while executing query")
//	}
//	fmt.Fprintf(os.Stdout, "Class query done\n")
//}
//
//func CreateTimetableTableQuery(dbPool *pgxpool.Pool) {
//	var createQuery string = "create table if not exists timetable (" +
//		"subject_id INTEGER," +
//		"class_id INTEGER," +
//		"time TIMESTAMP," +
//		"primary key (subject_id, class_id)," +
//		"foreign key (subject_id) references subject (id)" +
//		"foreign key (class_id) references classes (id)" +
//		")"
//
//	_, err := dbPool.Query(context.Background(), createQuery)
//	if err != nil {
//		log.Fatal("error while executing query")
//	}
//	fmt.Fprintf(os.Stdout, "Timetable query done\n")
//}
