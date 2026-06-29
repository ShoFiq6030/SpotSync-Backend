package parkingZone

import (
	"SpotSync/internal/domain/parkingZone/dto"
	"SpotSync/internal/httpresponse"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateParkingZone(c *echo.Context) error {
	var req dto.CreateParkingZoneRequest

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

	response, err := h.service.CreateParkingZone(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: "Failed to create parking zone",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, httpresponse.Success{
		Success: true,
		Code:    http.StatusCreated,
		Message: "Parking zone created successfully",
		Data:    response,
	})
}

func (h *handler) GetParkingZones(c *echo.Context) error {
	zones, err := h.service.GetParkingZones()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve parking zones",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, httpresponse.Success{
		Success: true,
		Code:    http.StatusOK,
		Message: "Parking zones retrieved successfully",
		Data:    zones,
	})
}

func (h *handler) GetParkingZoneByID(c *echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "Invalid zone ID",
			Details: err.Error(),
		})
	}

	zone, err := h.service.GetParkingZoneByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve parking zone",
			Details: err.Error(),
		})
	}

	if zone == nil {
		return c.JSON(http.StatusNotFound, httpresponse.Error{
			Success: false,
			Code:    http.StatusNotFound,
			Message: "Parking zone not found",
		})
	}

	return c.JSON(http.StatusOK, httpresponse.Success{
		Success: true,
		Code:    http.StatusOK,
		Message: "Parking zone retrieved successfully",
		Data:    zone,
	})
}
