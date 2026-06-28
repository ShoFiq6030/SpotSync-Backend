package parkingZone

import (
	"SpotSync/internal/domain/parkingZone/dto"

	"gorm.io/gorm"
)

type ParkingZoneType string

const (
	ZoneGeneral    ParkingZoneType = "general"
	ZoneEVCharging ParkingZoneType = "ev_charging"
	ZoneCovered    ParkingZoneType = "covered"
)

type ParkingZone struct {
	gorm.Model

	Name          string  `json:"name" gorm:"type:varchar(150);not null"`
	Type          ParkingZoneType   `json:"type" gorm:"type:varchar(30);not null"`
	TotalCapacity int     `json:"total_capacity" gorm:"not null"`
	PricePerHour  float64 `json:"price_per_hour" gorm:"type:decimal(10,2);not null"`
}

func ( p *ParkingZone) ToResponse() dto.ParkingZoneResponse {
	return dto.ParkingZoneResponse{
		ID:              p.ID,
		Name:            p.Name,
		Type:            string(p.Type),
		TotalCapacity:   p.TotalCapacity,
		AvailableSpots:  p.TotalCapacity, // Assuming all spots are available initially
		PricePerHour:    p.PricePerHour,
		CreatedAt:       p.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:       p.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}