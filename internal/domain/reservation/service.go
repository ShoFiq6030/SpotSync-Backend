package reservation

import (
	"SpotSync/internal/domain/reservation/dto"
	"errors"

	"gorm.io/gorm"
)

var ErrZoneFull = errors.New("parking zone is full")
var ErrZoneNotFound = errors.New("parking zone not found")

type Service interface {
    CreateReservation(userID uint, req dto.CreateReservationRequest) (*dto.ReservationResponse, error)
    GetMyReservations(userID uint) ([]dto.ReservationResponse, error)
}

type service struct {
    repo Repository
    db   *gorm.DB
}

func NewService(repo Repository, db *gorm.DB) *service {
    return &service{repo: repo, db: db}
}

func (s *service) CreateReservation(userID uint, req dto.CreateReservationRequest) (*dto.ReservationResponse, error) {
    reservation, err := s.repo.CreateReservation(userID, req.LicensePlate, req.ZoneID)
    if err != nil {
        return nil, err
    }

    return reservation.ToResponse(), nil
}

func (s *service) GetMyReservations(userID uint) ([]dto.ReservationResponse, error) {
    reservations, err := s.repo.GetMyReservations(userID)
    if err != nil {
        return nil, err
    }

    responses := make([]dto.ReservationResponse, len(reservations))
    for i, reservation := range reservations {
        responses[i] = *reservation.ToResponse()
    }

    return responses, nil
}
