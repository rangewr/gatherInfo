package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

//区县信息表
type Region struct {
	Id         int    `orm:"column(id);pk,auto"`          //编号
	Regionname string `orm:"column(regionname)"`          //地区名称
	Isactive   int    `orm:"column(isactive);default(1)"` //是否可用(1表示可用,2表示不可用)
}

// 自定义表名（系统自动调用）
func (u *Region) TableName() string {
	return "admin_region"
}
func init() {
	orm.RegisterModel(new(Region)) //把LogInfo注册到orm中
}

func GetRegionList() ([]Region, error) {
	var regionList []Region
	o := orm.NewOrm()
	o.Using("default")
	var strSql string = "select * from admin_region where isactive=1 "
	_, err := o.Raw(strSql).QueryRows(&regionList)
	if err != nil {
		return nil, errors.New("获取区县列表失败")
	} else {
		return regionList, nil
	}
}
