package repository

import (
	"context"
	"os"

	"github.com/Michael-Andryeer/crud-go/src/configuration/logger"
	"github.com/Michael-Andryeer/crud-go/src/configuration/rest_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(
	userId string,
) *rest_errors.RestError {
	logger.Info("Init deleteUser repository",
		zap.String("journey", "deleteUser"),
	)

	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error trying to  DeleteUser",
			err,
			zap.String("journey", "deleteUser"),
		)
		return rest_errors.NewInternalServerError(err.Error())
	}

	logger.Info("DeleteUser repository executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "DeleteUser"),
	)

	return nil
}
