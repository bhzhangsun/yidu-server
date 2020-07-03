package services

import "yidu.4se.cool/models"

type User models.User

func (u *User) Register() error {
	_, err := DB.InsertOne(u)
	return err
}

func (u *User) Login() error {
	var result = User{
		Phone: u.Phone,
	}
	if _, err := DB.Get(&result); err != nil {
		return u.Register()
	}
	*u = result
	return nil
}

func (u *User) ChangeName(name string) error {
	u.Nickname = name
	if _, err := DB.Update(u, User{ID: u.ID}); err != nil {
		return err
	}
	return nil
}
