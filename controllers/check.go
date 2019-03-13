package controllers

import (
	"fmt"
	"gatherInfo/models"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type GetRealnameList struct {
	beego.Controller
}

type AddRealnameInfo struct {
	beego.Controller
}

type DelRealnameInfo struct {
	beego.Controller
}

//获取实名制列表
func (this *GetRealnameList) Get() {
	var result models.ResultInfo
	var strWhere string
	_, err := checkLogin(this.GetSession("userinfo"))
	if err != nil {
		result.ErrCode = -5 //-5表示未登录状态
		result.ErrMsg = err.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	checkinfoId, err0 := this.GetInt("checkinfoId")
	if err0 != nil {
		result.ErrCode = -17
		result.ErrMsg = "查询实名制记录获取参数失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	if checkinfoId > 0 {
		strWhere += " checkinfoid = " + strconv.Itoa(checkinfoId)
	}
	realList, err1 := models.GetRealnameList(strWhere)
	if err1 != nil {
		result.ErrCode = -18
		result.ErrMsg = "查询实名制记录失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	result.ErrCode = 0
	result.ErrMsg = ""
	result.Result = realList
	this.Data["json"] = result
	this.ServeJSON()
	return
}

//添加一条实名制检查记录
func (this *AddRealnameInfo) Post() {
	var result models.ResultInfo
	_, err := checkLogin(this.GetSession("userinfo"))
	if err != nil {
		result.ErrCode = -5 //-5表示未登录状态
		result.ErrMsg = err.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	checkinfoId, err0 := this.GetInt("checkinfoId")
	Interviewername := this.GetString("Interviewername")
	Interviewerbirth := this.GetString("Interviewerbirth")
	Reason := this.GetString("Reason")
	if err0 != nil {
		result.ErrCode = -15
		result.ErrMsg = "添加实名制记录获取参数失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	var realInfo models.Realname
	realInfo.Checkinfoid = checkinfoId
	realInfo.Interviewername = Interviewername
	realInfo.Interviewerbirth = Interviewerbirth
	realInfo.Reason = Reason
	realInfo.Isactive = 1
	realInfo.Createtime = time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(checkinfoId)
	fmt.Println(Interviewername)
	fmt.Println(Interviewerbirth)
	fmt.Println(Reason)
	err1 := models.AddRealnameInfo(realInfo)
	if err1 != nil {
		result.ErrCode = -16
		result.ErrMsg = err1.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	result.ErrCode = 0
	result.ErrMsg = ""
	this.Data["json"] = result
	this.ServeJSON()
	return
}

//删除一条实名制记录
func (this *DelRealnameInfo) Post() {
	var result models.ResultInfo
	_, err := checkLogin(this.GetSession("userinfo"))
	if err != nil {
		result.ErrCode = -5 //-5表示未登录状态
		result.ErrMsg = err.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	id, err0 := this.GetInt("id")
	if err0 != nil {
		result.ErrCode = -19
		result.ErrMsg = "删除实名制记录获取参数失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	err1 := models.DelRealnameInfo(id)
	if err1 != nil {
		result.ErrCode = -20
		result.ErrMsg = "删除实名制记录内部错误"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	result.ErrCode = 0
	result.ErrMsg = ""
	this.Data["json"] = result
	this.ServeJSON()
	return
}
