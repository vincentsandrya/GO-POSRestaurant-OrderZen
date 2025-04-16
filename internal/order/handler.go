package order

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/display"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/dto"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

func (hand *Handler) AddOrder(c *gin.Context) {
	var req dto.OrderRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(display.ErrorInvalidBody.MessageErr, display.ErrorInvalidBody.CodeErr))
		return
	}

	//customer only
	roleId, _ := c.Get("role_id")
	if roleId != 4 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ResponseFailed("unauthorized", http.StatusUnauthorized))
		return
	}

	userId, _ := c.Get("user_id")
	req.UserId = userId.(int)

	email, _ := c.Get("email")
	req.CreatedBy = email.(string)
	req.CreatedDate = time.Now()
	req.UpdatedBy = email.(string)
	req.UpdatedDate = time.Now()

	res, err := hand.Service.AddOrder(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to add new order, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to add new order ", res))
}

func (hand *Handler) GetOrder(c *gin.Context) {
	res, err := hand.Service.GetOrder()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to retrieve order, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to retrieve order ", res))
}

func (hand *Handler) GetOrderById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(display.ErrorInvalidParamID.MessageErr, display.ErrorInvalidParamID.CodeErr))
		return
	}

	res, err := hand.Service.GetOrderById(int(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to retrieve order, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to retrieve order ", res))
}

func (hand *Handler) GetPayment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(display.ErrorInvalidParamID.MessageErr, display.ErrorInvalidParamID.CodeErr))
		return
	}

	res, err := hand.Service.GetPayment(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to retrieve order, "+err.Error(), http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to retrieve order ", res))
}

func (hand *Handler) CheckPayment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(display.ErrorInvalidParamID.MessageErr, display.ErrorInvalidParamID.CodeErr))
		return
	}

	res, err := hand.Service.CheckPayment(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to retrieve order, "+err.Error(), http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to retrieve order ", res))
}

// PROMO
func (hand *Handler) GetPromo(c *gin.Context) {
	res, err := hand.Service.GetPromo()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to retrieve Promo, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to retrieve Promo ", res))
}

func (hand *Handler) GetPromoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(display.ErrorInvalidParamID.MessageErr, display.ErrorInvalidParamID.CodeErr))
		return
	}

	res, err := hand.Service.GetPromoById(int(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to retrieve Promo, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to retrieve Promo ", res))
}

func (hand *Handler) EditPromoById(c *gin.Context) {
	var req dto.PromoRequest
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

	PromoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(display.ErrorInvalidParamID.MessageErr, display.ErrorInvalidParamID.CodeErr))
		return
	}

	res, err := hand.Service.EditPromoById(PromoID, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to edit Promo, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to edit Promo ", res))
}
func (hand *Handler) DeletePromoById(c *gin.Context) {
	PromoIDStr := c.Param("id")
	PromoID, err := strconv.Atoi(PromoIDStr)
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

	err = hand.Service.DeletePromoById(PromoID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed("failed to delete Promo, "+err.Error(), http.StatusBadRequest))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesNoData("success to delete Promo with ID "+PromoIDStr))
}

func (hand *Handler) AddPromo(c *gin.Context) {
	var req dto.PromoRequest
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

	res, err := hand.Service.AddPromo(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to add new Promo, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to add new Promo ", res))
}
