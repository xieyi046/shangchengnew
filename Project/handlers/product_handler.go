package handlers

import (
	"github.com/shangcheng/Project/Project/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shangcheng/Project/Project/internal/models"
	"github.com/shangcheng/Project/Project/internal/services"
)

// ProductHandler 处理产品的请求
type ProductHandler struct {
	ProductService *services.ProductService
}

// 添加产品
func (h *ProductHandler) AddProduct(c *gin.Context) {
	var req models.Product
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newProduct, err := h.ProductService.AddProduct(req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newProduct)
}

// 获取所有产品
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	// 使用 types.BasePage 从请求中获取分页参数，默认值为第1页，每页10条记录
	var page types.BasePage
	if err := c.ShouldBindQuery(&page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	if page.PageNum <= 0 {
		page.PageNum = 1
	}
	if page.PageSize <= 0 || page.PageSize > 100 { // 限制最大页面大小
		page.PageSize = 10
	}

	products, total, err := h.ProductService.GetPaginatedProducts(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 返回分页结果
	response := map[string]interface{}{
		"total":    total,
		"items":    products,
		"page":     page.PageNum,
		"pageSize": page.PageSize,
	}
	c.JSON(http.StatusOK, response)
}

// 获取单个产品
func (h *ProductHandler) GetProductById(c *gin.Context) {
	idStr := c.DefaultQuery("id", "")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}

	product, err := h.ProductService.GetProductById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// 更新产品信息
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.DefaultQuery("id", "")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}

	var req models.Product
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	product, err := h.ProductService.UpdateProduct(id, req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// 删除产品
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.DefaultQuery("id", "")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}

	err = h.ProductService.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
