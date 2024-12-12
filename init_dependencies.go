package main

import (
	"github.com/Michael-Andryeer/crud-go/src/controller"
	"github.com/Michael-Andryeer/crud-go/src/model/repository"
	"github.com/Michael-Andryeer/crud-go/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
