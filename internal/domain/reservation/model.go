package reservation

import (
	"SpotSync/internal/domain/parkingZone"
	"SpotSync/internal/domain/reservation/dto"

	"gorm.io/gorm"
)

type ReservationStatus string

const (
	ReservationActive   ReservationStatus = "active"
	ReservationCanceled ReservationStatus = "canceled"
)

type Reservation struct {
	gorm.Model

	UserID       uint              `json:"user_id" gorm:"not null"`
	ZoneID       uint              `json:"zone_id" gorm:"not null"`
	LicensePlate string            `json:"license_plate" gorm:"type:varchar(15);not null"`
	Status       ReservationStatus `json:"status" gorm:"type:varchar(20);not null;default:'active'"`
	Zone         parkingZone.ParkingZone `json:"zone,omitempty" gorm:"foreignKey:ZoneID"`
}

func (r *Reservation) ToResponse() *dto.ReservationResponse {
	var zone *dto.ReservationZoneResponse
	if r.Zone.ID != 0 {
		zone = &dto.ReservationZoneResponse{
			ID:   r.Zone.ID,
			Name: r.Zone.Name,
			Type: string(r.Zone.Type),
		}
	}

	return &dto.ReservationResponse{
		ID:           r.ID,
		UserID:       r.UserID,
		ZoneID:       r.ZoneID,
		LicensePlate: r.LicensePlate,
		Status:       string(r.Status),
		Zone:         zone,
		CreatedAt:    r.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:    r.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}
