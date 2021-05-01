package controllers

import (
	"encoding/json"
	"fyoukuApi/models"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
)

type BarrageController struct {
	beego.Controller
}

type WsData struct {
	CurrentTime int
	EpisodesId  int
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// BarrageWs 获取弹幕websocket
// @router /barrage/ws [*]
func (bc *BarrageController) BarrageWs() {
	var (
		conn     *websocket.Conn
		err      error
		data     []byte
		barrages []models.BarrageData
	)
	if conn, err = upgrader.Upgrade(bc.Ctx.ResponseWriter, bc.Ctx.Request, nil); err != nil {
		goto ERR
	}
	for {
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		var wsData WsData
		_ = json.Unmarshal([]byte(data), &wsData)
		endTime := wsData.CurrentTime + 60
		//获取弹幕数据
		_, barrages, err = models.BarrageList(wsData.EpisodesId, wsData.CurrentTime, endTime)
		if err == nil {
			if err := conn.WriteJSON(barrages); err != nil {
				goto ERR
			}
		}

	}
ERR:
	conn.Close()
}

// Save 保存弹幕
// @router /barrage/save [*]
func (bc *BarrageController) Save() {
	uid, _ := bc.GetInt("uid")
	content := bc.GetString("content")
	currentTime, _ := bc.GetInt("currentTime")
	episodesId, _ := bc.GetInt("episodesId")
	videoId, _ := bc.GetInt("videoId")

	if content == "" {
		bc.Data["json"] = ReturnError(4001, "弹幕不能为空")
		bc.ServeJSON()
	}
	if uid == 0 {
		bc.Data["json"] = ReturnError(4002, "请先登录")
		bc.ServeJSON()
	}
	if episodesId == 0 {
		bc.Data["json"] = ReturnError(4003, "必须指定剧集ID")
		bc.ServeJSON()
	}
	if videoId == 0 {
		bc.Data["json"] = ReturnError(4005, "必须指定视频ID")
		bc.ServeJSON()
	}
	if currentTime == 0 {
		bc.Data["json"] = ReturnError(4006, "必须指定视频播放时间")
		bc.ServeJSON()
	}
	err := models.SaveBarrage(episodesId, videoId, currentTime, uid, content)

	if err != nil {
		bc.Data["json"] = ReturnError(5000, err)
		bc.ServeJSON()
	} else {
		bc.Data["json"] = ReturnSuccess(0, "success", "", 1)
		bc.ServeJSON()
	}
}
