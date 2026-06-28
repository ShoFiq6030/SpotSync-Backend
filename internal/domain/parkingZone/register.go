package parkingZone

import (
	"SpotSync/internal/auth"
	"SpotSync/internal/config"
	"SpotSync/internal/middlewares"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	parkingZoneRepository := NewRepository(db)
	parkingZoneService := NewService(parkingZoneRepository)
	parkingZoneHandler := NewHandler(parkingZoneService)
	jwtService := auth.NewJWTService(cfg.JwtSecret)


	api := e.Group("/api/v1/parking-zones")
	api.POST("", parkingZoneHandler.CreateParkingZone, middlewares.AuthMiddleware(jwtService, "ADMIN"))
}


