package routers

import (
	"gatherInfo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//===================登录退出=======================
	beego.Router("/proc/Login", &controllers.UserLogin{})   //登录方法 Post
	beego.Router("/proc/Logout", &controllers.UserLogout{}) //登出方法 Get
	//====================后台任务管理===================
	beego.Router("/admin/GetTaskList", &controllers.GetTaskList{})       //获取列表 Get
	beego.Router("/admin/AddTaskInfo", &controllers.AddTaskInfo{})       //添加一个任务 Post
	beego.Router("/admin/DelTaskInfo", &controllers.DelTaskInfo{})       //删除一个任务的详细信息 Post
	beego.Router("/admin/UpdateTaskInfo", &controllers.UpdateTaskInfo{}) //更新一个任务的详细信息 Post
	//===================手持端任务管理==================
	beego.Router("/admin/GetBarList", &controllers.GetBarList{})       //获取网吧列表 Get
	beego.Router("/admin/GetRegionList", &controllers.GetRegionList{}) //获取区县列表 Get
	beego.Router("/admin/GetVisitList", &controllers.GetVisitList{})   //获取调查员列表 Get
	beego.Router("/admin/GetItemsList", &controllers.GetItemsList{})   //获取检查项目列表 Get
	//===================实名制检查记录==================
	beego.Router("/admin/GetRealnameList", &controllers.GetRealnameList{}) //获取实名制列表 Get
	beego.Router("/admin/AddRealnameInfo", &controllers.AddRealnameInfo{}) //添加实名制信息 Post
	beego.Router("/admin/DelRealnameInfo", &controllers.DelRealnameInfo{}) //删除实名制信息记录 Post
	//===================检查项目结果记录================
	beego.Router("/admin/AddItemresultInfo", &controllers.AddItemresultInfo{}) //保存检查项目结果记录 Post
	beego.Router("/admin/GetItemresultList", &controllers.GetItemresultList{}) //获取检查项目结果记录 Get

}
