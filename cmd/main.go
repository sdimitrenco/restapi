package main

import (
	"github.com/StanislavDimitrenco/restapi/pkg/handler"
	"github.com/StanislavDimitrenco/restapi/pkg/repository"
	"github.com/StanislavDimitrenco/restapi/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"

	"github.com/StanislavDimitrenco/restapi"
	"github.com/joho/godotenv"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Can't read config file - %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Can't load env config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("filed to initiate db: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)

	if err := server.Run(viper.GetString("port"), handlers.InitRouts()); err != nil {
		logrus.Fatal("error")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
