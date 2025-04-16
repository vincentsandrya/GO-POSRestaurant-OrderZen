package models

import "time"

type Menu struct {
	MenuId         int           `json:"menu_id" gorm:"primaryKey;autoIncrement"`
	Name           string        `json:"name" gorm:"not null"`
	Description    string        `json:"description"`
	Price          int           `json:"price"`
	MenuCategoryId int           `json:"menu_category_id"`
	MenuCategory   *MenuCategory `json:"menu_category" gorm:"foreignKey:MenuCategoryId;references:MenuCategoryId"`
	CreatedDate    time.Time     `json:"created_date" gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy      string        `json:"created_by"`
	UpdatedDate    time.Time     `json:"updated_date" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedBy      string        `json:"updated_by"`
}

type MenuCategory struct {
	MenuCategoryId int       `json:"menu_category_id" gorm:"primaryKey;autoIncrement"`
	CategoryName   string    `json:"category_name" gorm:"not null;unique"`
	CreatedDate    time.Time `json:"created_date" gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy      string    `json:"created_by"`
	UpdatedDate    time.Time `json:"updated_date" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedBy      string    `json:"updated_by"`
}
