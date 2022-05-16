package main

import (
	"github.com/joho/godotenv"
	"github.com/krinya32/pizzaApp"
	"github.com/krinya32/pizzaApp/pkg/handlers"
	"github.com/krinya32/pizzaApp/pkg/repository"
	"github.com/krinya32/pizzaApp/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := InitConfig(); err != nil {
		logrus.Fatalf("Error initializing configs %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("Failed to initialize DB %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handlers.NewHandler(services)

	srv := new(pizzaApp.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured while running http server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
