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

func (s *service) GetParkingZones() ([]dto.ParkingZoneResponse, error) {
	zones, err := s.repo.GetParkingZones()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.ParkingZoneResponse, len(zones))
	for i, zone := range zones {
		responses[i] = zone.ToResponse()
	}

	return responses, nil
}

func (s *service) GetParkingZoneByID(id uint) (*dto.ParkingZoneResponse, error) {
	zone, err := s.repo.GetParkingZoneByID(id)
	if err != nil {
		return nil, err
	}

	if zone == nil {
		return nil, nil
	}

	response := zone.ToResponse()
	return &response, nil
}
