package service

import (
	"github.com/xiashaung/demitasse.cn/model"
)

type ShopInfo struct {

}

func (ShopInfo) GetById(shopId int)  (model.ShopInfo) {
	var shopInfo model.ShopInfo
	model.Table(model.ShopInfo{}).First(&shopInfo,shopId)
	return shopInfo
}
