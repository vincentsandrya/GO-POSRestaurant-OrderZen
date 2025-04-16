package order

import (
	"fmt"
	"strconv"

	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/dto"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/internal/menu"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/internal/midtrans"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/models"
)

type Service struct {
	Repository  *Repository
	MenuService *menu.Service
}

func NewService(repository *Repository, menuService *menu.Service) *Service {
	return &Service{
		Repository:  repository,
		MenuService: menuService,
	}
}

func (svc *Service) AddOrder(req *dto.OrderRequest) (*dto.OrderResponse, error) {
	orderItem := *req.OrderItem
	var amount = 0

	for i, item := range orderItem {
		menuData, err := svc.MenuService.GetMenuById(item.MenuId)
		if err != nil {
			return nil, err
		}

		orderItem[i].Price = int64(menuData.Price)
		orderItem[i].TotalPrice = int64(menuData.Price) * int64(item.Quantity)
		amount += int(orderItem[i].TotalPrice)
	}

	req.Amount = int64(amount)

	dataPromo, err := svc.Repository.GetBestPromoForAmount(req.Amount)
	req.PromoId = dataPromo.PromoId
	req.Discount = dataPromo.Discount
	req.TotalAmount = req.Amount - req.Discount

	req.OrderStatus = "Waiting for Payment"

	data, err := svc.Repository.AddOrder(req)
	if err != nil {
		return nil, err
	}

	for _, item := range orderItem {
		item.OrderId = data.Id
		_, err := svc.Repository.AddOrderItem(&item)
		if err != nil {
			return nil, err
		}
	}

	res, err := svc.Repository.GetOrderById(uint(data.Id))
	res.LinkPayment = midtrans.CreateOrder(int(data.Id), req.TotalAmount)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc *Service) GetOrderById(id int) (*dto.OrderResponse, error) {
	resp, err := svc.Repository.GetOrderById(uint(id))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *Service) GetOrder() (*[]dto.OrderResponse2, error) {
	resp, err := svc.Repository.GetOrder()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *Service) CheckPayment(id int) (*models.PaymentMidtrans, error) {

	midtransData, err := midtrans.CheckMidtransPaymentStatus(strconv.Itoa(id))
	if err != nil {
		return nil, fmt.Errorf("failed to check Midtrans status: %w", err)
	}

	orderDate, err := svc.Repository.GetOrderById(uint(id))
	if err != nil {
		return nil, err
	}

	if orderDate.OrderStatus == "Waiting for Payment" && (midtransData.TransactionStatus == "settlement" || midtransData.TransactionStatus == "capture") {

		if _, err := svc.Repository.AddPaymentMidtrans(midtransData); err != nil {
			return nil, fmt.Errorf("failed to add Midtrans payment: %w", err)
		}

		updateReq := &dto.OrderRequest{
			OrderStatus: "Paid",
			UpdatedBy:   "CheckPayment",
		}

		if err := svc.Repository.EditOrderById(id, updateReq); err != nil {
			return nil, fmt.Errorf("failed to update order status: %w", err)
		}
	}

	return midtransData, nil
}

func (svc *Service) GetPayment(id int) (*models.PaymentMidtrans, error) {
	midtransData, err := midtrans.CheckMidtransPaymentStatus(strconv.Itoa(id))
	if err != nil {
		return nil, fmt.Errorf("failed to check Midtrans status: %w", err)
	}

	return midtransData, nil
}

// PROMO
func (svc *Service) GetPromoById(id int) (*dto.PromoResponse, error) {
	resp, err := svc.Repository.GetPromoById(uint(id))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *Service) GetPromo() (*[]dto.PromoResponse, error) {
	resp, err := svc.Repository.GetPromo()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *Service) EditPromoById(id int, req *dto.PromoRequest) (*dto.PromoResponse, error) {
	err := svc.Repository.EditPromoById(id, req)
	if err != nil {
		return nil, err
	}

	resp, err := svc.Repository.GetPromoById(uint(id))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *Service) DeletePromoById(id int) error {
	_, err := svc.Repository.GetPromoById(uint(id))
	if err != nil {
		return err
	}

	err = svc.Repository.DeletePromoById(id)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) AddPromo(req *dto.PromoRequest) (*dto.PromoResponse, error) {
	data, err := svc.Repository.AddPromo(req)
	if err != nil {
		return nil, err
	}

	resp, err := svc.Repository.GetPromoById(uint(data.Id))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
