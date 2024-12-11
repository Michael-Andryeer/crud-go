package main

import (
	"context"
	"log"

	"github.com/Michael-Andryeer/crud-go/src/configuration/database/mongodb"
	"github.com/Michael-Andryeer/crud-go/src/configuration/logger"
	"github.com/Michael-Andryeer/crud-go/src/controller"
	"github.com/Michael-Andryeer/crud-go/src/controller/routes"
	"github.com/Michael-Andryeer/crud-go/src/model/repository"
	"github.com/Michael-Andryeer/crud-go/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting the application...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error connecting to database: %s \n", err.Error())
	}
	//Inicializar dependencias
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
