package reservation

import (
	"SpotSync/internal/domain/reservation/dto"
	"SpotSync/internal/httpresponse"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
    service Service
}

func NewHandler(service Service) *handler {
    return &handler{service: service}
}

func (h *handler) CreateReservation(c *echo.Context) error {
    claims := c.Get("user_id")
    userID, ok := claims.(uint)
    if !ok {
        return c.JSON(http.StatusUnauthorized, httpresponse.Error{
            Success: false,
            Code:    http.StatusUnauthorized,
            Message: "Unauthorized",
            Details: "user context is missing",
        })
    }

    var req dto.CreateReservationRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, httpresponse.Error{
            Success: false,
            Code:    http.StatusBadRequest,
            Message: "Invalid request payload",
            Details: err.Error(),
        })
    }

    if err := c.Validate(&req); err != nil {
        return c.JSON(http.StatusBadRequest, httpresponse.Error{
            Success: false,
            Code:    http.StatusBadRequest,
            Message: "Validation failed",
            Details: err.Error(),
        })
    }

    reservation, err := h.service.CreateReservation(userID, req)
    if err != nil {
        switch err {
        case ErrZoneFull:
            return c.JSON(http.StatusUnprocessableEntity, httpresponse.Error{
                Success: false,
                Code:    http.StatusUnprocessableEntity,
                Message: "Parking zone is full",
                Details: err.Error(),
            })
        case ErrZoneNotFound:
            return c.JSON(http.StatusNotFound, httpresponse.Error{
                Success: false,
                Code:    http.StatusNotFound,
                Message: "Parking zone not found",
                Details: err.Error(),
            })
        default:
            return c.JSON(http.StatusInternalServerError, httpresponse.Error{
                Success: false,
                Code:    http.StatusInternalServerError,
                Message: "Failed to create reservation",
                Details: err.Error(),
            })
        }
    }

    return c.JSON(http.StatusCreated, httpresponse.Success{
        Success: true,
        Code:    http.StatusCreated,
        Message: "Reservation confirmed successfully",
        Data:    reservation,
    })
}
