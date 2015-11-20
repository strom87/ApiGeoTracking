package db

// Image model for image
type Image struct {
	Thumb string `json:"thumb" bson:"thumb"`
	Icon  string `json:"icon"  bson:"icon"`
}
