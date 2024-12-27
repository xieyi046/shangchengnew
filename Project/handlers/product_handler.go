package handlers

import (
	"net/http"
	"strconv"

	"github.com/shangcheng/Project/types"

	"github.com/gin-gonic/gin"
	"github.com/shangcheng/Project/internal/models"
	"github.com/shangcheng/Project/internal/services"
)

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

// 获取产品列表
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	var page types.BasePage
	if err := c.ShouldBindQuery(&page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数有误"})
		return
	}

	if page.PageNum <= 0 {
		page.PageNum = 1
	}
	if page.PageSize <= 0 || page.PageSize > 100 {
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

// 获取产品详情
func (h *ProductHandler) GetProductById(c *gin.Context) {
	idStr := c.DefaultQuery("id", "")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id不能不填"})
		return
	}

	// 转换ID为整数并检查是否出错
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id无效"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID无效"})
		return
	}

	var req models.Product
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数有错"})
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
