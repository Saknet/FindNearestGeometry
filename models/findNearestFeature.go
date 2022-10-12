package models

type FindNearestInput struct {
	Feature_data []struct {
		Latitude  float64
		Longitude float64
		Key       string
	}
	Point_latitude  float64 `json:"point_latitude" binding:"required"`
	Point_Longitude float64 `json:"point_Longitude" binding:"required"`
}
