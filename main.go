package main

import (
	controller "github.com/forumviriumhelsinki/findnearestgeometry/controllers"
	"github.com/spf13/viper"

	mw "github.com/forumviriumhelsinki/findnearestgeometry/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(mw.CORSMiddleware())
	r.POST("/findnearestgeometry", controller.FindNearestGeometry)

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	r.Run(viper.Get("PATH").(string))

}
