package route

import (
	"github.com/gin-gonic/gin"
	"github.com/xiashaung/demitasse.cn/controller/api"
)

func InitApiRoute(r *gin.Engine) {
	r.GET("/user/add",api.ListUser)
	r.GET("/talent/get",api.TalentInfo)

	//时间api
	r.POST("/time/timestamp",api.ToTimestamp)
	r.GET("/time/now",api.Nowtime)
	r.POST("/time/datetime",api.ToDatetime)
}