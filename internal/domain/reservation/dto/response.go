package dto

// ReservationZoneResponse is the nested zone payload returned when reservation zone data is preloaded.
type ReservationZoneResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// ReservationResponse is the payload returned for reservation endpoints.
type ReservationResponse struct {
	ID           uint                     `json:"id"`
	UserID       uint                     `json:"user_id"`
	ZoneID       uint                     `json:"zone_id"`
	LicensePlate string                   `json:"license_plate"`
	Status       string                   `json:"status"`
	Zone         *ReservationZoneResponse `json:"zone,omitempty"`
	CreatedAt    string                   `json:"created_at"`
	UpdatedAt    string                   `json:"updated_at"`
}
