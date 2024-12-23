package models

type MultiPay struct {
	UserId int       `json:"user_id" binding:"required"`
	Orders []PayItem `json:"orders" binding:"required,dive"`
}

type PayItem struct {
	OrderId   int     `json:"order_id" binding:"required"`
	ProductId int     `json:"product_id" binding:"required"`
	Money     float64 `json:"money" binding:"required"`
	Num       int     `json:"num" binding:"required"`
}
