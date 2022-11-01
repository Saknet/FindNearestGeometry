package main

import (
	controller "github.com/ForumViriumHelsinki/WPS-FindNearestFeature/controllers"
	"github.com/spf13/viper"

	mw "github.com/ForumViriumHelsinki/WPS-FindNearestFeature/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(mw.CORSMiddleware())
	r.POST("/findnearestfeature", controller.FindNearestFeature)

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	r.Run(viper.Get("PATH").(string))

}
