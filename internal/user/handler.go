package user

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/display"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/dto"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

func (hand *Handler) Login(c *gin.Context) {
	var loginInfo dto.LoginRequest
	err := c.ShouldBindJSON(&loginInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.
			StatusBadRequest, dto.ResponseFailed(display.ErrorInvalidBody.MessageErr, display.ErrorInvalidBody.CodeErr))
		return
	}

	response, err := hand.Service.Login(loginInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ResponseFailed(err.Error(), http.StatusUnauthorized))
		return
	}

	// Generate JWT token on successful login
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &dto.JWTClaims{
		UserId: int(response.UserId),
		Email:  response.Email,
		RoleId: int(response.RoleId),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	jwtKey := []byte(os.Getenv("API_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("Could not create token", http.StatusInternalServerError))
		return
	}

	c.SetCookie(
		"authToken",
		tokenString,
		int(expirationTime.Unix()-time.Now().Unix()),
		"/",
		"",   // Domain, can be set to your specific domain
		true, // Secure, set to true in production
		true, // HttpOnly
	)

	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("login successfully", response))
}

func (hand *Handler) Register(c *gin.Context) {
	var req dto.UserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(display.ErrorInvalidBody.MessageErr, display.ErrorInvalidBody.CodeErr))
		return
	}

	req.RoleId = 4 //customer

	req.CreatedBy = req.Email
	req.CreatedDate = time.Now()
	req.UpdatedBy = req.Email
	req.UpdatedDate = time.Now()

	res, err := hand.Service.AddUser(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to add new user, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to add new user ", res))
}

func (hand *Handler) AddUser(c *gin.Context) {
	var req dto.UserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(display.ErrorInvalidBody.MessageErr, display.ErrorInvalidBody.CodeErr))
		return
	}

	//admin only
	roleId, _ := c.Get("role_id")
	if roleId != 1 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ResponseFailed("unauthorized", http.StatusUnauthorized))
		return
	}

	email, _ := c.Get("email")
	req.CreatedBy = email.(string)
	req.CreatedDate = time.Now()
	req.UpdatedBy = email.(string)
	req.UpdatedDate = time.Now()

	res, err := hand.Service.AddUser(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to add new user, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to add new user ", res))
}

func (hand *Handler) GetUser(c *gin.Context) {
	//admin only
	roleId, _ := c.Get("role_id")
	if roleId != 1 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ResponseFailed("unauthorized", http.StatusUnauthorized))
		return
	}

	res, err := hand.Service.GetUser()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to retrieve user, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to retrieve user ", res))
}

func (hand *Handler) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(display.ErrorInvalidParamID.MessageErr, display.ErrorInvalidParamID.CodeErr))
		return
	}

	res, err := hand.Service.GetUserById(int(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to retrieve user, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to retrieve user ", res))
}

func (hand *Handler) EditUserById(c *gin.Context) {
	var req dto.UserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(display.ErrorInvalidBody.MessageErr, display.ErrorInvalidBody.CodeErr))
		return
	}

	//admin only
	roleId, _ := c.Get("role_id")
	if roleId != 1 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ResponseFailed("unauthorized", http.StatusUnauthorized))
		return
	}

	email, _ := c.Get("email")
	req.UpdatedBy = email.(string)
	req.UpdatedDate = time.Now()

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(display.ErrorInvalidParamID.MessageErr, display.ErrorInvalidParamID.CodeErr))
		return
	}

	res, err := hand.Service.EditUserById(userID, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to edit user, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to edit user ", res))
}
func (hand *Handler) DeleteUserById(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(display.ErrorInvalidParamID.MessageErr, display.ErrorInvalidParamID.CodeErr))
		return
	}

	//admin only
	roleId, _ := c.Get("role_id")
	if roleId != 1 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ResponseFailed("unauthorized", http.StatusUnauthorized))
		return
	}

	err = hand.Service.DeleteUserById(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed("failed to delete user, "+err.Error(), http.StatusBadRequest))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesNoData("success to delete user with ID "+userIDStr))
}
