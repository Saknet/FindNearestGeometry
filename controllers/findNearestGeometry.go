package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	model "github.com/forumviriumhelsinki/findnearestgeometry/models"
	service "github.com/forumviriumhelsinki/findnearestgeometry/services"

	"github.com/gin-gonic/gin"
)

func FindNearestGeometry(c *gin.Context) {

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
