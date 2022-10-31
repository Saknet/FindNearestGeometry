package main

import (
	controller "github.com/ForumViriumHelsinki/WPS-FindNearestFeature/controllers"

	mw "github.com/ForumViriumHelsinki/WPS-FindNearestFeature/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(mw.CORSMiddleware())
	r.POST("/findnearestfeature", controller.FindNearestFeature)

	r.Run("localhost:8087")

}
