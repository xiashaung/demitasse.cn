package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xiashaung/demitasse.cn/model"
	"net/http"
)

func ListUser(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{
		"code": 0,
		"data":"",
	})

}

func TalentInfo(c *gin.Context){
	talentInfo := &model.TalentInfo{}
	model.Table(talentInfo).First(talentInfo)
	c.JSON(http.StatusOK,gin.H{
		"code": 0,
		"data":talentInfo,
	})
}
