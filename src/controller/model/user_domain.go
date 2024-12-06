package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/Michael-Andryeer/crud-go/src/configuration/rest_errors"
)

func NewUserDomain(
	email, password, name string,
	age int8,
) *UserDomain {
	return &UserDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	} //Construtor
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncryptPassowrd() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser(UserDomain) *rest_errors.RestError
	UpdateUser(string) *rest_errors.RestError
	DeleteUser(string) *rest_errors.RestError
	FindUser(string) *rest_errors.RestError
}
