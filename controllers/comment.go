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

// Save 保存评论
// @router /comment/save [*]
func (cc *CommentController) Save() {
	content := cc.GetString("content")
	uid,_ := cc.GetInt("uid")
	episodesId,_ := cc.GetInt("episodesId")
	videoId,_ := cc.GetInt("videoId")

	if content == "" {
		cc.Data["json"] = ReturnError(4001,"内容不能为空")
		cc.ServeJSON()
	}
	if uid == 0 {
		cc.Data["json"] = ReturnError(4002,"请先登录")
		cc.ServeJSON()
	}
	if episodesId == 0 {
		cc.Data["json"] = ReturnError(4003,"必须指定评论剧集ID")
		cc.ServeJSON()
	}
	if videoId == 0 {
		cc.Data["json"] = ReturnError(4005,"必须指定视频ID")
		cc.ServeJSON()
	}
	err := models.SaveComment(content,uid,episodesId,videoId)
	if err != nil {
		cc.Data["json"] = ReturnError(5000,err)
		cc.ServeJSON()
	}else {
		cc.Data["json"] = ReturnSuccess(0,"success","",1)
		cc.ServeJSON()
	}
}