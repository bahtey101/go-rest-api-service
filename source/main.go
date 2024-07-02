package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/bahtey101/go-rest-api-service/package/handler"
	"github.com/bahtey101/go-rest-api-service/package/repository"
	"github.com/bahtey101/go-rest-api-service/package/service"
	"github.com/bahtey101/go-rest-api-service/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	repos := repository.NewRepository(viper.GetString("path"))
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	server := new(server.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
