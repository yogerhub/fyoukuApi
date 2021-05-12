package controllers

import (
	"fyoukuApi/models"
	"fyoukuApi/services/es"
	"github.com/astaxie/beego"
	"strconv"
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

// ChannelVideo 根据传入参数获取视频列表
// @router /channel/video [*]
func (vc *VideoController) ChannelVideo() {
	//获取频道ID
	channelId, _ := vc.GetInt("channelId")
	//获取频道地区ID
	regionId, _ := vc.GetInt("regionId")
	//获取频道类型ID
	typeId, _ := vc.GetInt("typeId")
	//获取状态
	end := vc.GetString("end")
	//获取排序
	sort := vc.GetString("sort")
	//获取页码
	limit, _ := vc.GetInt("limit")
	offset, _ := vc.GetInt("offset")

	if channelId == 0 {
		vc.Data["json"] = ReturnError(4001, "必须指定频道")
		vc.ServeJSON()
	}
	if limit == 0 {
		limit = 12
	}

	num, videos, err := models.GetChannelVideoList(channelId, regionId, typeId, end, sort, offset, limit)
	if err != nil {
		vc.Data["json"] = ReturnError(4004, "没有相关信息")
		vc.ServeJSON()
	} else {
		vc.Data["json"] = ReturnSuccess(0, "success", videos, num)
		vc.ServeJSON()
	}
}

// VideoInfo 获取视频详情
// @router /video/info [*]
func (vc *VideoController) VideoInfo() {
	videoId, _ := vc.GetInt("videoId")
	if videoId == 0 {
		vc.Data["json"] = ReturnError(4001, "必须指定视频ID")
		vc.ServeJSON()
	}
	video, err := models.RedisGetVideoInfo(videoId)
	if err != nil {
		vc.Data["json"] = ReturnError(4004, "请求数据失败")
		vc.ServeJSON()
	} else {
		vc.Data["json"] = ReturnSuccess(0, "success", video, 1)
		vc.ServeJSON()
	}
}

// VideoEpisodesList 获取视频剧集列表
// @router /video/episodes/list [*]
func (vc *VideoController) VideoEpisodesList() {
	videoId, _ := vc.GetInt("videoId")
	if videoId == 0 {
		vc.Data["json"] = ReturnError(4001, "必须指定视频ID")
		vc.ServeJSON()
	}
	num, episodes, err := models.RedisGetVideoEpisodesList(videoId)
	if err != nil {
		vc.Data["json"] = ReturnError(4004, "请求数据失败")
		vc.ServeJSON()
	} else {
		vc.Data["json"] = ReturnSuccess(0, "success", episodes, num)
		vc.ServeJSON()
	}
}

// UserVideo 我的视频管理
// @router /user/video [*]
func (vc *VideoController) UserVideo() {
	uid, _ := vc.GetInt("uid")
	if uid == 0 {
		vc.Data["json"] = ReturnError(4001, "必须指定用户")
		vc.ServeJSON()
	}
	num, videos, err := models.GetUserVideo(uid)
	if err != nil {
		vc.Data["json"] = ReturnError(4004, "没有相关信息")
		vc.ServeJSON()
	} else {
		vc.Data["json"] = ReturnSuccess(0, "success", videos, num)
		vc.ServeJSON()
	}
}

// VideoSave 保存用户上传视频信息
// @router /video/save [*]
func (vc *VideoController) VideoSave() {
	playUrl := vc.GetString("playUrl")
	title := vc.GetString("title")
	subTitle := vc.GetString("subTitle")
	channelId, _ := vc.GetInt("channelId")
	typeId, _ := vc.GetInt("typeId")
	regionId, _ := vc.GetInt("regionId")
	uid, _ := vc.GetInt("uid")
	if uid == 0 {
		vc.Data["json"] = ReturnError(4001, "必须指定用户")
		vc.ServeJSON()
	}
	if playUrl == "" {
		vc.Data["json"] = ReturnError(4002, "视频地址不能为空")
		vc.ServeJSON()
	}
	err := models.SaveVideo(title, subTitle, channelId, regionId, typeId, playUrl, uid)
	if err != nil {
		vc.Data["json"] = ReturnError(5000, err)
		vc.ServeJSON()
	} else {
		vc.Data["json"] = ReturnSuccess(0, "success", nil, 1)
		vc.ServeJSON()
	}

}

// SendEs 导入ES脚本
// @router /video/send/es [*]
func (vc *VideoController) SendEs() {
	_, data, _ := models.GetAllList()
	for _, v := range data {
		body := map[string]interface{}{
			"id":                   v.Id,
			"title":                v.Title,
			"sub_title":            v.SubTitle,
			"add_time":             v.AddTime,
			"img":                  v.Img,
			"img1":                 v.Img1,
			"episodes_count":       v.EpisodesCount,
			"is_end":               v.IsEnd,
			"channel_id":           v.ChannelId,
			"status":               v.Status,
			"region_id":            v.RegionId,
			"type_id":              v.TypeId,
			"episodes_update_time": v.EpisodesUpdateTime,
			"comment":              v.Comment,
			"user_id":              v.UserId,
			"is_recommend":         v.IsRecommend,
		}
		es.EsAdd("fyouku_video", "video-"+strconv.Itoa(v.Id), body)
	}
	vc.Ctx.WriteString("ES")
}
