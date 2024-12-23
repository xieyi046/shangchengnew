package services

import (
	"errors"

	"github.com/shangcheng/Project/internal/consts"
	"github.com/shangcheng/Project/internal/dao"
	"github.com/shangcheng/Project/internal/models"
	"gorm.io/gorm"
)

type PaysService struct {
	OrderDao   *dao.OrderDao
	UserDao    *dao.UserDao
	ProductDao *dao.ProductDao
}

func NewPaysService(orderDao *dao.OrderDao, userDao *dao.UserDao, productDao *dao.ProductDao) *PaysService {
	return &PaysService{
		OrderDao:   orderDao,
		UserDao:    userDao,
		ProductDao: productDao,
	}
}

func (s *PaysService) MultiPayOrder(pay models.MultiPay) error {
	// 1. 查询用户信息
	user, err := s.UserDao.GetUserByID(pay.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("用户不存在")
		}
		return err
	}

	// 2. 计算总支付金额
	totalPayAmount := 0.0
	for _, payItem := range pay.Orders {
		// 3. 查询订单信息
		order, err := s.OrderDao.GetOrderById(payItem.OrderId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("订单不存在")
			}
			return err
		}

		// 4. 判断订单是否已支付
		if order.PayStatus != consts.OrderStatusPend {
			return errors.New("订单已支付")
		}

		// 5. 计算该订单的支付金额
		payAmount := payItem.Money * float64(payItem.Num)
		totalPayAmount += payAmount

		// 6. 更新订单状态为已支付
		order.PayStatus = consts.OrderStatuspay
		if err := s.OrderDao.UpdateOrderById(uint(payItem.OrderId), uint(pay.UserId), order); err != nil {
			return err
		}

		// 7. 更新库存
		if err := s.ProductDao.UpdateProductStock(payItem.ProductId, payItem.Num); err != nil {
			return err
		}
	}

	// 8. 检查余额是否足够
	if user.Money < totalPayAmount {
		return errors.New("余额不足")
	}

	// 9. 扣除用户余额
	user.Money -= totalPayAmount
	if err := s.UserDao.UpdateUserById(uint(pay.UserId), user); err != nil {
		return err
	}

	return nil
}
