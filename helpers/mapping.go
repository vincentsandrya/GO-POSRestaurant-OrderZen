package helpers

import (
	dto "github.com/vincentsandrya/GO-POSRestaurant-OrderZen/dto"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/models"
)

func MappingUserModel(req *dto.UserRequest) (*models.User, error) {
	hashedPass, err := hashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	return &models.User{
		Name:        req.Name,
		Email:       req.Email,
		RoleId:      int(req.RoleId),
		Password:    hashedPass,
		CreatedDate: req.CreatedDate,
		CreatedBy:   req.CreatedBy,
		UpdatedDate: req.UpdatedDate,
		UpdatedBy:   req.UpdatedBy,
	}, nil
}

func MappingMenuCategoryModel(req *dto.MenuCategoryRequest) (*models.MenuCategory, error) {

	return &models.MenuCategory{
		CategoryName: req.CategoryName,
		CreatedDate:  req.CreatedDate,
		CreatedBy:    req.CreatedBy,
		UpdatedDate:  req.UpdatedDate,
		UpdatedBy:    req.UpdatedBy,
	}, nil
}

func MappingMenuModel(req *dto.MenuRequest) (*models.Menu, error) {

	return &models.Menu{
		Name:           req.Name,
		Description:    req.Description,
		Price:          req.Price,
		MenuCategoryId: req.MenuCategoryId,
		CreatedDate:    req.CreatedDate,
		CreatedBy:      req.CreatedBy,
		UpdatedDate:    req.UpdatedDate,
		UpdatedBy:      req.UpdatedBy,
	}, nil
}

func MappingOrderModel(req *dto.OrderRequest) (*models.Order, error) {

	return &models.Order{
		TableNumber: req.TableNumber,
		UserId:      req.UserId,
		PromoId:     req.PromoId,
		Amount:      req.Amount,
		Discount:    req.Discount,
		TotalAmount: req.TotalAmount,
		OrderStatus: req.OrderStatus,
		CreatedDate: req.CreatedDate,
		CreatedBy:   req.CreatedBy,
		UpdatedDate: req.UpdatedDate,
		UpdatedBy:   req.UpdatedBy,
	}, nil
}

func MappingOrderItemModel(req *dto.OrderItemRequest) (*models.OrderItem, error) {

	return &models.OrderItem{
		OrderId:    req.OrderId,
		MenuId:     req.MenuId,
		Price:      req.Price,
		Quantity:   req.Quantity,
		TotalPrice: req.TotalPrice,
	}, nil
}

func MappingPromoModel(req *dto.PromoRequest) (*models.Promo, error) {

	return &models.Promo{
		PromoName:          req.PromoName,
		MinimumAmount:      req.MinimumAmount,
		DiscountPercentage: req.DiscountPercentage,
		MaxDiscountAmount:  req.MaxDiscountAmount,
		IsActive:           req.IsActive,
		CreatedDate:        req.CreatedDate,
		CreatedBy:          req.CreatedBy,
		UpdatedDate:        req.UpdatedDate,
		UpdatedBy:          req.UpdatedBy,
	}, nil
}
