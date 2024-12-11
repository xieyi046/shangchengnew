package services

import (
	"errors"
	"github.com/shangcheng/Project/Project/internal/models"
	"github.com/shangcheng/Project/Project/internal/dao"
	"github.com/shangcheng/Project/Project/types"
)

type ProductService struct {
	ProductDao *dao.ProductDao
}

// 创建商品
func (s *ProductService) AddProduct(product models.Product) (*models.Product, error) {

	if product.Name == "" {
		return nil, errors.New("product name is required")
	}

	err := s.ProductDao.CreateProduct(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// 获取所有商品
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
	product, err := s.ProductDao.GetProductById(uint(id))
	if err != nil {
		return nil, err
	}
	return product, nil
}

// 更新商品
func (s *ProductService) UpdateProduct(id int, updatedProduct models.Product) (*models.Product, error) {
	product, err := s.ProductDao.GetProductById(uint(id))
	if err != nil {
		return nil, errors.New("product not found")
	}

	product.Name = updatedProduct.Name
	product.Price = updatedProduct.Price
	product.Struct = updatedProduct.Struct

	err = s.ProductDao.UpdateProduct(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// 删除商品
func (s *ProductService) DeleteProduct(id int) error {

	err := s.ProductDao.DeleteProduct(uint(id))
	if err != nil {
		return err
	}

	return nil
}
