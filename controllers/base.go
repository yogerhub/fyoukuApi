package controllers

import (
	"fyoukuApi/models"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

// ChannelRegion 获取频道地区列表
// @router /channel/region [*]
func (bc *BaseController) ChannelRegion() {
	channelId, _ := bc.GetInt("channelId")
	if channelId == 0 {
		bc.Data["json"] = ReturnError(4001, "必须指定频道")
		bc.ServeJSON()
	}
	num, regions, err := models.GetChannelRegion(channelId)
	if err != nil {
		bc.Data["json"] = ReturnError(4004, "没有相关信息")
		bc.ServeJSON()
	} else {
		bc.Data["json"] = ReturnSuccess(0, "success", regions, num)
		bc.ServeJSON()
	}
}

// ChannelType 获取频道类型列表
// @router /channel/type [*]
func (bc *BaseController) ChannelType() {
	channelId, _ := bc.GetInt("channelId")
	if channelId == 0 {
		bc.Data["json"] = ReturnError(4001, "必须指定频道")
		bc.ServeJSON()
	}
	num, types, err := models.GetChannelType(channelId)
	if err != nil {
		bc.Data["json"] = ReturnError(4004, "没有相关信息")
		bc.ServeJSON()
	} else {
		bc.Data["json"] = ReturnSuccess(0, "success", types, num)
		bc.ServeJSON()
	}

}
