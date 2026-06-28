package parkingZone

import (
	"SpotSync/internal/domain/parkingZone/dto"
	"fmt"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) CreateParkingZone(req dto.CreateParkingZoneRequest) (*dto.ParkingZoneResponse, error) {
	zoneType := ParkingZoneType(req.Type)
	switch zoneType {
	case ZoneGeneral, ZoneEVCharging, ZoneCovered:
	default:
		return nil, fmt.Errorf("invalid parking zone type: %s", req.Type)
	}

	zone := &ParkingZone{
		Name:          req.Name,
		Type:          zoneType,
		TotalCapacity: req.TotalCapacity,
		PricePerHour:  req.PricePerHour,
	}

	if err := s.repo.CreateParkingZone(zone); err != nil {
		return nil, err
	}

	response := zone.ToResponse()
	return &response, nil
}
