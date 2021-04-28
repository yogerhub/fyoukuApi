package controllers

import (
	"fyoukuApi/models"
	"github.com/astaxie/beego"
	"regexp"
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
		uc.Data["json"] = ReturnSuccess(0, "登录成功", map[string]interface{}{"uid":uid,"username": name},1)
		uc.ServeJSON()
	} else {
		uc.Data["json"] = ReturnError(4004, "手机号或密码不正确")
		uc.ServeJSON()
	}
}
