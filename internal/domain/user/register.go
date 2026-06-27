package user

import (
	"SpotSync/internal/config"

	"github.com/labstack/echo/v5"

	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	userRepository := NewRepository(db)
	
	userService := NewService(userRepository)
	userHandler := NewHandler(userService)

	api := e.Group("/api/v1/auth")

	api.POST("/register", userHandler.CreateUser) // api/v1/auth/register
	

}