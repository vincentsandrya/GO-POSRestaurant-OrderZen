package user

import (
	"errors"
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

func (r *Repository) Login(email string) (*models.User, error) {
	var cust models.User
	err := r.DB.Debug().First(&cust, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &cust, nil
}

func (r *Repository) GetUser() (*[]dto.UserResponse, error) {
	var res []dto.UserResponse

	err := r.DB.Model(&models.User{}).Scan(&res).Error
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *Repository) GetUserById(id uint) (*dto.UserResponse, error) {
	var err error
	res := dto.UserResponse{}

	err = r.DB.Model(&models.User{}).Where("user_id = ?", id).First(&res).Error

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *Repository) GetUserByEmail(email string) (*dto.UserResponse, error) {
	var err error
	res := dto.UserResponse{}

	err = r.DB.Model(&models.User{}).Where("email = ?", email).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *Repository) AddUser(req *dto.UserRequest) (*dto.PayloadID, error) {
	model, err := helpers.MappingUserModel(req)
	if err != nil {
		return nil, err
	}
	res := r.DB.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}
	return &dto.PayloadID{Id: int(model.UserId)}, nil
}

func (r *Repository) EditUserById(userID int, req *dto.UserRequest) error {
	updateData := map[string]interface{}{
		"email":        req.Email,
		"name":         req.Name,
		"password":     req.Password,
		"role_id":      req.RoleId,
		"updated_by":   req.UpdatedBy,
		"updated_date": time.Now(),
	}

	response := r.DB.Model(&models.User{}).Where("user_id = ?", userID).Updates(updateData)
	if response.Error != nil {
		return response.Error
	}
	return nil
}

func (r *Repository) DeleteUserById(id int) error {
	err := r.DB.Debug().Delete(&models.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
