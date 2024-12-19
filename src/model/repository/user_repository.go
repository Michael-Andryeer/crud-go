package repository

import (
	"github.com/Michael-Andryeer/crud-go/src/configuration/rest_errors"
	"github.com/Michael-Andryeer/crud-go/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_errors.RestError)

	UpdateUser(
		userId string,
		userDomain model.UserDomainInterface,
	) *rest_errors.RestError

	DeleteUser(
		userId string,
	) *rest_errors.RestError

	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *rest_errors.RestError)

	FindUserByEmailAndPassword(
		email, password string,
	) (model.UserDomainInterface, *rest_errors.RestError)

	FindUserByID(
		id string,
	) (model.UserDomainInterface, *rest_errors.RestError)
}
