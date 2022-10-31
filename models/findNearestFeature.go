package models

type FindNearestInput struct {
	Feature_data []struct {
		Latitude  float64
		Longitude float64
		Key       string
	}
	Point_latitude  float64 `json:"point_latitude" binding:"required"`
	Point_longitude float64 `json:"point_longitude" binding:"required"`
}

type GeoJSONPointRequest struct {
	Features        []GeometryPoint `json:"features" binding:"required"`
	Point_latitude  float64         `json:"point_latitude" binding:"required"`
	Point_longitude float64         `json:"point_longitude" binding:"required"`
}

type GeoJSONPointResponse struct {
	Nearest GeometryPoint `json:"nearest"`
}

type FindNearestResponse struct {
	Nearest string `json:"nearest"`
}

type GeometryPoint struct {
	Type        string    `json:"type,omitempty"`
	Coordinates []float64 `json:"coordinates,omitempty"`
}

type GeometryLineString struct {
	Type        string      `json:"type,omitempty"`
	Coordinates [][]float64 `json:"coordinates,omitempty"`
}

type GeometryPolygon struct {
	Type        string        `json:"type,omitempty"`
	Coordinates [][][]float64 `json:"coordinates,omitempty"`
}

type GeometryMultiPoly struct {
	Type        string          `json:"type,omitempty"`
	Coordinates [][][][]float64 `json:"coordinates,omitempty"`
}
