package models

type Product struct {
	Id     int     `json:"id"`     //主键
	Name   string  `json:"name"`   //商品名称
	Price  float64 `json:"price"`  //价格
	Struct string  `json:"struct"` //库存
}
