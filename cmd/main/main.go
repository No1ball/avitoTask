package main

import (
	"github.com/No1ball/avitoTask/internal/config"
	"github.com/No1ball/avitoTask/internal/handlers"
	"github.com/No1ball/avitoTask/internal/repository"
	"github.com/No1ball/avitoTask/internal/server"
	"github.com/No1ball/avitoTask/internal/services"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("error init: %s", err.Error())
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
		logrus.Fatalf("error with db connection: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	service := services.NewService(repos)
	handler := handlers.NewHandler(service)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRout()); err != nil {
		logrus.Fatalf("error occurated while running http server: %s", err.Error())
	}
}
