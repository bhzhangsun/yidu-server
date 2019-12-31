package services

import (
	"mobi.4se.tech/models"
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
	ok, _ := db.Get(&this.User)
	return ok
}

func (this *UserService) Register(user *models.User) bool {

	count, _ := db.InsertOne(*user)
	this.User = *user
	return count > 0
}
