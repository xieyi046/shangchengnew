package services

import (
	"errors"
	"time"

	"github.com/shangcheng/Project/Project/internal/dao"
	"github.com/shangcheng/Project/Project/internal/models"
	"gorm.io/gorm"
)

type OrderService struct {
	OrderDao *dao.OrderDao
}

// 创建订单
func (s *OrderService) CreateOrder(userId int, orderPrice float64, payType string) (*models.Order, error) {
	if orderPrice <= 0 {
		return nil, errors.New("订单价格必须大于零")
	}

	order := &models.Order{
		UserId:     userId,
		On:         generateOrderNumber(),
		OrderPrice: orderPrice,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		PayType:    payType,
		PayStatus:  "1", // 未支付
	}

	err := s.OrderDao.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// 查询订单
func (s *OrderService) GetOrderById(orderId int) (*models.Order, error) {
	order, err := s.OrderDao.GetOrderById(orderId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("订单不存在")
		}
		return nil, err
	}

	return order, nil
}

// 删除订单
func (s *OrderService) DeleteOrder(orderId, userId int) error {
	return s.OrderDao.DeleteOrderById(uint(orderId), uint(userId))
}

// 生成订单号
func generateOrderNumber() string {
	return "ORDER" + time.Now().Format("20060102150405")
}
