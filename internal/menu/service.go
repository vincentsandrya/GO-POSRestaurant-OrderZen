package menu

import (
	"errors"

	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/dto"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

func (svc *Service) AddMenuCategory(req *dto.MenuCategoryRequest) (*dto.MenuCategoryResponse, error) {

	checkExists, err := svc.Repository.GetMenuCategoryByCategoryName(req.CategoryName)
	if err != nil {
		return nil, err
	}

	if checkExists != nil {
		return nil, errors.New("category already exist")
	}

	data, err := svc.Repository.AddMenuCategory(req)
	if err != nil {
		return nil, err
	}

	resp, err := svc.Repository.GetMenuCategoryById(uint(data.Id))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *Service) GetMenuCategoryById(id int) (*dto.MenuCategoryResponse, error) {
	resp, err := svc.Repository.GetMenuCategoryById(uint(id))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *Service) GetMenuCategory() (*[]dto.MenuCategoryResponse, error) {
	resp, err := svc.Repository.GetMenuCategory()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *Service) EditMenuCategoryById(MenuCategoryID int, req *dto.MenuCategoryRequest) (*dto.MenuCategoryResponse, error) {
	err := svc.Repository.EditMenuCategoryById(MenuCategoryID, req)
	if err != nil {
		return nil, err
	}

	resp, err := svc.Repository.GetMenuCategoryById(uint(MenuCategoryID))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *Service) DeleteMenuCategoryById(MenuCategoryID int) error {
	_, err := svc.Repository.GetMenuCategoryById(uint(MenuCategoryID))
	if err != nil {
		return err
	}

	err = svc.Repository.DeleteMenuCategoryById(MenuCategoryID)
	if err != nil {
		return err
	}

	return nil
}

// MENU

func (svc *Service) AddMenu(req *dto.MenuRequest) (*dto.MenuResponse, error) {

	data, err := svc.Repository.AddMenu(req)
	if err != nil {
		return nil, err
	}

	resp, err := svc.Repository.GetMenuById(uint(data.Id))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *Service) GetMenuById(id int) (*dto.MenuResponse, error) {
	resp, err := svc.Repository.GetMenuById(uint(id))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *Service) GetMenu() (*[]dto.MenuResponse, error) {
	resp, err := svc.Repository.GetMenu()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *Service) EditMenuById(MenuID int, req *dto.MenuRequest) (*dto.MenuResponse, error) {
	err := svc.Repository.EditMenuById(MenuID, req)
	if err != nil {
		return nil, err
	}

	resp, err := svc.Repository.GetMenuById(uint(MenuID))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *Service) DeleteMenuById(MenuID int) error {
	_, err := svc.Repository.GetMenuById(uint(MenuID))
	if err != nil {
		return err
	}

	err = svc.Repository.DeleteMenuById(MenuID)
	if err != nil {
		return err
	}

	return nil
}
