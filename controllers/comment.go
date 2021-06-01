package controllers

import (
	"fmt"
	"fyoukuApi/models"
	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

type CommentInfo struct {
	Id           int             `json:"id"`
	Content      string          `json:"content"`
	AddTime      int64           `json:"addTime"`
	AddTimeTitle string          `json:"addTimeTitle"`
	UserId       int             `json:"UserId"`
	Stamp        int             `json:"stamp"`
	PraiseCount  int             `json:"praiseCount"`
	UserInfo     models.UserInfo `json:"userInfo"`
}

//// List 获取评论列表
//// @router /comment/list [*]
//func (cc *CommentController) List() {
//	//获取剧集数
//	episodesId,_ := cc.GetInt("episodesId")
//	//获取页码信息
//	limit,_ := cc.GetInt("limit")
//	offset,_ := cc.GetInt("offset")
//	if episodesId == 0{
//		cc.Data["json"] = ReturnError(4001,"必须指定视频剧集")
//		cc.ServeJSON()
//	}
//	if limit == 0 {
//		limit = 12
//	}
//	num,comments,err := models.GetCommentList(episodesId,offset,limit)
//	if err != nil {
//		cc.Data["json"] = ReturnError(4001,"没有相关信息")
//		cc.ServeJSON()
//	}else {
//		var data []CommentInfo
//		var commentInfo CommentInfo
//		for _,v := range comments{
//			commentInfo.Id = v.Id
//			commentInfo.Content = v.Content
//			commentInfo.AddTime = v.AddTime
//			commentInfo.AddTimeTitle = DataFormat(v.AddTime)
//			commentInfo.UserId = v.UserId
//			commentInfo.Stamp = v.Stamp
//			commentInfo.PraiseCount = v.PraiseCount
//			//获取用户信息
//			commentInfo.UserInfo,_ = models.RedisGetUserInfo(v.UserId)
//			data = append(data,commentInfo)
//		}
//		cc.Data["json"] = ReturnSuccess(0,"success",data,num)
//		cc.ServeJSON()
//	}
//}


// List 获取评论列表
// @router /comment/list [*]
func (cc *CommentController) List() {
	//获取剧集数
	episodesId, _ := cc.GetInt("episodesId")
	//获取页码信息
	limit, _ := cc.GetInt("limit")
	offset, _ := cc.GetInt("offset")
	if episodesId == 0 {
		cc.Data["json"] = ReturnError(4001, "必须指定视频剧集")
		cc.ServeJSON()
		return
	}
	if limit == 0 {
		limit = 12
	}
	num, comments, err := models.GetCommentList(episodesId, offset, limit)
	fmt.Println(comments)
	if err != nil {
		cc.Data["json"] = ReturnError(4001, "没有相关信息")
		cc.ServeJSON()
		return
	} else {
		var data []CommentInfo
		var commentInfo CommentInfo

		//获取uid channel
		uidChan := make(chan int, 12)
		closeChan := make(chan bool, 12)
		resChan := make(chan models.UserInfo, 12)
		//把获取到到uid放到channel中
		go func() {
			for _, v := range comments {
				uidChan <- v.UserId
			}
			close(uidChan)
		}()
		//处理uidChan中到信息
		for i := 0; i < 12; i++ {
			go chanGetUserInfo(uidChan, resChan, closeChan)
		}
		//判断是否执行完成，信息聚合
		go func() {
			for i := 0; i < 12; i++ {
				<-closeChan
			}
			close(resChan)
			close(closeChan)
		}()
		userInfoMap := make(map[int]models.UserInfo)
		for r := range resChan {
			userInfoMap[r.Id] = r
		}
		for _, v := range comments {
			commentInfo.Id = v.Id
			commentInfo.Content = v.Content
			commentInfo.AddTime = v.AddTime
			commentInfo.AddTimeTitle = DataFormat(v.AddTime)
			commentInfo.UserId = v.UserId
			commentInfo.Stamp = v.Stamp
			commentInfo.PraiseCount = v.PraiseCount
			//获取用户信息
			commentInfo.UserInfo, _ = userInfoMap[v.UserId]
			data = append(data, commentInfo)
		}

		cc.Data["json"] = ReturnSuccess(0, "success", data, num)
		cc.ServeJSON()
		return
	}
}


func chanGetUserInfo(uidChan chan int, resChan chan models.UserInfo, closeChan chan bool) {
	for uid := range uidChan {
		res, err := models.RedisGetUserInfo(uid)
		fmt.Println(res)
		if err == nil {
			resChan <- res
		}
		closeChan <- true
	}
}

// Save 保存评论
// @router /comment/save [*]
func (cc *CommentController) Save() {
	content := cc.GetString("content")
	uid, _ := cc.GetInt("uid")
	episodesId, _ := cc.GetInt("episodesId")
	videoId, _ := cc.GetInt("videoId")

	if content == "" {
		cc.Data["json"] = ReturnError(4001, "内容不能为空")
		cc.ServeJSON()
		return
	}
	if uid == 0 {
		cc.Data["json"] = ReturnError(4002, "请先登录")
		cc.ServeJSON()
		return
	}
	if episodesId == 0 {
		cc.Data["json"] = ReturnError(4003, "必须指定评论剧集ID")
		cc.ServeJSON()
		return
	}
	if videoId == 0 {
		cc.Data["json"] = ReturnError(4005, "必须指定视频ID")
		cc.ServeJSON()
		return
	}
	err := models.SaveComment(content, uid, episodesId, videoId)
	if err != nil {
		cc.Data["json"] = ReturnError(5000, err)
		cc.ServeJSON()
		return
	} else {
		cc.Data["json"] = ReturnSuccess(0, "success", "", 1)
		cc.ServeJSON()
		return
	}
}
