package order

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/dto"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/helpers"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/models"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetOrder() (*[]dto.OrderResponse2, error) {
	var err error
	res := []dto.OrderResponse2{}

	err = r.DB.Model(&models.Order{}).Scan(&res).Error

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *Repository) GetOrderById(id uint) (*dto.OrderResponse, error) {
	var err error
	res := dto.OrderResponse{}

	err = r.DB.Model(&models.Order{}).Where("order_id = ?", id).First(&res).Error

	if err != nil {
		return nil, err
	}

	var items []dto.OrderItemResponse
	err = r.DB.Table("order_items").
		Select(`
			order_items.order_item_id,
			order_items.order_id,
			order_items.menu_id,
			menus.name as menu_name,
			order_items.quantity,
			order_items.price,
			order_items.total_price`).
		Joins("LEFT JOIN menus ON menus.menu_id = order_items.menu_id").
		Where("order_items.order_id = ?", id).
		Scan(&items).Error

	if err != nil {
		return nil, err
	}

	res.OrderItem = &items

	return &res, nil
}

func (r *Repository) AddOrder(req *dto.OrderRequest) (*dto.PayloadID, error) {
	model, err := helpers.MappingOrderModel(req)
	if err != nil {
		return nil, err
	}
	res := r.DB.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}
	return &dto.PayloadID{Id: int(model.OrderId)}, nil
}

func (r *Repository) AddOrderItem(req *dto.OrderItemRequest) (*dto.PayloadID, error) {
	model, err := helpers.MappingOrderItemModel(req)
	if err != nil {
		return nil, err
	}
	res := r.DB.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}
	return &dto.PayloadID{Id: int(model.OrderItemId)}, nil
}

func (r *Repository) AddPaymentMidtrans(model *models.PaymentMidtrans) (*dto.PayloadID, error) {
	res := r.DB.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}

	id, _ := strconv.Atoi(model.OrderId)

	return &dto.PayloadID{Id: id}, nil
}

func (r *Repository) GetPaymentMidtransById(id uint) (*models.PaymentMidtrans, error) {
	res := models.PaymentMidtrans{}
	err := r.DB.Model(&models.PaymentMidtrans{}).Where("order_id = ?", strconv.Itoa(int(id))).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *Repository) EditOrderById(id int, req *dto.OrderRequest) error {
	updateData := map[string]interface{}{
		"order_status": req.OrderStatus,
		"updated_by":   req.UpdatedBy,
		"updated_date": time.Now(),
	}

	fmt.Printf("EditOrderById: updating order_id = %d with %+v\n", id, updateData)

	response := r.DB.Model(&models.Order{}).Where("order_id = ?", id).Updates(updateData)
	if response.Error != nil {
		return response.Error
	}

	if response.RowsAffected == 0 {
		return fmt.Errorf("no order found with id %d", id)
	}

	return nil
}

//PROMO

func (r *Repository) GetPromo() (*[]dto.PromoResponse, error) {
	var err error
	res := []dto.PromoResponse{}

	err = r.DB.Model(&models.Promo{}).Scan(&res).Error

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *Repository) GetPromoById(id uint) (*dto.PromoResponse, error) {
	var err error
	res := dto.PromoResponse{}

	err = r.DB.Model(&models.Promo{}).Where("promo_id = ?", id).First(&res).Error

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *Repository) EditPromoById(id int, req *dto.PromoRequest) error {
	updateData := map[string]interface{}{
		"promo_name":          req.PromoName,
		"minimum_amount":      req.MinimumAmount,
		"discount_percentage": req.DiscountPercentage,
		"max_discount_amount": req.MaxDiscountAmount,
		"is_active":           req.IsActive,
		"updated_by":          req.UpdatedBy,
		"updated_date":        time.Now(),
	}

	response := r.DB.Model(&models.Promo{}).Where("promo_id = ?", id).Updates(updateData)
	if response.Error != nil {
		return response.Error
	}
	return nil
}

func (r *Repository) DeletePromoById(id int) error {
	err := r.DB.Debug().Delete(&models.Promo{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) AddPromo(req *dto.PromoRequest) (*dto.PayloadID, error) {
	model, err := helpers.MappingPromoModel(req)
	if err != nil {
		return nil, err
	}
	res := r.DB.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}
	return &dto.PayloadID{Id: int(model.PromoId)}, nil
}

func (r *Repository) GetBestPromoForAmount(amount int64) (*dto.PromoResponse2, error) {
	res := dto.PromoResponse2{}
	var err error

	query := `
		SELECT promo_id,
		       CASE 
		           WHEN ? * discount_percentage < max_discount_amount 
		           THEN ? * discount_percentage 
		           ELSE max_discount_amount 
		       END AS discount
		FROM promos
		WHERE minimum_amount <= ?
		ORDER BY 
		    CASE 
		        WHEN ? * discount_percentage < max_discount_amount 
		        THEN ? * discount_percentage 
		        ELSE max_discount_amount 
		    END DESC
		LIMIT 1;
	`

	err = r.DB.Raw(query, amount, amount, amount, amount, amount).Scan(&res).Error
	if err != nil {
		return nil, err
	}

	return &res, nil
}
