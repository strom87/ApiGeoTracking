package models

// GeoLocationModel struct
type GeoLocationModel struct {
	ID           string  `json:"string"`
	Lat          float64 `json:"lat"`
	Long         float64 `json:"long"`
	Disconnected bool    `json:"disconnected"`
}
