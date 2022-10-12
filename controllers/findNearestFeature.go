package controllers

import (
	"log"
	"net/http"

	model "github.com/ForumViriumHelsinki/WPS-FindNearestFeature/models"
	service "github.com/ForumViriumHelsinki/WPS-FindNearestFeature/services"

	"github.com/gin-gonic/gin"
)

func FindNearestFeature(c *gin.Context) {
	// Validate input
	//	var input TimeseriesInput

	var input model.FindNearestInput

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, service.FindNearestFeature(input))

}
