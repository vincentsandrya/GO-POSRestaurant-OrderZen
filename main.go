package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/database"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/internal/menu"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/internal/order"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/internal/user"
	"github.com/vincentsandrya/GO-POSRestaurant-OrderZen/middleware"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Connect to PostgreSQL database using GORM
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	fmt.Println(db)

	routing(db)
}

func routing(db *gorm.DB) {
	// Initialize repository, service, and handler
	repoUser := user.NewRepository(db)
	servUser := user.NewService(repoUser)
	handUser := user.NewHandler(servUser)

	repoMenu := menu.NewRepository(db)
	servMenu := menu.NewService(repoMenu)
	handMenu := menu.NewHandler(servMenu)

	repoOrder := order.NewRepository(db)
	servOrder := order.NewService(repoOrder, servMenu)
	handOrder := order.NewHandler(servOrder)

	// Initialize Gin router
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())

	// login doesnt need token
	v1 := r.Group("/api/v1")
	v1.POST("/login", handUser.Login)
	v1.POST("/register", handUser.Register)

	const user = "/user"

	v1.POST(user, middleware.AuthorizeHandlerCookies(), handUser.AddUser)
	v1.GET(user, middleware.AuthorizeHandlerCookies(), handUser.GetUser)
	v1.GET(user+"/:id", middleware.AuthorizeHandlerCookies(), handUser.GetUserById)
	v1.PUT(user+"/:id", middleware.AuthorizeHandlerCookies(), handUser.EditUserById)
	v1.DELETE(user+"/:id", middleware.AuthorizeHandlerCookies(), handUser.DeleteUserById)

	const menuCategory = "/menu/category"

	v1.POST(menuCategory, middleware.AuthorizeHandlerCookies(), handMenu.AddMenuCategory)
	v1.GET(menuCategory, middleware.AuthorizeHandlerCookies(), handMenu.GetMenuCategory)
	v1.GET(menuCategory+"/:id", middleware.AuthorizeHandlerCookies(), handMenu.GetMenuCategoryById)
	v1.PUT(menuCategory+"/:id", middleware.AuthorizeHandlerCookies(), handMenu.EditMenuCategoryById)
	v1.DELETE(menuCategory+"/:id", middleware.AuthorizeHandlerCookies(), handMenu.DeleteMenuCategoryById)

	const menu = "/menu"

	v1.POST(menu, middleware.AuthorizeHandlerCookies(), handMenu.AddMenu)
	v1.GET(menu, middleware.AuthorizeHandlerCookies(), handMenu.GetMenu)
	v1.GET(menu+"/:id", middleware.AuthorizeHandlerCookies(), handMenu.GetMenuById)
	v1.PUT(menu+"/:id", middleware.AuthorizeHandlerCookies(), handMenu.EditMenuById)
	v1.DELETE(menu+"/:id", middleware.AuthorizeHandlerCookies(), handMenu.DeleteMenuById)

	// const promo = "/promo"
	// ...

	const order = "/order"

	v1.POST(order, middleware.AuthorizeHandlerCookies(), handOrder.AddOrder)
	v1.GET(order, middleware.AuthorizeHandlerCookies(), handOrder.GetOrder)
	v1.GET(order+"/:id", middleware.AuthorizeHandlerCookies(), handOrder.GetOrderById)

	const payment = "/order/payment"
	v1.GET(payment+"/:id", middleware.AuthorizeHandlerCookies(), handOrder.GetPayment)
	v1.POST(payment+"/:id", middleware.AuthorizeHandlerCookies(), handOrder.CheckPayment)

	const promo = "/promo"
	v1.GET(promo, middleware.AuthorizeHandlerCookies(), handOrder.GetPromo)
	v1.GET(promo+"/:id", middleware.AuthorizeHandlerCookies(), handOrder.GetPromoById)
	v1.POST(promo, middleware.AuthorizeHandlerCookies(), handOrder.AddPromo)
	v1.PUT(promo+"/:id", middleware.AuthorizeHandlerCookies(), handOrder.EditPromoById)
	v1.DELETE(promo+"/:id", middleware.AuthorizeHandlerCookies(), handOrder.DeletePromoById)

	// Start the server
	r.Run(":8080")
}
