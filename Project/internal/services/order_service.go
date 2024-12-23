package services

import (
	"errors"
	"time"

	"github.com/shangcheng/Project/internal/dao"
	"github.com/shangcheng/Project/internal/models"
	"gorm.io/gorm"
)

type OrderService struct {
	OrderDao *dao.OrderDao
}

// 创建订单
func (s *OrderService) CreateOrder(order *models.Order) error {
	if order.Money <= 0 || order.Num <= 0 || order.OrderPrice <= 0 {
		return errors.New("金额、数量和订单价格必须大于零")
	}

	order.On = generateOrderNumber() // 生成订单号
	order.CreateTime = time.Now()
	order.UpdateTime = time.Now()
	order.PayStatus = 1 // 默认未支付状态

	err := s.OrderDao.CreateOrder(order)
	if err != nil {
		return err
	}

	return nil
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

// 获取订单列表
func (s *OrderService) GetOrdersByUserId(userId, page, pageSize int) ([]*models.Order, error) {
	var orders []*models.Order
	offset := (page - 1) * pageSize

	// 根据用户ID和分页查询订单
	err := s.OrderDao.GetOrdersByUserId(userId, offset, pageSize, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// 生成订单号
func generateOrderNumber() string {
	return "ORDER" + time.Now().Format("20060102150405")
}
