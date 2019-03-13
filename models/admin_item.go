package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

//检查实名制表
type Item struct {
	Id         int    `orm:"column(id);pk,auto"`          //编号
	Itemname   string `orm:"column(itemname)"`            //检查项目名称
	Createtime string `orm:"column(createtime)"`          //创建时间
	Isactive   int    `orm:"column(isactive);default(1)"` //是否可用(1表示可用,2表示不可用)
}

// 自定义表名（系统自动调用）
func (u *Item) TableName() string {
	return "admin_item"
}
func init() {
	orm.RegisterModel(new(Item)) //把LogInfo注册到orm中
}

//获取检查项目列表
func GetItmesList() ([]Item, error) {
	var itemList []Item
	o := orm.NewOrm()
	o.Using("default")
	var strSql string = "select * from admin_item where isactive=1 order by id asc"
	_, err := o.Raw(strSql).QueryRows(&itemList)
	if err != nil {
		return nil, errors.New("获取检查项目列表失败")
	} else {
		return itemList, nil
	}
}
