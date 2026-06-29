package reservation

import (
	"SpotSync/internal/domain/reservation/dto"
	"errors"

	"gorm.io/gorm"
)

var ErrZoneFull = errors.New("parking zone is full")
var ErrZoneNotFound = errors.New("parking zone not found")
var ErrReservationNotFound = errors.New("reservation not found")
var ErrForbiddenCancellation = errors.New("forbidden to cancel this reservation")
var ErrCanceled = errors.New("Already canceled reservation cannot be canceled again")

type Service interface {
    CreateReservation(userID uint, req dto.CreateReservationRequest) (*dto.ReservationResponse, error)
    GetMyReservations(userID uint) ([]dto.ReservationResponse, error)
    GetAllReservations() ([]dto.ReservationResponse, error)
    CancelReservation(userID uint, role string, reservationID uint) error
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

func (s *service) GetAllReservations() ([]dto.ReservationResponse, error) {
    reservations, err := s.repo.GetAllReservations()
    if err != nil {
        return nil, err
    }

    responses := make([]dto.ReservationResponse, len(reservations))
    for i, reservation := range reservations {
        responses[i] = *reservation.ToResponse()
    }

    return responses, nil
}

func (s *service) CancelReservation(userID uint, role string, reservationID uint) error {
    reservation, err := s.repo.GetReservationByID(reservationID)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return ErrReservationNotFound
        }
        return err
    }

    if role == "DRIVER" && reservation.UserID != userID {
        return ErrForbiddenCancellation
    }
    if reservation.Status == "canceled" {
        return ErrCanceled
    }

    if err := s.repo.CancelReservation(reservationID); err != nil {
        return err
    }

    return nil
}
