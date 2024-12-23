package services

import (
	"errors"

	"github.com/shangcheng/Project/internal/consts"
	"github.com/shangcheng/Project/internal/dao"
	"github.com/shangcheng/Project/internal/models"
	"gorm.io/gorm"
)

type PayService struct {
	OrderDao   *dao.OrderDao
	UserDao    *dao.UserDao
	ProductDao *dao.ProductDao
}

func NewPayService(orderDao *dao.OrderDao, userDao *dao.UserDao, productDao *dao.ProductDao) *PayService {
	return &PayService{
		OrderDao:   orderDao,
		UserDao:    userDao,
		ProductDao: productDao,
	}
}

func (s *PayService) PayOrder(pay models.Pay) error {
	// 1. 查询订单是否存在
	order, err := s.OrderDao.GetOrderById(pay.OrderId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("订单不存在")
		}
	}
	// 2. 判断订单状态是否为未支付
	if order.PayStatus != consts.OrderStatusPend {
		return errors.New("订单已支付")
	}
	// 3. 计算单个购买情况下的金额
	money := order.Money
	num := order.Num
	paymoney := money * float64(num)
	// 4.获取用户余额
	user, err := s.UserDao.GetUserByID(pay.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("没有查询到该用户")
		}
	}
	// 5. 判断余额是否足够支付
	if user.Money-paymoney < 0.0 {
		return errors.New("余额不足")
	}
	// 6. 更新订单状态
	order.PayStatus = consts.OrderStatuspay
	err = s.OrderDao.UpdateOrderById(uint(pay.OrderId), uint(pay.UserId), order)
	if err != nil {
		return err
	}
	// 7. 更新用户余额
	user.Money = user.Money - paymoney
	err = s.UserDao.UpdateUserById(uint(pay.UserId), user)
	if err != nil {
		return err
	}
	// 8. 库存更新
	err = s.ProductDao.UpdateProductStock(pay.ProductId, order.Num)
	if err != nil {
		return err
	}
	return nil
}
