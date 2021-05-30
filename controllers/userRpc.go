package controllers

import (
	"context"
	userRpcProto "fyoukuApi/micro/user/proto"
	"fyoukuApi/models"
	"github.com/astaxie/beego"
	"regexp"
)

type UserRpcController struct {
	beego.Controller
}

// LoginDo 用户登录
// @router /login/do [*]
func (uc *UserRpcController) LoginDo(ctx context.Context, req *userRpcProto.RequestLoginDo, res *userRpcProto.ResponseLoginDo) error {
	var (
		userLoginProto userRpcProto.LoginUser
		isMobile       bool
		uid            int
		name           string
	)

	mobile := req.Mobile
	password := req.Password

	if mobile == "" {
		res.Code = 4001
		res.Msg = "手机号不能为空"
		goto ERR
	}
	isMobile, _ = regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !isMobile {
		res.Code = 4002
		res.Msg = "手机号格式不正确"
		goto ERR
	}
	if password == "" {
		res.Code = 4003
		res.Msg = "密码不能为空"
		goto ERR
	}

	uid, name = models.IsMobileLogin(mobile, MD5V(password))
	if uid != 0 {
		userLoginProto.Uid = int64(uid)
		userLoginProto.Username = name

		res.Code = 0
		res.Msg = "登录成功"
		res.Items = &userLoginProto
		res.Count = 1
		return nil
		goto ERR
	} else {
		res.Code = 4004
		res.Msg = "手机号或密码不正确"
		goto ERR
	}

ERR:
	res.Items = &userLoginProto
	res.Count = 0
	return nil
}
