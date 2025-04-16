package menu

import (
	"errors"
	"time"

	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/dto"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/helpers"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/models"
	"gorm.io/gorm"
)

// Repository defines methods for accessing products in the database
type Repository struct {
	DB *gorm.DB
}

// NewRepository creates a new product repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

//START MENU CATEGORY

func (r *Repository) GetMenuCategory() (*[]dto.MenuCategoryResponse, error) {
	var res []dto.MenuCategoryResponse

	err := r.DB.Model(&models.MenuCategory{}).Scan(&res).Error
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *Repository) GetMenuCategoryById(id uint) (*dto.MenuCategoryResponse, error) {
	var err error
	res := dto.MenuCategoryResponse{}

	err = r.DB.Model(&models.MenuCategory{}).Where("menu_category_id = ?", id).First(&res).Error

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *Repository) GetMenuCategoryByCategoryName(categoryName string) (*dto.MenuCategoryResponse, error) {
	var err error
	res := dto.MenuCategoryResponse{}

	err = r.DB.Model(&models.MenuCategory{}).Where("category_name = ?", categoryName).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *Repository) AddMenuCategory(req *dto.MenuCategoryRequest) (*dto.PayloadID, error) {
	model, err := helpers.MappingMenuCategoryModel(req)
	if err != nil {
		return nil, err
	}
	res := r.DB.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}
	return &dto.PayloadID{Id: int(model.MenuCategoryId)}, nil
}

func (r *Repository) EditMenuCategoryById(id int, req *dto.MenuCategoryRequest) error {

	updateData := map[string]interface{}{
		"category_name": req.CategoryName,
		"updated_by":    req.UpdatedBy,
		"updated_date":  time.Now(),
	}

	response := r.DB.Model(&models.MenuCategory{}).Where("menu_category_id = ?", id).Updates(updateData)

	if response.Error != nil {
		return response.Error
	}
	return nil
}

func (r *Repository) DeleteMenuCategoryById(id int) error {
	err := r.DB.Debug().Delete(&models.MenuCategory{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

//START MENU

func (r *Repository) GetMenu() (*[]dto.MenuResponse, error) {
	var res []dto.MenuResponse

	err := r.DB.Table("menus").
		Select(`
			menus.menu_id, 
			menus.name, 
			menus.description, 
			menus.price, 
			menus.menu_category_id, 
			menu_categories.category_name,
			menus.created_by,
			menus.created_date,
			menus.updated_by,
			menus.updated_date`).
		Joins("left join menu_categories on menu_categories.menu_category_id = menus.menu_category_id").
		Scan(&res).Error

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *Repository) GetMenuById(id uint) (*dto.MenuResponse, error) {
	var res dto.MenuResponse

	err := r.DB.Table("menus").
		Select(`
			menus.menu_id, 
			menus.name, 
			menus.description, 
			menus.price, 
			menus.menu_category_id, 
			menu_categories.category_name,
			menus.created_by,
			menus.created_date,
			menus.updated_by,
			menus.updated_date`).
		Joins("left join menu_categories on menu_categories.menu_category_id = menus.menu_category_id").
		Where("menus.menu_id = ?", id).
		First(&res).Error

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *Repository) AddMenu(req *dto.MenuRequest) (*dto.PayloadID, error) {
	model, err := helpers.MappingMenuModel(req)
	if err != nil {
		return nil, err
	}
	res := r.DB.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}
	return &dto.PayloadID{Id: int(model.MenuId)}, nil
}

func (r *Repository) EditMenuById(id int, req *dto.MenuRequest) error {
	updateData := map[string]interface{}{
		"name":             req.Name,
		"description":      req.Description,
		"price":            req.Price,
		"menu_category_id": req.MenuCategoryId,
		"updated_by":       req.UpdatedBy,
		"updated_date":     time.Now(),
	}

	response := r.DB.Model(&models.Menu{}).Where("menu_id = ?", id).Updates(updateData)

	if response.Error != nil {
		return response.Error
	}
	return nil
}

func (r *Repository) DeleteMenuById(id int) error {
	err := r.DB.Debug().Delete(&models.Menu{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
