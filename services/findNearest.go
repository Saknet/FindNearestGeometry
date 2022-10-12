package services

import (
	"math"

	model "github.com/ForumViriumHelsinki/WPS-FindNearestFeature/models"
)

func FindNearestFeature(data model.FindNearestInput) string {
	closestDistance := 9999999999.9
	closest := data.Feature_data[0].Key

	for i := 0; i < len(data.Feature_data); i++ {

		distance := getDistance(data.Point_Longitude, data.Point_latitude, data.Feature_data[i].Longitude, data.Feature_data[i].Latitude)
		if distance < closestDistance {
			closestDistance = distance
			closest = data.Feature_data[0].Key

		}

	}

	return closest
}

func getDistance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	lon1 := x1 * math.Pi / 180
	lon2 := x2 * math.Pi / 180
	lat1 := y1 * math.Pi / 180
	lat2 := y2 * math.Pi / 180

	dlon := lon2 - lon1
	dlat := lat2 - lat1

	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)

	c := 2 * math.Asin(math.Sqrt(a))

	//  c * Radius of earth in meters

	return c * 6371000

}
