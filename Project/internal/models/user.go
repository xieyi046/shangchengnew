package models

import (
	"errors"
	"regexp"
	"time"
)

type User struct {
	Id        int       `json:"id"`         //主键
	UserName  string    `json:"username"`   //账号
	PassWord  string    `json:"password"`   //密码
	Phone     string    `json:"phone"`      //手机号
	StartTime time.Time `json:"start_time"` //账号创建时间
	Momey     int       `json:"momey"`      //余额
}

// 验证注册信息
func (u *User) Validate() error {
	if u.UserName == "" {
		return errors.New("用户名不能为空")
	}

	if len(u.PassWord) < 6 {
		return errors.New("密码长度必须至少为6个字符")
	}

	if u.Phone == "" || !isValidPhone(u.Phone) {
		return errors.New("无效的手机号")
	}

	return nil
}

// 验证手机号
func isValidPhone(phone string) bool {
	re := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return re.MatchString(phone)
}
