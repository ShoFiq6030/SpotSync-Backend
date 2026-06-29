package dto

// ReservationUserResponse is the nested user payload returned when reservation user data is preloaded.
type ReservationUserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

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
	User         *ReservationUserResponse `json:"user,omitempty"`
	Zone         *ReservationZoneResponse `json:"zone,omitempty"`
	CreatedAt    string                   `json:"created_at"`
	UpdatedAt    string                   `json:"updated_at"`
}
