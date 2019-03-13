package controllers

import (
	"crypto/md5"
	"fmt"
	"gatherInfo/models"
	"github.com/astaxie/beego"
)

type UserLogin struct {
	beego.Controller
}

type UserLogout struct {
	beego.Controller
}

//登录
func (this *UserLogin) Post() {
	var result models.ResultInfo
	loginName := this.GetString("loginname")
	loginpwd := this.GetString("loginpwd")
	if loginpwd == "" {

	}
	userinfo, err := models.Login(loginName, fmt.Sprintf("%x", md5.Sum([]byte("loginpwd"))))
	if err != nil {
		result.ErrCode = -1
		result.ErrMsg = err.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	userinfo.Loginpwd = ""
	if userinfo.Id > 0 {
		this.SetSession("userinfo", userinfo) //加入session管理
	}
	result.ErrCode = 0
	result.ErrMsg = ""
	result.Result = userinfo
	this.Data["json"] = result
	this.ServeJSON()
	return
}

//登出
func (this *UserLogout) Get() {
	this.SetSession("userinfo", nil)
	var result models.ResultInfo
	result.ErrCode = 0
	result.ErrMsg = "退出登录成功"
	this.Data["json"] = result
	this.ServeJSON()
}
