package vo

type CheckUserReq struct {
	OpenId string `json:"open_id"`
	UserId string `json:"user_id"`
}

type CheckUserResp struct {
	CommonVo
	Data CheckUserData `json:"data"`
}

type CheckUserData struct {
	Status          string `json:"status"`
	PricePlanId     string `json:"price_plan_id"`
	IsTrial         bool   `json:"is_trial"`
	ServiceStopTime string `json:"service_stop_time"`
}

type GetOrderListReq struct {
	Status    *string `json:"status"`
	PageSize  int     `json:"page_size"`
	PageToken *string `json:"page_token"`
	TenantKey *string `json:"tenant_key"`
}

type GetOrderListResp struct {
	CommonVo
	Data GetOrderListData `json:"data"`
}

type GetOrderListData struct {
	Total     int64       `json:"total"`
	HasMore   bool        `json:"has_more"`
	PageToken string      `json:"page_token"`
	OrderList []OrderInfo `json:"order_list"`
}

type OrderInfo struct {
	OrderId       string `json:"order_id"`
	PricePlanId   string `json:"price_plan_id"`
	PricePlanType string `json:"price_plan_type"`
	Seats         int    `json:"seats"`
	BuyCount      int    `json:"buy_count"`
	CreateTime    string `json:"create_time"`
	PayTime       string `json:"pay_time"`
	Status        string `json:"status"`
	BuyType       string `json:"buy_type"`
	SrcOrderId    string `json:"src_order_id"`
	DstOrderId    string `json:"dst_order_id"`
	OrderPayPrice int64  `json:"order_pay_price"`
	TenantKey     string `json:"tenant_key"`
}

type OrderInfoResp struct {
	CommonVo
	Data OrderInfoData `json:"data"`
}

type OrderInfoData struct {
	Order OrderInfo `json:"order"`
}
