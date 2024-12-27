package services

import (
	"errors"
	"log"
	"time"

	"github.com/shangcheng/Project/internal/dao"
	"github.com/shangcheng/Project/internal/models"
	jwt "github.com/shangcheng/Project/pkg/util/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserDao *dao.UserDao
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) CreateUser(user *models.User) error {
	if user.UserName == "" || user.Phone == "" {
		return errors.New("所有字段都是必填项")
	}

	if len(user.PassWord) < 6 {
		return errors.New("密码不能少于6位")
	}

	// 对密码进行哈希处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PassWord), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("失败")
	}
	user.PassWord = string(hashedPassword)

	user.Money = 0.0
	user.StartTime = time.Now()

	if err := s.UserDao.CreateUser(user); err != nil {
		return err
	}

	log.Printf("用户 %s 创建成功", user.UserName)
	return nil
}

// 登录
func (s *UserService) Login(userName, password string) (string, string, error) {
	// 验证用户名和密码不能为空
	if userName == "" || password == "" {
		return "", "", errors.New("用户名和密码不能为空")
	}

	// 查询用户
	user, err := s.UserDao.GetUserByUsername(userName)
	if err != nil {
		return "", "", errors.New("用户不存在或登录失败")
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password)); err != nil {
		return "", "", errors.New("密码错误")
	}

	// 登录成功，生成 Token
	accessToken, refreshToken, err := jwt.GenerateToken(user.Id, user.UserName)
	if err != nil {
		return "", "", err
	}

	log.Printf("用户 %s 登录成功", userName)
	return accessToken, refreshToken, nil
}

// 更新信息
func (s *UserService) UpdateUser(user *models.User) error {
	// 验证数据完整性
	if user.UserName == "" || user.Phone == "" {
		return errors.New("用户名和手机号不能为空")
	}

	if len(user.PassWord) < 6 && user.PassWord != "" {
		return errors.New("密码不能少于6位")
	}

	// 如果提供了新密码，则对密码进行哈希处理
	if user.PassWord != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PassWord), bcrypt.DefaultCost)
		if err != nil {
			return errors.New("密码哈希失败")
		}
		user.PassWord = string(hashedPassword)
	}

	// 调用 DAO 层进行更新
	if err := s.UserDao.UpdateUser(user); err != nil {
		return err
	}

	log.Printf("用户 %s 信息更新成功", user.UserName)
	return nil
}
