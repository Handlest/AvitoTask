package main

import (
	avito "AvitoTask"
	"AvitoTask/pkg/handler"
	"AvitoTask/pkg/repository"
	service2 "AvitoTask/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

// @title Avito Trainee Application
// @version 1.0
// @description API Application task for backend trainee position

// @host localhost:8000
// @BasePath /

func main() {

	//Считали логи из файла, сохраняя их в viper
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("failed to start db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service2.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(avito.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running server: %s", err.Error())
	}
}

// Считали файлы конфигураций из папки configs для дальнейшего их использования
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
