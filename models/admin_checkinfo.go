package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"strconv"
)

//检查任务表
type Checkinfo struct {
	Id             int    `orm:"column(id);pk,auto"`          //编号
	Communityid    int    `orm:"column(communityid)"`         //区县编号
	Internetbarid  int    `orm:"column(internetbarid)"`       //网吧编号
	Investigatorid int    `orm:"column(investigatorid)"`      //调查员编号
	Updatetime     string `orm:"column(updatetime)"`          //修改时间
	Createtime     string `orm:"column(createtime)"`          //创建时间
	Orderstate     int    `orm:"column(orderstate)"`          //任务状态(1表示未办结,2表示已完结)
	Isactive       int    `orm:"column(isactive);default(1)"` //是否可用(1表示可用,2表示不可用)
}

// 自定义表名（系统自动调用）
func (u *Checkinfo) TableName() string {
	return "admin_checkinfo"
}
func init() {
	orm.RegisterModel(new(Checkinfo)) //把LogInfo注册到orm中
}

//获取任务集合
func GetTaskList(strWhere string) ([]Checkinfo, error) {
	var taskList []Checkinfo
	o := orm.NewOrm()
	o.Using("default")
	var strSql string = "select * from admin_checkinfo where isactive=1 "
	if strWhere != "" {
		strSql += " and " + strWhere
	}
	strSql += " order by createtime desc"
	_, err := o.Raw(strSql).QueryRows(&taskList)
	//err:=o.Read(&user) //这种方法只适合主键查询
	if err != nil {
		return nil, errors.New("获取任务列表失败")
	} else {
		return taskList, nil
	}
}

//添加新的任务
func AddTaskInfo(checkinfo Checkinfo) error {
	if checkinfo.Internetbarid == 0 {
		return errors.New("网吧编码为0")
	}
	o := orm.NewOrm()
	o.Using("default")
	var checkinfos []Checkinfo
	var strSql = "select * from admin_checkinfo where isactive=1 and internetbarid='" + strconv.Itoa(checkinfo.Internetbarid) + "'"
	num, err := o.Raw(strSql).QueryRows(&checkinfos)
	if err != nil {
		return err
	}
	//判断用户名是否存在
	if num > 0 {
		return errors.New("此任务已存在")
	}
	_, err = o.Insert(&checkinfo) //向数据库中插入记录
	if err != nil {
		return errors.New("新增任务内部错误")
	}
	return err
}

//更新任务信息
func UpdateTaskInfo(task Checkinfo) (error) {
	o := orm.NewOrm()
	o.Using("default")
	var check Checkinfo
	var strSql1 = "select * from admin_checkinfo where isactive=1 and id='" + strconv.Itoa(task.Id) + "'"
	err0 := o.Raw(strSql1).QueryRow(&check)
	if err0 != nil {
		return err0
	}
	task.Createtime = check.Createtime
	task.Orderstate = check.Orderstate
	task.Isactive = 1
	_, err1 := o.Update(&task)
	if err1 != nil {
		return errors.New("更新失败")
	}
	return nil
}

//删除指定任务
func DelTaskInfo(taskId int) (error) {
	o := orm.NewOrm()
	o.Using("default")
	var strSql1 = "update admin_checkinfo set isactive=2 where id='" + strconv.Itoa(taskId) + "'"
	_, err := o.Raw(strSql1).Exec()
	if err != nil {
		return errors.New("删除任务错误")
	} else {
		return nil
	}
}

//修改任务状态
func UpdateTaskState(id int) (error) {
	o := orm.NewOrm()
	o.Using("default")
	var strSql1 = "update admin_checkinfo set orderstate = 2 where id=" + strconv.Itoa(id)
	_, err := o.Raw(strSql1).Exec()
	if err != nil {
		return errors.New("修改任务转态错误")
	} else {
		return nil
	}
}
