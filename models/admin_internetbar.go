package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

//网吧信息表
type Internetbar struct {
	Id             int    `orm:"column(id);pk,auto"`          //编号
	Chargeperson   string `orm:"column(chargeperson)"`        //负责人
	Enterprisename string `orm:"column(enterprisename)"`      //企业名称
	Communityid    int    `orm:"column(communityid)"`         //区县编号
	Address        string `orm:"column(address)"`             //详细地址
	Connect        string `orm:"column(connect)"`             //联系方式
	Enterprisecode int    `orm:"column(enterprisecode)"`      //企业编码
	Isactive       int    `orm:"column(isactive);default(1)"` //是否可用(1表示可用,2表示不可用)
}

// 自定义表名（系统自动调用）
func (u *Internetbar) TableName() string {
	return "admin_internetbar"
}
func init() {
	orm.RegisterModel(new(Internetbar)) //把LogInfo注册到orm中
}

func GetBarList() ([]Internetbar, error) {
	var barList []Internetbar
	o := orm.NewOrm()
	o.Using("default")
	var strSql string = "select * from admin_internetbar where isactive=1 "
	_, err := o.Raw(strSql).QueryRows(&barList)
	if err != nil {
		return nil, errors.New("获取网吧列表失败")
	} else {
		return barList, nil
	}
}
