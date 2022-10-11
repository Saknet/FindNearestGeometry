package main

import (
	"log"
	"net/http"

	mw "github.com/ForumViriumHelsinki/WPS-FindNearestFeature/middleware"

	"github.com/gin-gonic/gin"
)

type FindNearestInput struct {
	Feature_data []struct {
		Latitude  float64
		Longitude float64
		Key       string
	}
	Point_latitude  float64 `json:"point_latitude" binding:"required"`
	Point_Longitude float64 `json:"point_Longitude" binding:"required"`
}

func findNearestFeature(c *gin.Context) {
	// Validate input
	//	var input TimeseriesInput

	var input FindNearestInput

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, FindNearestFeature(input))

}

func main() {

	r := gin.Default()
	r.Use(mw.CORSMiddleware())
	r.POST("/wps/findnearestfeature", findNearestFeature)
	//	obs.GetObservations(dbUrl)

	r.Run("localhost:8087")

}
