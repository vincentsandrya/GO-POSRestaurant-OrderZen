package models

import "time"

type Order struct {
	OrderId     int       `json:"order_id" gorm:"primaryKey;autoIncrement"`
	TableNumber int       `json:"table_number" gorm:"not null"`
	UserId      int       `json:"user_id"`
	User        *User     `json:"user" gorm:"foreignKey:UserId;references:UserId"`
	PromoId     int       `json:"promo_id"`
	Promo       *Promo    `json:"promo" gorm:"foreignKey:PromoId;references:PromoId"`
	Amount      int64     `json:"amount"`
	Discount    int64     `json:"discount"`
	TotalAmount int64     `json:"total_amount"`
	OrderStatus string    `json:"order_status"`
	CreatedDate time.Time `json:"created_date" gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy   string    `json:"created_by"`
	UpdatedDate time.Time `json:"updated_date" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedBy   string    `json:"updated_by"`
}

type OrderItem struct {
	OrderItemId int    `json:"order_item_id" gorm:"primaryKey;autoIncrement"`
	OrderId     int    `json:"order_id" gorm:"not null"`
	Order       *Order `json:"order" gorm:"foreignKey:OrderId;references:OrderId"`
	MenuId      int    `json:"menu_id" gorm:"not null"`
	Menu        *Menu  `json:"menu" gorm:"foreignKey:MenuId;references:MenuId"`
	Price       int64  `json:"price" gorm:"not null"`
	Quantity    int    `json:"quantity" gorm:"not null"`
	TotalPrice  int64  `json:"total_price" gorm:"not null"`
}

type Promo struct {
	PromoId            int       `json:"promo_id" gorm:"primaryKey;autoIncrement"`
	PromoName          string    `json:"promo_name" gorm:"not null;unique"`
	MinimumAmount      int64     `json:"minimum_amount" gorm:"not null"`
	DiscountPercentage int32     `json:"discount_percentage" gorm:"not null"`
	MaxDiscountAmount  int64     `json:"max_discount_amount" gorm:"not null"`
	IsActive           bool      `json:"is_active" gorm:"not null"`
	CreatedDate        time.Time `json:"created_date" gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy          string    `json:"created_by"`
	UpdatedDate        time.Time `json:"updated_date" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedBy          string    `json:"updated_by"`
}

type PaymentMidtrans struct {
	TransactionId     string    `json:"transaction_id" gorm:"not null;unique"`
	OrderId           string    `json:"order_id" gorm:"not null"`
	Order             *Order    `json:"order" gorm:"foreignKey:OrderId;references:OrderId"`
	GrossAmount       int64     `json:"gross_amount" gorm:"not null"`
	Currency          string    `json:"currency" gorm:"not null"`
	PaymentType       string    `json:"payment_type" gorm:"not null"`
	TransactionStatus string    `json:"transaction_status" gorm:"not null"`
	StatusMessage     string    `json:"status_message" gorm:"not null"`
	MerchantId        string    `json:"merchant_id" gorm:"not null"`
	TransactionTime   time.Time `json:"transaction_time" gorm:"not null"`
}
