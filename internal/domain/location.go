package domain


// Location represents the geographical location of a shuttle.
type Location struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
}

