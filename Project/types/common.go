package types

type BasePage struct {
	PageNum  int `form:"page_num"`
	PageSize int `form:"page_size"`
}

type OrderListReq struct {
	Page      int  `json:"page" binding:"gte=1"`              // 当前页码，默认为1
	PageSize  int  `json:"page_size" binding:"gte=1,lte=100"` // 每页显示的数量，默认为10，最大不超过100
	ProductID uint `json:"product_id,omitempty"`              // 可选：产品ID，用于筛选特定产品的订单
	Status    int8 `json:"status,omitempty"`                  // 可选：订单状态，用于筛选特定状态的订单
}
