package dao

import (
	"errors"
	"fmt"

	"github.com/shangcheng/Project/internal/models"
	"github.com/shangcheng/Project/types"
	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

// 根据id获取商品
func (dao *ProductDao) GetProductById(id int) (*models.Product, error) {
	var product models.Product
	result := dao.DB.Model(&models.Product{}).Where("id = ?", id).First(&product)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("查找到商品出错: %d", id)
		}
		return nil, fmt.Errorf("无法获取该商品id的商品 %d: %v", id, result.Error)
	}
	return &product, nil
}

// 获取商品列表
func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, page types.BasePage) ([]*models.Product, error) {
	var products []*models.Product

	query := dao.DB.Where(condition).
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize)

	result := query.Find(&products)
	if result.Error != nil {
		return nil, fmt.Errorf("获取有误: %v", result.Error)
	}

	return products, nil
}

// 获取商品总数
func (dao *ProductDao) CountProducts(count *int64) error {
	return dao.DB.Model(&models.Product{}).Count(count).Error
}

// 创建商品
func (dao *ProductDao) CreateProduct(product *models.Product) error {
	return dao.DB.Model(&models.Product{}).
		Create(&product).Error
}

// 删除商品
func (dao *ProductDao) DeleteProduct(id int) error {
	result := dao.DB.Delete(&models.Product{}, id)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("未找到该商品")
	}
	return nil
}

// 更新商品
func (dao *ProductDao) UpdateProductStock(productId int, quantity int) error {
	// 获取商品信息
	var product models.Product
	if err := dao.DB.Model(&models.Product{}).Where("id = ?", productId).First(&product).Error; err != nil {
		return errors.New("商品不存在")
	}

	// 判断库存是否足够
	if product.Struct < quantity {
		return errors.New("库存不足")
	}

	// 更新库存
	newStock := product.Struct - quantity
	if err := dao.DB.Model(&models.Product{}).Where("id = ?", productId).Update("struct", newStock).Error; err != nil {
		return err
	}

	return nil
}

// 搜索
func (dao *ProductDao) SearchProduct(info string, page types.BasePage) (products []*models.Product, count int64, err error) {
	err = dao.DB.Model(&models.Product{}).
		Where("name LIKE ? OR info LIKE ?", "%"+info+"%", "%"+info+"%").
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).
		Find(&products).Error

	if err != nil {
		return
	}
	return
}
