package main

import (
	_ "github.com/TheKruasan/CRM-server/docs"
	"github.com/TheKruasan/CRM-server/pkg/handler"
	"github.com/TheKruasan/CRM-server/pkg/repository"
	"github.com/TheKruasan/CRM-server/pkg/service"
	"github.com/TheKruasan/CRM-server/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//@title CRM App API

//@version 1.0
//@description API Server for CRM System Application

// @host localhost:8000
// @BasePath /

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err)
	}
	// db, err := repository.NewPostgresDb(repository.Config{
	// 	Host:     viper.GetString("db.host"),
	// 	Port:     viper.GetString("db.pqport"),
	// 	Username: viper.GetString("db.username"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	Database: viper.GetString("db.database"),
	// 	SSLMode:  viper.GetString("db.sslmode"),
	// })
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
	if err := initConfig(); err != nil {
		logrus.Fatal(err)
	}
	// repos := repository.NewRepository(db)
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)
	server := new(server.Server)
	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatal(err)
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
