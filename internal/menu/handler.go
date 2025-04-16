package menu

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

func (hand *Handler) AddMenuCategory(c *gin.Context) {
	var req dto.MenuCategoryRequest
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

	res, err := hand.Service.AddMenuCategory(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to add new category, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to add new category ", res))
}

func (hand *Handler) GetMenuCategory(c *gin.Context) {
	res, err := hand.Service.GetMenuCategory()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to retrieve category, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to retrieve category ", res))
}

func (hand *Handler) GetMenuCategoryById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(display.ErrorInvalidParamID.MessageErr, display.ErrorInvalidParamID.CodeErr))
		return
	}

	res, err := hand.Service.GetMenuCategoryById(int(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to retrieve category, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to retrieve category ", res))
}

func (hand *Handler) EditMenuCategoryById(c *gin.Context) {
	var req dto.MenuCategoryRequest
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

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(display.ErrorInvalidParamID.MessageErr, display.ErrorInvalidParamID.CodeErr))
		return
	}

	res, err := hand.Service.EditMenuCategoryById(id, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to edit category, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to edit category ", res))
}
func (hand *Handler) DeleteMenuCategoryById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
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

	err = hand.Service.DeleteMenuCategoryById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed("failed to delete category, "+err.Error(), http.StatusBadRequest))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesNoData("success to delete category with ID "+idStr))
}

// MENU

func (hand *Handler) AddMenu(c *gin.Context) {
	var req dto.MenuRequest
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

	res, err := hand.Service.AddMenu(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to add new category, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to add new category ", res))
}

func (hand *Handler) GetMenu(c *gin.Context) {
	res, err := hand.Service.GetMenu()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to retrieve category, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to retrieve menu ", res))
}

func (hand *Handler) GetMenuById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(display.ErrorInvalidParamID.MessageErr, display.ErrorInvalidParamID.CodeErr))
		return
	}

	res, err := hand.Service.GetMenuById(int(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to retrieve menu, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to retrieve menu ", res))
}

func (hand *Handler) EditMenuById(c *gin.Context) {
	var req dto.MenuRequest
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

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed(display.ErrorInvalidParamID.MessageErr, display.ErrorInvalidParamID.CodeErr))
		return
	}

	res, err := hand.Service.EditMenuById(id, &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResponseFailed("failed to edit menu, "+err.Error(), http.StatusInternalServerError))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesWithData("success to edit menu ", res))
}
func (hand *Handler) DeleteMenuById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
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

	err = hand.Service.DeleteMenuById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseFailed("failed to delete menu, "+err.Error(), http.StatusBadRequest))
		return
	}
	c.JSON(http.StatusOK, dto.ResponseSuccesNoData("success to delete menu with ID "+idStr))
}
