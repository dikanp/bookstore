package services

import (
	"github.com/dikanp/domain/users"
	"github.com/dikanp/utils/errors"
)


func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	// if userId <= 0 {
	// 	return nil, errors.NewBadRequestError("invalid user id")
	// }
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// func GetAllUser() ([]*users.User, *errors.RestErr) {
// 	result := &users.User{}
// 	result := result.GetAll()
// 	// if err := result.GetAll(); err != nil {
// 	// 	return nil, err
// 	// }
// 	return result, nil
// }