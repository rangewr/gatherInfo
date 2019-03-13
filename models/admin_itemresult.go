package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"strconv"
)

//检查项目结果表
type Itemresult struct {
	Id          int    `orm:"column(id);pk,auto"`          //编号
	Checkinfoid int    `orm:"column(checkinfoid)"`         //检查单编号id
	Itemid      int    `orm:"column(itemid)"`              //不合格的检查项目编号id
	Imgpath     string `orm:"column(imgpath)"`             //不合格的检查项目图片路径
	Createtime  string `orm:"column(createtime)"`          //创建时间
	Isactive    int    `orm:"column(isactive);default(1)"` //是否可用(1表示可用,2表示不可用)
}

// 自定义表名（系统自动调用）
func (u *Itemresult) TableName() string {
	return "admin_itemresult"
}
func init() {
	orm.RegisterModel(new(Itemresult)) //把LogInfo注册到orm中
}

//插入检查项目结果
func AddItemresultInfo(itemresult Itemresult) (error) {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Insert(&itemresult) //向数据库中插入记录
	if err != nil {
		return errors.New("添加检查结果记录内部错误")
	}
	return err
}

//获取检查结果项目id集合
func GetItemresultIdList(id int) ([]Itemresult, error) {
	var itemresultList []Itemresult
	o := orm.NewOrm()
	o.Using("default")
	var strSql string = "select * from admin_itemresult where isactive = 1 and checkinfoid = " + strconv.Itoa(id) + " group by itemid"
	_, err := o.Raw(strSql).QueryRows(&itemresultList)
	if err != nil {
		return nil, errors.New("获取检查结果项目id异常")
	}
	return itemresultList, nil
}

//获取检查结果集合
func GetItemresultList(id int) ([]Itemresult, error) {
	var itemresultList []Itemresult
	o := orm.NewOrm()
	o.Using("default")
	var strSql string = "select * from admin_itemresult where isactive = 1 and checkinfoid = " + strconv.Itoa(id)
	_, err := o.Raw(strSql).QueryRows(&itemresultList)
	if err != nil {
		return nil, errors.New("获取检查项目结果异常")
	}
	return itemresultList, nil
}
