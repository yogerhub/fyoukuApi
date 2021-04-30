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

// ChannelHotList 频道页获取热播视频
// @router /channel/hot [*]
func (vc *VideoController) ChannelHotList() {
	channelId, _ := vc.GetInt("channelId")
	if channelId == 0 {
		vc.Data["json"] = ReturnError(4001, "必须指定频道")
		vc.ServeJSON()
	}

	num, videos, err := models.GetChannelHotList(channelId)
	if err != nil {
		vc.Data["json"] = ReturnError(4004, "没有相关内容")
		vc.ServeJSON()
	} else {
		vc.Data["json"] = ReturnSuccess(0, "success", videos, num)
		vc.ServeJSON()
	}
}

// ChannelRecommendRegionList 频道页-根据频道地区获取推荐的视频列表
// @router /channel/recommend/region [*]
func (vc *VideoController) ChannelRecommendRegionList() {
	channelId, _ := vc.GetInt("channelId")
	regionId, _ := vc.GetInt("regionId")
	if channelId == 0 {
		vc.Data["json"] = ReturnError(4001, "必须指定频道")
		vc.ServeJSON()
	}
	if regionId == 0 {
		vc.Data["json"] = ReturnError(4002, "必须指定频道地区")
		vc.ServeJSON()
	}

	num, videos, err := models.GetChannelRecommendRegionList(channelId, regionId)
	if err != nil {
		vc.Data["json"] = ReturnError(4004, "没有相关内容")
		vc.ServeJSON()
	} else {
		vc.Data["json"] = ReturnSuccess(0, "success", videos, num)
		vc.ServeJSON()
	}
}

// GetChannelRecommendTypeList 频道页-根据频道类型获取推荐视频
// @router /channel/recommend/type [*]
func (vc *VideoController) GetChannelRecommendTypeList() {
	channelId, _ := vc.GetInt("channelId")
	typeId, _ := vc.GetInt("typeId")
	if channelId == 0 {
		vc.Data["json"] = ReturnError(4001, "必须指定频道")
		vc.ServeJSON()
	}
	if typeId == 0 {
		vc.Data["json"] = ReturnError(4002, "必须指定频道类型")
		vc.ServeJSON()
	}
	num, videos, err := models.GetChannelRecommendTypeList(channelId, typeId)
	if err != nil {
		vc.Data["json"] = ReturnError(4004, "没有相关信息")
		vc.ServeJSON()
	} else {
		vc.Data["json"] = ReturnSuccess(0, "success", videos, num)
		vc.ServeJSON()
	}
}
