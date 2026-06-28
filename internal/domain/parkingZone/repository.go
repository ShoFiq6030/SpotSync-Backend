package parkingZone

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateParkingZone(zone *ParkingZone) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateParkingZone(zone *ParkingZone) error {
	result := r.db.Create(zone)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
