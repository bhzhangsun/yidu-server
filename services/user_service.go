package services

import (
	"yidu.4se.tech/models"
)

type UserService struct {
	models.User
}

// NewUserService.
func NewUserService() *UserService {
	return &UserService{}
}

func (this *UserService) GetUserRepoById(id string) bool {
	this.UserId = id
	ok, _ := DB.Get(&this.User)
	return ok
}

func (this *UserService) Register(user *models.User) bool {

	count, _ := DB.InsertOne(*user)
	this.User = *user
	return count > 0
}
