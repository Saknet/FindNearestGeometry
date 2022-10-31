package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	model "github.com/ForumViriumHelsinki/WPS-FindNearestFeature/models"
	service "github.com/ForumViriumHelsinki/WPS-FindNearestFeature/services"

	"github.com/gin-gonic/gin"
)

func FindNearestFeature(c *gin.Context) {

	data, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Println(err)
	}

	var input model.GeoJSONPointRequest

	if err := json.Unmarshal(data, &input); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, service.FindFromGeoJSONPoint(input))

}
