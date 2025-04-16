package dto

import "time"

type MenuRequest struct {
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Price          int       `json:"price"`
	MenuCategoryId int       `json:"menu_category_id"`
	CreatedDate    time.Time `json:"created_date"`
	CreatedBy      string    `json:"created_by"`
	UpdatedDate    time.Time `json:"updated_date"`
	UpdatedBy      string    `json:"updated_by"`
}

type MenuResponse struct {
	MenuId       int       `json:"menu_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Price        int       `json:"price"`
	CategoryName string    `json:"category_name"`
	CreatedDate  time.Time `json:"created_date"`
	CreatedBy    string    `json:"created_by"`
	UpdatedDate  time.Time `json:"updated_date"`
	UpdatedBy    string    `json:"updated_by"`
}

type MenuCategoryRequest struct {
	CategoryName string    `json:"category_name"`
	CreatedDate  time.Time `json:"created_date"`
	CreatedBy    string    `json:"created_by"`
	UpdatedDate  time.Time `json:"updated_date"`
	UpdatedBy    string    `json:"updated_by"`
}

type MenuCategoryResponse struct {
	MenuCategoryId int       `json:"menu_category_id"`
	CategoryName   string    `json:"category_name"`
	CreatedDate    time.Time `json:"created_date"`
	CreatedBy      string    `json:"created_by"`
	UpdatedDate    time.Time `json:"updated_date"`
	UpdatedBy      string    `json:"updated_by"`
}
