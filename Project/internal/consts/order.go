package consts

const (
	OrderStatusPend = iota + 1
	OrderStatuspay
)

var OrderStatusText = map[int]string{
	OrderStatusPend: "未支付",
	OrderStatuspay:  "已支付",
}
