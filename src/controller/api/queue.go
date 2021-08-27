package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xiashaung/demitasse.cn/lib"
)

func QueueProducer(c *gin.Context){
	message := c.Query("msg")
	producer := lib.InitProducer("test_exchange","test_queue","test_key")
	producer.Send(message)
}

