package dto

import "time"

type OrderRequest struct {
	TableNumber int                 `json:"table_number"`
	UserId      int                 `json:"user_id"`
	PromoId     int                 `json:"promo_id"`
	Amount      int64               `json:"amount"`
	Discount    int64               `json:"discount"`
	TotalAmount int64               `json:"total_amount"`
	OrderStatus string              `json:"order_status"`
	OrderItem   *[]OrderItemRequest `json:"order_item"`
	CreatedDate time.Time           `json:"created_date"`
	CreatedBy   string              `json:"created_by"`
	UpdatedDate time.Time           `json:"updated_date"`
	UpdatedBy   string              `json:"updated_by"`
}

type OrderItemRequest struct {
	OrderId    int   `json:"order_id"`
	MenuId     int   `json:"menu_id"`
	Price      int64 `json:"price"`
	Quantity   int   `json:"quantity"`
	TotalPrice int64 `json:"total_price"`
}

type OrderResponse struct {
	OrderId     int                  `json:"order_id"`
	TableNumber int                  `json:"table_number"`
	UserId      int                  `json:"user_id"`
	PromoId     int                  `json:"promo_id"`
	Amount      int64                `json:"amount"`
	Discount    int64                `json:"discount"`
	TotalAmount int64                `json:"total_amount"`
	OrderStatus string               `json:"order_status"`
	OrderItem   *[]OrderItemResponse `json:"order_item"`
	LinkPayment string               `json:"link_payment"`
	CreatedDate time.Time            `json:"created_date"`
	CreatedBy   string               `json:"created_by"`
	UpdatedDate time.Time            `json:"updated_date"`
	UpdatedBy   string               `json:"updated_by"`
}

type OrderResponse2 struct {
	OrderId     int       `json:"order_id"`
	TableNumber int       `json:"table_number"`
	UserId      int       `json:"user_id"`
	PromoId     int       `json:"promo_id"`
	Amount      int64     `json:"amount"`
	Discount    int64     `json:"discount"`
	TotalAmount int64     `json:"total_amount"`
	OrderStatus string    `json:"order_status"`
	CreatedDate time.Time `json:"created_date"`
	CreatedBy   string    `json:"created_by"`
	UpdatedDate time.Time `json:"updated_date"`
	UpdatedBy   string    `json:"updated_by"`
}

type OrderItemResponse struct {
	OrderItemId int   `json:"order_item_id"`
	OrderId     int   `json:"order_id"`
	MenuId      int   `json:"menu_id"`
	Price       int64 `json:"price"`
	Quantity    int   `json:"quantity"`
	TotalPrice  int64 `json:"total_price"`
}

type PromoRequest struct {
	PromoName          string    `json:"promo_name"`
	MinimumAmount      int64     `json:"minimum_amount"`
	DiscountPercentage int32     `json:"discount_percentage"`
	MaxDiscountAmount  int64     `json:"max_discount_amount"`
	IsActive           bool      `json:"is_active"`
	CreatedDate        time.Time `json:"created_date"`
	CreatedBy          string    `json:"created_by"`
	UpdatedDate        time.Time `json:"updated_date"`
	UpdatedBy          string    `json:"updated_by"`
}

type PromoResponse struct {
	PromoId            int       `json:"promo_id"`
	PromoName          string    `json:"promo_name"`
	MinimumAmount      int64     `json:"minimum_amount"`
	DiscountPercentage int32     `json:"discount_percentage"`
	MaxDiscountAmount  int64     `json:"max_discount_amount"`
	IsActive           bool      `json:"is_active"`
	CreatedDate        time.Time `json:"created_date"`
	CreatedBy          string    `json:"created_by"`
	UpdatedDate        time.Time `json:"updated_date"`
	UpdatedBy          string    `json:"updated_by"`
}

type PromoResponse2 struct {
	PromoId  int   `json:"promo_id"`
	Discount int64 `json:"discount"`
}
