package dao

import (

	"github.com/shangcheng/Project/Project/internal/config"
	"github.com/shangcheng/Project/Project/internal/models"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDAO() *UserDao {
	return &UserDao{
		DB: config.DB,
	}
}

// 创建用户
func (dao *UserDao) CreateUser(user *models.User) error {
	result := dao.DB.Create(user)
	return result.Error
}

// 根据ID获取用户
func (dao *UserDao) GetUserByID(id int) (*models.User, error) {
	var user models.User
	result := dao.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// 根据用户名获取用户
func (dao *UserDao) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := dao.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// 获取用户详细信息
func (dao *UserDao) GetUserDetails(id int) (*models.User, error) {
	var user models.User
	result := dao.DB.Where("user_id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

//获取用户余额
