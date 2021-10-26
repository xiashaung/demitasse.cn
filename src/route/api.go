package route

import (
	"github.com/gin-gonic/gin"
	"github.com/xiashaung/demitasse.cn/controller/api"
)

func InitApiRoute(r *gin.Engine) {
	r.GET("/queue/producer",api.QueueProducer)
	r.GET("/shop/info",api.GetShopInfo)
}