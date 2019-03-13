package controllers

import (
	"errors"
	"gatherInfo/models"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type GetTaskList struct {
	beego.Controller
}

type AddTaskInfo struct {
	beego.Controller
}

type UpdateTaskInfo struct {
	beego.Controller
}

type DelTaskInfo struct {
	beego.Controller
}

//获取任务列表
func (this *GetTaskList) Get() {
	var result models.ResultInfo
	var strWhere string
	userinfo, err := checkLogin(this.GetSession("userinfo"))
	if err != nil {
		result.ErrCode = -5 //-5表示未登录状态
		result.ErrMsg = err.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	if userinfo.Id > 0 {
		if userinfo.Usertype != 1 {
			strWhere += " investigatorid = " + strconv.Itoa(userinfo.Id)
		}
	}
	orderState, err1 := this.GetInt("orderstate")

	if err1 != nil {
		result.ErrCode = -6
		result.ErrMsg = "获取任务列表时获取参数失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	if orderState != 3 { //前台传递3代表查询全部状态的任务,1代表未办结任务,2代表已办结任务
		if orderState == 1 {
			if len(strWhere) > 0 {
				strWhere += " and "
			}
			strWhere += " orderstate = 1 "
		} else if orderState == 2 {
			if len(strWhere) > 0 {
				strWhere += " and "
			}
			strWhere += " orderstate = 2 "
		}
	}
	taskList, err2 := models.GetTaskList(strWhere)
	if err2 != nil {
		result.ErrCode = -2
		result.ErrMsg = "获取任务列表失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	result.ErrCode = 0
	result.ErrMsg = ""
	result.Result = taskList
	this.Data["json"] = result
	this.ServeJSON()
	return
}

//添加一个新任务
func (this *AddTaskInfo) Post() {
	var result models.ResultInfo
	_, err := checkLogin(this.GetSession("userinfo"))
	if err != nil {
		result.ErrCode = -5 //-5表示未登录状态
		result.ErrMsg = err.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	regionId, err1 := this.GetInt("regionId")
	enterpriseId, err2 := this.GetInt("enterpriseId")
	visitId, err3 := this.GetInt("visitId")
	if err1 != nil || err2 != nil || err3 != nil {
		result.ErrCode = -7
		result.ErrMsg = "添加新任务获取参数失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	var taskInfo models.Checkinfo
	taskInfo.Communityid = regionId
	taskInfo.Internetbarid = enterpriseId
	taskInfo.Investigatorid = visitId
	taskInfo.Isactive = 1
	taskInfo.Orderstate = 1
	taskInfo.Createtime = time.Now().Format("2006-01-02 15:04:05")
	taskInfo.Updatetime = time.Now().Format("2006-01-02 15:04:05")
	err4 := models.AddTaskInfo(taskInfo)
	if err4 != nil {
		result.ErrCode = -8
		result.ErrMsg = err4.Error()
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

//修改任务信息
func (this *UpdateTaskInfo) Post() {
	var result models.ResultInfo
	_, err := checkLogin(this.GetSession("userinfo"))
	if err != nil {
		result.ErrCode = -5 //-5表示未登录状态
		result.ErrMsg = err.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	id, err0 := this.GetInt("orderId")
	regionId, err1 := this.GetInt("regionId")
	enterpriseId, err2 := this.GetInt("enterpriseId")
	visitId, err3 := this.GetInt("visitId")
	if err0 != nil || err1 != nil || err2 != nil || err3 != nil {
		result.ErrCode = -9
		result.ErrMsg = "修改任务获取参数失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	var taskInfo models.Checkinfo
	taskInfo.Communityid = regionId
	taskInfo.Internetbarid = enterpriseId
	taskInfo.Investigatorid = visitId
	taskInfo.Id = id
	taskInfo.Updatetime = time.Now().Format("2006-01-02 15:04:05")
	err4 := models.UpdateTaskInfo(taskInfo)
	if err4 != nil {
		result.ErrCode = -10
		result.ErrMsg = err4.Error()
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

//删除任务
func (this *DelTaskInfo) Post() {
	var result models.ResultInfo
	_, err := checkLogin(this.GetSession("userinfo"))
	if err != nil {
		result.ErrCode = -5 //-5表示未登录状态
		result.ErrMsg = err.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	id, err0 := this.GetInt("orderId")
	if err0 != nil {
		result.ErrCode = -11
		result.ErrMsg = "删除任务获取参数失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	err1 := models.DelTaskInfo(id)
	if err1 != nil {
		result.ErrCode = -12
		result.ErrMsg = err.Error()
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

//登录验证
func checkLogin(ss interface{}) (models.UserInfo, error) {

	var userss models.UserInfo
	if ss == nil {
		return userss, errors.New("用户未登录")
	}
	userss = ss.(models.UserInfo)
	if userss.Id <= 0 {
		return userss, errors.New("登录超时,请重新登录")
	}
	return userss, nil
}
