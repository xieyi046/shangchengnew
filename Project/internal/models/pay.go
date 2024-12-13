package models

type Pay struct {
OrderId int `json:"order_id"`
OrderNo string `json:"order_no"`
ProductId int `json:"product_id"`
ProductNum int `json:"product_num"
PayMoney int `json:"product_money"`

}