package dao

import (
	"errors"
	"fmt"

	"github.com/shangcheng/Project/internal/config"
	"github.com/shangcheng/Project/internal/models"

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
	var user *models.User
	result := dao.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return user, nil
}

// 更新用户余额
func (dao *UserDao) UpdateUserById(id uint, user *models.User) error {
	// 首先检查用户是否存在
	var existingUser models.User
	result := dao.DB.First(&existingUser, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("用户不存在")
		}
		return fmt.Errorf("获取用户信息失败: %v", result.Error)
	}

	updateData := map[string]interface{}{
		"money": user.Money,
	}
	// 使用 GORM 的 Model 和 Updates 方法进行更新
	result = dao.DB.Model(&existingUser).Updates(updateData)
	if result.Error != nil {
		return fmt.Errorf("更新用户信息失败: %v", result.Error)
	}
	return nil
}

// 更新用户身份信息
func (dao *UserDao) UpdateUser(user *models.User) error {
	// 更新数据库中的用户信息
	result := dao.DB.Exec("UPDATE users SET username = ?, password = ?, phone = ? WHERE id = ?", user.UserName, user.PassWord, user.Phone, user.Id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
