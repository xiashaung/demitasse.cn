package route

import (
	"github.com/gin-gonic/gin"
	"github.com/xiashaung/demitasse.cn/controller/api"
)

func InitApiRoute(r *gin.Engine) {
	//时间api
	timeApi := r.Group("/time")
	timeApi.POST("/timestamp",api.ToTimestamp)
	timeApi.GET("/now",api.Nowtime)
	timeApi.POST("/datetime",api.ToDatetime)

	//微信授权接口
	mpApi := r.Group("/account/")
	mpApi.POST("wx-auth",api.MpAuth)

	r.GET("/queue/producer",api.QueueProducer)
}