package reservation

import (
	"SpotSync/internal/domain/reservation/dto"
	"SpotSync/internal/httpresponse"
	"net/http"
	"strconv"

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

func (h *handler) CancelReservation(c *echo.Context) error {
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

    roleClaim := c.Get("user_role")
    userRole, ok := roleClaim.(string)
    if !ok {
        return c.JSON(http.StatusUnauthorized, httpresponse.Error{
            Success: false,
            Code:    http.StatusUnauthorized,
            Message: "Unauthorized",
            Details: "user role context is missing",
        })
    }

    reservationIDParam := c.Param("id")
    reservationIDUint64, err := strconv.ParseUint(reservationIDParam, 10, 64)
    if err != nil {
        return c.JSON(http.StatusBadRequest, httpresponse.Error{
            Success: false,
            Code:    http.StatusBadRequest,
            Message: "Invalid reservation ID",
            Details: err.Error(),
        })
    }

    if err := h.service.CancelReservation(userID, userRole, uint(reservationIDUint64)); err != nil {
        switch err {
        case ErrReservationNotFound:
            return c.JSON(http.StatusNotFound, httpresponse.Error{
                Success: false,
                Code:    http.StatusNotFound,
                Message: "Reservation not found",
                Details: err.Error(),
            })
        case ErrForbiddenCancellation:
            return c.JSON(http.StatusForbidden, httpresponse.Error{
                Success: false,
                Code:    http.StatusForbidden,
                Message: "Forbidden",
                Details: err.Error(),
            })
        case ErrCanceled:
            return c.JSON(http.StatusConflict, httpresponse.Error{
                Success: false,
                Code:    http.StatusConflict,
                Message: "Conflict",
                Details: err.Error(),
            })
        default:
            return c.JSON(http.StatusInternalServerError, httpresponse.Error{
                Success: false,
                Code:    http.StatusInternalServerError,
                Message: "Failed to cancel reservation",
                Details: err.Error(),
            })
        }
    }

    return c.JSON(http.StatusOK, httpresponse.Success{
        Success: true,
        Code:    http.StatusOK,
        Message: "Reservation cancelled successfully",
    })
}

func (h *handler) GetMyReservations(c *echo.Context) error {
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

    reservations, err := h.service.GetMyReservations(userID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, httpresponse.Error{
            Success: false,
            Code:    http.StatusInternalServerError,
            Message: "Failed to retrieve reservations",
            Details: err.Error(),
        })
    }

    return c.JSON(http.StatusOK, httpresponse.Success{
        Success: true,
        Code:    http.StatusOK,
        Message: "My reservations retrieved successfully",
        Data:    reservations,
    })
}

func (h *handler) GetAllReservations(c *echo.Context) error {
    reservations, err := h.service.GetAllReservations()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, httpresponse.Error{
            Success: false,
            Code:    http.StatusInternalServerError,
            Message: "Failed to retrieve reservations",
            Details: err.Error(),
        })
    }

    return c.JSON(http.StatusOK, httpresponse.Success{
        Success: true,
        Code:    http.StatusOK,
        Message: "All reservations retrieved successfully",
        Data:    reservations,
    })
}

