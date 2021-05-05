package controllers

import (
	"fyoukuApi/models"
	"github.com/astaxie/beego"
)

type TopController struct {
	beego.Controller
}

// ChannelTop 根据频道获取排行榜
// @router /channel/top [*]
func (tc *TopController) ChannelTop() {
	//获取频道ID
	channelId, _ := tc.GetInt("channelId")
	if channelId == 0 {
		tc.Data["json"] = ReturnError(4001, "必须指定频道")
		tc.ServeJSON()
	}
	num, videos, err := models.RedisGetChannelTop(channelId)

	if err != nil {
		tc.Data["json"] = ReturnError(4004, "没有相关信息")
		tc.ServeJSON()
	} else {
		tc.Data["json"] = ReturnSuccess(0, "success", videos, num)
		tc.ServeJSON()
	}
}

// TypeTop 根据类型获取排行榜
// @router /type/top [*]
func (tc *TopController) TypeTop() {
	typeId ,_ := tc.GetInt("typeId")
	if typeId  == 0{
		tc.Data["json"] = ReturnError(4001,"必须指定类型")
		tc.ServeJSON()
	}
	num, videos, err := models.RedisGetTypeTop(typeId)

	if err != nil {
		tc.Data["json"] = ReturnError(4004, "没有相关信息")
		tc.ServeJSON()
	} else {
		tc.Data["json"] = ReturnSuccess(0, "success", videos, num)
		tc.ServeJSON()
	}
}