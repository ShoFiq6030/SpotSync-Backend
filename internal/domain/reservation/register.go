package reservation

import (
	"SpotSync/internal/auth"
	"SpotSync/internal/config"
	"SpotSync/internal/middlewares"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
    reservationRepository := NewRepository(db)
    reservationService := NewService(reservationRepository, db)
    reservationHandler := NewHandler(reservationService)
    jwtService := auth.NewJWTService(cfg.JwtSecret)

    api := e.Group("/api/v1/reservations")
    api.POST("", reservationHandler.CreateReservation, middlewares.AuthMiddleware(jwtService, "DRIVER", "ADMIN"))
    api.GET("/my-reservations", reservationHandler.GetMyReservations, middlewares.AuthMiddleware(jwtService, "DRIVER", "ADMIN"))
}
