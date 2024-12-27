package services

import (
	"errors"
	"fmt"

	"github.com/shangcheng/Project/internal/dao"
	"github.com/shangcheng/Project/internal/models"
	"github.com/shangcheng/Project/types"
)

type ProductService struct {
	ProductDao *dao.ProductDao
}

// 创建商品
func (s *ProductService) AddProduct(product models.Product) (*models.Product, error) {
	if product.Name == "" {
		return nil, errors.New("商品名不能为空")
	}
	if product.Price <= 0 {
		return nil, errors.New("商品价格要大于零")
	}
	if product.Struct < 0 {
		return nil, errors.New("商品库存要大于等于零")
	}

	// 可以考虑使用事务来保证操作的一致性
	err := s.ProductDao.CreateProduct(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// 获取商品列表
func (s *ProductService) GetPaginatedProducts(page types.BasePage) ([]*models.Product, int64, error) {
	var products []*models.Product
	var total int64

	// 查询总记录数
	err := s.ProductDao.CountProducts(&total)
	if err != nil {
		return nil, 0, err
	}

	// 查询分页后的产品列表
	products, err = s.ProductDao.ListProductByCondition(nil, page)
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

// 根据ID获取商品
func (s *ProductService) GetProductById(id int) (*models.Product, error) {
	product, err := s.ProductDao.GetProductById(id)
	if err != nil {
		return nil, fmt.Errorf("获取商品时出错 %d: %w", id, err)
	}
	return product, nil
}

// 更新商品
func (s *ProductService) UpdateProduct(id int, updatedProduct models.Product) (*models.Product, error) {
	// 1. 获取商品信息
	product, err := s.ProductDao.GetProductById(id)
	if err != nil {
		return nil, errors.New("商品不存在")
	}

	// 2. 更新商品信息
	product.Name = updatedProduct.Name
	product.Price = updatedProduct.Price
	product.Struct = updatedProduct.Struct

	// 3. 更新库存
	err = s.ProductDao.UpdateProductStock(product.Id, product.Struct)
	if err != nil {
		return nil, err
	}

	// 4. 返回更新后的商品信息
	return product, nil
}

// 删除商品
func (s *ProductService) DeleteProduct(id int) error {

	err := s.ProductDao.DeleteProduct(id)
	if err != nil {
		return err
	}

	return nil
}
