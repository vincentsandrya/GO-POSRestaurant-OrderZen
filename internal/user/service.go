package user

import (
	"errors"

	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/display"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/dto"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/helpers"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{Repository: repository}
}

func (svc *Service) Login(req dto.LoginRequest) (*dto.UserResponseLogin, error) {
	user, err := svc.Repository.Login(req.Email)
	if err != nil {
		return nil, display.ErrorWrongCredentialsLogin
	}

	if !helpers.ComparePassword(user.Password, req.Password) {
		return nil, display.ErrorWrongCredentialsLogin
	}

	token, err := dto.GenerateAccessToken(dto.JWTClaims{
		UserId: int(user.UserId),
		Email:  user.Email,
		RoleId: user.RoleId,
	})

	if err != nil {
		return nil, errors.New("error generate access token")
	}

	resp, err := svc.Repository.GetUserById(uint(user.UserId))
	if err != nil {
		return nil, err
	}

	response := &dto.UserResponseLogin{
		UserId: resp.UserId,
		Name:   resp.Name,
		Email:  resp.Email,
		RoleId: resp.RoleId,
	}

	response.Token = token

	return response, nil
}

func (svc *Service) AddUser(req *dto.UserRequest) (*dto.UserResponse, error) {

	checkEmailExist, err := svc.Repository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if checkEmailExist != nil {
		return nil, errors.New("email already exist")
	}

	data, err := svc.Repository.AddUser(req)
	if err != nil {
		return nil, err
	}

	resp, err := svc.Repository.GetUserById(uint(data.Id))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *Service) GetUserById(id int) (*dto.UserResponse, error) {
	resp, err := svc.Repository.GetUserById(uint(id))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *Service) GetUser() (*[]dto.UserResponse, error) {
	resp, err := svc.Repository.GetUser()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *Service) EditUserById(userID int, req *dto.UserRequest) (*dto.UserResponse, error) {
	err := svc.Repository.EditUserById(userID, req)
	if err != nil {
		return nil, err
	}

	resp, err := svc.Repository.GetUserById(uint(userID))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (svc *Service) DeleteUserById(userID int) error {
	_, err := svc.Repository.GetUserById(uint(userID))
	if err != nil {
		return err
	}

	err = svc.Repository.DeleteUserById(userID)
	if err != nil {
		return err
	}

	return nil
}
