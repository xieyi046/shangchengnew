package models

import "time"

type Order struct {
	Id         int       `json:"id"`          //主键
	UserId     int       `json:"user_id"`     //用户id
	On         string    `json:"on"`          //订单号
	Money      float64   `json:"money"`       //商品金额
	Num        int       `json:"num"`         //商品数量
	OrderPrice float64   `json:"order_price"` //订单金额
	CreateTime time.Time `json:"create_time"` //订单创建时间
	UpdateTime time.Time `json:"update_time"` //订单更新时间
	PayType    string    `json:"pay_type"`    //支付方式
	PayStatus  int       `json:"pay_status"`  //支付状态 1=未支付 2=已支付
}
