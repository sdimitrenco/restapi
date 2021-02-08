package main

import (
	"github.com/StanislavDimitrenco/restapi/pkg/handler"
	"github.com/StanislavDimitrenco/restapi/pkg/repository"
	"github.com/StanislavDimitrenco/restapi/pkg/service"
	"github.com/spf13/viper"
	"log"
	"os"

	"github.com/StanislavDimitrenco/restapi"
	"github.com/joho/godotenv"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("Can't read config file - %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Can't load env config: %s", err.Error())
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
		log.Fatalf("filed to initiate db: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)

	if err := server.Run(viper.GetString("port"), handlers.InitRouts()); err != nil {
		log.Fatal("error")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
