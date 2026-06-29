package dto

type CreateReservationRequest struct {
    ZoneID       uint   `json:"zone_id" validate:"required,gt=0"`
    LicensePlate string `json:"license_plate" validate:"required,min=1,max=15"`
}