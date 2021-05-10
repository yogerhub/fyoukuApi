package controllers

import (
	"fmt"
	"fyoukuApi/models"
	"github.com/astaxie/beego"
	"regexp"
	"strconv"
	"strings"
)

type UserController struct {
	beego.Controller
}

// SaveRegister 用户注册功能
// @router /register/save [post]
func (uc *UserController) SaveRegister() {
	var (
		mobile   string
		password string
	)
	mobile = uc.GetString("mobile")
	password = uc.GetString("password")

	if mobile == "" {
		uc.Data["json"] = ReturnError(4001, "手机号不能为空")
		uc.ServeJSON()
	}
	isMobile, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !isMobile {
		uc.Data["json"] = ReturnError(4002, "手机号格式不正确")
		uc.ServeJSON()
	}
	if password == "" {
		uc.Data["json"] = ReturnError(4003, "密码不能为空")
		uc.ServeJSON()
	}
	//判断手机号是否注册
	status := models.IsUserMobile(mobile)
	if status {
		uc.Data["json"] = ReturnError(4005, "此手机号已经注册")
		uc.ServeJSON()
	} else {
		//保存用户
		err := models.UserSave(mobile, MD5V(password))

		if err != nil {
			uc.Data["json"] = ReturnError(5000, err)
			uc.ServeJSON()
		} else {
			uc.Data["json"] = ReturnSuccess(0, "注册成功", nil, 0)
			uc.ServeJSON()
		}
	}
}

// LoginDo 用户登录
// @router /login/do [*]
func (uc *UserController) LoginDo() {
	mobile := uc.GetString("mobile")
	password := uc.GetString("password")

	if mobile == "" {
		uc.Data["json"] = ReturnError(4001, "手机号不能为空")
		uc.ServeJSON()
	}
	isMobile, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !isMobile {
		uc.Data["json"] = ReturnError(4002, "手机号格式不正确")
		uc.ServeJSON()
	}
	if password == "" {
		uc.Data["json"] = ReturnError(4003, "密码不能为空")
		uc.ServeJSON()
	}

	uid, name := models.IsMobileLogin(mobile, MD5V(password))
	if uid != 0 {
		uc.Data["json"] = ReturnSuccess(0, "登录成功", map[string]interface{}{"uid": uid, "username": name}, 1)
		uc.ServeJSON()
	} else {
		uc.Data["json"] = ReturnError(4004, "手机号或密码不正确")
		uc.ServeJSON()
	}
}

//// SendMessageDo 批量发送通知消息
//// @router /send/message [*]
//func (uc *UserController) SendMessageDo() {
//	uids := uc.GetString("uids")
//	content := uc.GetString("content")
//	if uids == "" {
//		uc.Data["json"] = ReturnError(4001, "请输入接收人")
//		uc.ServeJSON()
//	}
//	if content == "" {
//		uc.Data["json"] = ReturnError(4002, "请填写发送内容")
//		uc.ServeJSON()
//	}
//	messageId, err := models.SendMessageDo(content)
//	if err != nil {
//		uc.Data["json"] = ReturnError(5000, "发送失败")
//		uc.ServeJSON()
//	} else {
//		uidConfig := strings.Split(uids, ",")
//		for _, v := range uidConfig {
//			userId, _ := strconv.Atoi(v)
//			//err = models.SendMessageUser(userId, messageId)
//			models.SendMessageUserMq(userId, messageId)
//		}
//		uc.Data["json"] = ReturnSuccess(0, "发送成功", "", 1)
//		uc.ServeJSON()
//	}
//}

type SendData struct {
	UserId    int
	MessageId int64
}

// SendMessageDo 批量发送通知消息
// @router /send/message [*]
func (uc *UserController) SendMessageDo() {
	uids := uc.GetString("uids")
	content := uc.GetString("content")
	if uids == "" {
		uc.Data["json"] = ReturnError(4001, "请输入接收人")
		uc.ServeJSON()
	}
	if content == "" {
		uc.Data["json"] = ReturnError(4002, "请填写发送内容")
		uc.ServeJSON()
	}
	messageId, err := models.SendMessageDo(content)
	if err != nil {
		uc.Data["json"] = ReturnError(5000, "发送失败")
		uc.ServeJSON()
	} else {
		uidConfig := strings.Split(uids, ",")
		count := len(uidConfig)
		sendChan := make(chan SendData, count)
		closeChan := make(chan bool, count)
		go func() {
			var data SendData
			for _, v := range uidConfig {
				userId, _ := strconv.Atoi(v)
				data.UserId = userId
				data.MessageId = messageId
				sendChan <- data
			}
			close(sendChan)
		}()
		for i := 0; i < count; i++ {
			go sendMessageFunc(sendChan, closeChan)
		}

		for i := 0; i < count; i++ {
			<-closeChan
		}
		close(closeChan)

		uc.Data["json"] = ReturnSuccess(0, "发送成功", "", 1)
		uc.ServeJSON()
	}
}

func sendMessageFunc(sendChan chan SendData, closeChan chan bool) {
	for t := range sendChan {
		fmt.Println(t)
		models.SendMessageUserMq(t.UserId, t.MessageId)
	}
	closeChan <- true
}
