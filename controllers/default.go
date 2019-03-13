package controllers

import (
	"gatherInfo/models"
	"github.com/astaxie/beego"
	"github.com/typa01/go-utils"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type MainController struct {
	beego.Controller
}

type AddItemresultInfo struct {
	beego.Controller
}

type GetItemresultList struct {
	beego.Controller
}

type GetBarList struct {
	beego.Controller
}

type GetRegionList struct {
	beego.Controller
}

type GetVisitList struct {
	beego.Controller
}

type GetItemsList struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

//获取网吧列表
func (this *GetBarList) Get() {
	var result models.ResultInfo
	_, err := checkLogin(this.GetSession("userinfo"))
	if err != nil {
		result.ErrCode = -5 //-5表示未登录状态
		result.ErrMsg = err.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	barList, err1 := models.GetBarList()
	if err1 != nil {
		result.ErrCode = -3
		result.ErrMsg = "获取网吧列表失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	result.ErrCode = 0
	result.ErrMsg = ""
	result.Result = barList
	this.Data["json"] = result
	this.ServeJSON()
	return
}

//获取区县信息
func (this *GetRegionList) Get() {
	var result models.ResultInfo
	_, err := checkLogin(this.GetSession("userinfo"))
	if err != nil {
		result.ErrCode = -5 //-5表示未登录状态
		result.ErrMsg = err.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	regionList, err1 := models.GetRegionList()
	if err1 != nil {
		result.ErrCode = -4
		result.ErrMsg = "获取地区列表失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	result.ErrCode = 0
	result.ErrMsg = ""
	result.Result = regionList
	this.Data["json"] = result
	this.ServeJSON()
	return
}

//上传文件(保存检查结果)
func (this *AddItemresultInfo) Post() {
	var result models.ResultInfo
	_, err := checkLogin(this.GetSession("userinfo"))
	if err != nil {
		result.ErrCode = -5 //-5表示未登录状态
		result.ErrMsg = err.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	checkinfoid, err0 := this.GetInt("checkinfoid")
	fileIndex := this.GetString("fileIndex")
	fileSlice := strings.Split(fileIndex, ",")
	if err0 != nil {
		result.ErrCode = -22
		result.ErrMsg = "保存检查项目结果获取任务id异常"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	for _, val := range fileSlice {
		list, err := this.GetFiles("file" + val)
		if err != nil {
			result.ErrCode = -13
			result.ErrMsg = "获取文件异常"
			this.Data["json"] = result
			this.ServeJSON()
			return
		}
		for _, v := range list {
			s, _ := v.Open()
			defer s.Close()
			fileName := tsgutils.GUID() + path.Ext(v.Filename)
			n, _ := os.Create("./static/upload/" + fileName)
			defer n.Close()
			io.Copy(n, s)
			//设置插入对象赋值,组装
			var itemresult models.Itemresult
			itemresult.Checkinfoid = checkinfoid //任务id
			itemid, err2 := strconv.Atoi(val)
			if err2 != nil {
				result.ErrCode = -23
				result.ErrMsg = "保存检查结果时转换项目id异常"
				this.Data["json"] = result
				this.ServeJSON()
				return
			}
			itemresult.Itemid = itemid                        //检查项目id
			itemresult.Imgpath = "/static/upload/" + fileName //图片路径
			itemresult.Isactive = 1
			itemresult.Createtime = time.Now().Format("2006-01-02 15:04:05")
			//插入数据库
			err3 := models.AddItemresultInfo(itemresult)
			if err3 != nil {
				result.ErrCode = -24
				result.ErrMsg = err3.Error()
				this.Data["json"] = result
				this.ServeJSON()
				return
			}
		}
	}
	err4 := models.UpdateTaskState(checkinfoid)
	if err4 != nil {
		result.ErrCode = -25
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

//获取调查员信息
func (this *GetVisitList) Get() {
	var result models.ResultInfo
	_, err := checkLogin(this.GetSession("userinfo"))
	if err != nil {
		result.ErrCode = -5 //-5表示未登录状态
		result.ErrMsg = err.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	visitList, err1 := models.GetVisiterList()
	if err1 != nil {
		result.ErrCode = -14
		result.ErrMsg = "获取调查员列表失败"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	result.ErrCode = 0
	result.ErrMsg = ""
	result.Result = visitList
	this.Data["json"] = result
	this.ServeJSON()
	return
}

//获取检查项目列表
func (this *GetItemsList) Get() {
	var result models.ResultInfo
	_, err := checkLogin(this.GetSession("userinfo"))
	if err != nil {
		result.ErrCode = -5 //-5表示未登录状态
		result.ErrMsg = err.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	itemList, err0 := models.GetItmesList()
	if err0 != nil {
		result.ErrCode = -21
		result.ErrMsg = err0.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	result.ErrCode = 0
	result.ErrMsg = ""
	result.Result = itemList
	this.Data["json"] = result
	this.ServeJSON()
	return
}

//获取检查项目结果列表
func (this *GetItemresultList) Get() {
	var result models.ResultInfo
	_, err := checkLogin(this.GetSession("userinfo"))
	if err != nil {
		result.ErrCode = -5 //-5表示未登录状态
		result.ErrMsg = err.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	checkinfoid, err0 := this.GetInt("checkid")
	if err0 != nil {
		result.ErrCode = -26
		result.ErrMsg = "获取检查项目结果获取参数错误"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	resultMap := make(map[string]interface{})
	itemIdList, err1 := models.GetItemresultIdList(checkinfoid)
	if err1 != nil {
		result.ErrCode = -27
		result.ErrMsg = err1.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	resultMap["itemIdList"] = itemIdList
	itemList, err2 := models.GetItemresultList(checkinfoid)
	if err2 != nil {
		result.ErrCode = -28
		result.ErrMsg = err1.Error()
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	resultMap["itemList"] = itemList
	result.ErrCode = 0
	result.ErrMsg = ""
	result.Result = resultMap
	this.Data["json"] = result
	this.ServeJSON()
	return

}
