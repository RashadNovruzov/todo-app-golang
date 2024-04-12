package main

import (
	"os"

	"github.com/RoshkANovruzov/todo-app-golang"
	"github.com/RoshkANovruzov/todo-app-golang/pkg/handler"
	"github.com/RoshkANovruzov/todo-app-golang/pkg/repository"
	"github.com/RoshkANovruzov/todo-app-golang/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// open a file
	// f, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	// if err != nil {
	// 	fmt.Printf("error opening file: %v", err)
	// }

	// // don't forget to close it
	// defer f.Close()

	// // assign it to the standard logger
	// logrus.SetOutput(f)

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	if err := initConfig(); err != nil {
		logrus.Fatalf("error occured while reading config file %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
		DBName:   viper.GetString("db.dbname"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(todo.Server)
	if err := server.Run(viper.GetString("app.port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
