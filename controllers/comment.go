package controllers

import (
	"fyoukuApi/models"
	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

type CommentInfo struct {
	Id int `json:"id"`
	Content string `json:"content"`
	AddTime int64 `json:"addTime"`
	AddTimeTitle string `json:"addTimeTitle"`
	UserId int `json:"UserId"`
	Stamp int `json:"stamp"`
	PraiseCount int `json:"praiseCount"`
	UserInfo models.UserInfo `json:"userInfo"`
}

// List 获取评论列表
// @router /comment/list [*]
func (cc *CommentController) List() {
	//获取剧集数
	episodesId,_ := cc.GetInt("episodesId")
	//获取页码信息
	limit,_ := cc.GetInt("limit")
	offset,_ := cc.GetInt("offset")
	if episodesId == 0{
		cc.Data["json"] = ReturnError(4001,"必须指定视频剧集")
		cc.ServeJSON()
	}
	if limit == 0 {
		limit = 12
	}
	num,comments,err := models.GetCommentList(episodesId,offset,limit)
	if err != nil {
		cc.Data["json"] = ReturnError(4001,"没有相关信息")
		cc.ServeJSON()
	}else {
		var data []CommentInfo
		var commentInfo CommentInfo
		for _,v := range comments{
			commentInfo.Id = v.Id
			commentInfo.Content = v.Content
			commentInfo.AddTime = v.AddTime
			commentInfo.AddTimeTitle = DataFormat(v.AddTime)
			commentInfo.UserId = v.UserId
			commentInfo.Stamp = v.Stamp
			commentInfo.PraiseCount = v.PraiseCount
			//获取用户信息
			commentInfo.UserInfo,_ = models.GetUserInfo(v.UserId)
			data = append(data,commentInfo)
		}
		cc.Data["json"] = ReturnSuccess(0,"success",data,num)
		cc.ServeJSON()
	}
}