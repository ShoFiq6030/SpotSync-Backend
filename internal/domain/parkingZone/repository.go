package parkingZone

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	CreateParkingZone(zone *ParkingZone) error
	GetParkingZones() ([]*ParkingZone, error)
	GetParkingZoneByID(id uint) (*ParkingZone, error)
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

func (r *repository) GetParkingZones() ([]*ParkingZone, error) {
	var zones []*ParkingZone
	result := r.db.Find(&zones)
	if result.Error != nil {
		return nil, result.Error
	}
	return zones, nil
}

func (r *repository) GetParkingZoneByID(id uint) (*ParkingZone, error) {
	var zone ParkingZone
	result := r.db.First(&zone, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &zone, nil
}
