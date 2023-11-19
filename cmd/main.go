package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"tender/internal/handler"
	"tender/internal/repository"
	"tender/internal/service"
	"tender/pkg/server"

	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//if err := initConfig(); err != nil {
	//	logrus.Fatalf("config init error: %s", err.Error())
	//}

	db, err := repository.NewPostgresDB(repository.Config{
		//Host:     viper.GetString("db.host"),
		//Port:     viper.GetString("db.port"),
		//Username: viper.GetString("db.username"),
		//DBName:   viper.GetString("db.dbname"),
		//SSLMode:  viper.GetString("db.sslmode"),
		//Password: viper.GetString("db.password"),

		Host:     "0.0.0.0",
		Port:     "5432",
		Username: "bruh",
		DBName:   "db",
		SSLMode:  "disable",
		Password: "bruh",
		//Host:     "postgres://bruh:bruh@0.0.0.0:5432/db?sslmode=disable",
	})

	if err != nil {
		logrus.Fatalf("failed to init db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := &server.Server{}
	go func() {
		if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
			logrus.Fatalf("an error occurred while running the server: %s", err.Error())
		}
	}()

	logrus.Print("Application started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("Application finished")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("server shutting down error: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("db connection close error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
