package converter

import (
	"github.com/Michael-Andryeer/crud-go/src/model"
	"github.com/Michael-Andryeer/crud-go/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
		Password: domain.GetPassword(),
	}
}
