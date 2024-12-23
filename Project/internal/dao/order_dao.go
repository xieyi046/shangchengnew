package dao

import (
	"errors"

	"github.com/shangcheng/Project/internal/models"
	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

// 创建订单
func (dao *OrderDao) CreateOrder(order *models.Order) error {
	return dao.DB.Create(&order).Error
}

// 更新订单详情
func (dao *OrderDao) UpdateOrderById(id, uId uint, order *models.Order) error {
	return dao.DB.Where("id = ? AND user_id = ?", id, uId).
		Updates(order).Error
}

// 删除订单
func (dao *OrderDao) DeleteOrderById(id, uId uint) error {
	return dao.DB.Model(&models.Order{}).
		Where("id = ? AND user_id = ?", id, uId).
		Delete(&models.Order{}).Error
}

// 查询订单
func (dao *OrderDao) GetOrderById(orderId int) (*models.Order, error) {
	var order models.Order

	err := dao.DB.Model(&models.Order{}).Where("id = ?", orderId).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("订单不存在")
		}
		return nil, err
	}

	return &order, nil
}

func (dao *OrderDao) GetOrdersByUserId(userId, offset, pageSize int, orders *[]*models.Order) error {
	// 使用分页查询
	err := dao.DB.Model(&models.Order{}).
		Where("user_id = ?", userId).
		Offset(offset).
		Limit(pageSize).
		Find(orders).Error
	if err != nil {
		return err
	}

	return nil
}
