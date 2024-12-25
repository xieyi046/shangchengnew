package models

//单商品购买
type Pay struct {
	PayId      int     `json:"pay_id"`      //支付ID
	OrderId    int     `json:"order_id"`    //订单I
	UserId     int     `json:"user_id"`     //用户ID
	PayTime    string  `json:"pay_time"`    //支付时间
	PayMoney   float64 `json:"paymoney"`    //支付金额
	ProductId  int     `json:"product_id"`  //商品名称
	ProductNum int     `json:"product_num"` //商品数量
}
