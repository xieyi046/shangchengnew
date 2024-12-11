package services

import (
	"errors"

	"github.com/shangcheng/Project/Project/internal/dao"
	"github.com/shangcheng/Project/Project/internal/models"
)

type UserService struct{
	UserDao *dao.UserDao
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) CreateUser(user *models.User) error {
	if user.UserName == "" || user.Phone == "" {
		return errors.New("all fields are required")
	}

	if len(user.PassWord) < 6 {
		return errors.New("密码不能少于6位")
	}

	if err := s.UserDao.CreateUser(user); err != nil {
		return err
	}
	return nil
}