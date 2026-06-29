package reservation

import (
	"SpotSync/internal/domain/parkingZone"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
    CreateReservation(userId uint, licensePlate string, zoneID uint)(*Reservation, error)
   
}

type repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
    return &repository{db: db}
}

func (r *repository) CreateReservation( userId uint, licensePlate string, zoneID uint) (*Reservation, error) {
   
    var reservation Reservation

    err := r.db.Transaction(func(tx *gorm.DB) error {
          var zone parkingZone.ParkingZone
            // 1. Lock the row!
		 if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&zone, zoneID).Error; err != nil {
        return err
    }
    // 2. Count current 'active' reservations for this zone
    var activeReservationsCount int64
    if err := tx.Model(&Reservation{}).Where("zone_id = ? AND status = ?", zoneID, "active").Count(&activeReservationsCount).Error; err != nil {
        return err
    }

     // 3. Check if active_count < zone.total_capacity
    if activeReservationsCount >= int64(zone.TotalCapacity) {
        return ErrZoneFull
    }

    // 4. If yes, create reservation. If no, return custom error (e.g., ErrZoneFull).
    reservation = Reservation{
        UserID:       userId,
        LicensePlate: licensePlate,
        ZoneID:       zoneID,
        Status:       "active",
    }
    if err := tx.Create(&reservation).Error; err != nil {
        return err
    }
		return nil

	})

    if err != nil {
        return nil, err
    }

    return &reservation, nil

}

