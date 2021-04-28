package controllers

import (
	"fyoukuApi/models"
	"github.com/astaxie/beego"
)

type VideoController struct {
	beego.Controller
}

// ChannelAdvert 频道页获取顶部广告
// @router /channel/advert [*]
func (vc *VideoController) ChannelAdvert() {
	channelId, _ := vc.GetInt("channelId")
	if channelId == 0 {
		vc.Data["json"] = ReturnError(4001, "必须指定频道")
		vc.ServeJSON()
	}
	num, videos, err := models.GetChannelAdvert(channelId)
	if err != nil {
		vc.Data["json"] = ReturnError(4004, "请求数据失败")
		vc.ServeJSON()
	} else {
		vc.Data["json"] = ReturnSuccess(0, "success", videos, num)
		vc.ServeJSON()
	}
}
