package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"strconv"
)

//检查实名制表
type Realname struct {
	Id               int    `orm:"column(id);pk,auto"`          //编号
	Checkinfoid      int    `orm:"column(checkinfoid)"`         //检查单编号id
	Interviewername  string `orm:"column(interviewername)"`     //被检查人姓名
	Interviewerbirth string `orm:"column(interviewerbirth)"`    //被检查人出生日期
	Reason           string `orm:"column(reason)"`              //原因及备注
	Createtime       string `orm:"column(createtime)"`          //创建时间
	Isactive         int    `orm:"column(isactive);default(1)"` //是否可用(1表示可用,2表示不可用)
}

// 自定义表名（系统自动调用）
func (u *Realname) TableName() string {
	return "admin_realname"
}
func init() {
	orm.RegisterModel(new(Realname)) //把LogInfo注册到orm中
}

//添加实名制检查记录
func AddRealnameInfo(realInfo Realname) error {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Insert(&realInfo) //向数据库中插入记录
	if err != nil {
		return errors.New("新增实名制记录内部错误")
	}
	return err
}

//查询实名制检查记录
func GetRealnameList(strWhere string) ([]Realname, error) {
	var realList []Realname
	o := orm.NewOrm()
	o.Using("default")
	var strSql string = "select * from admin_realname where isactive=1 "
	if strWhere != "" {
		strSql += " and " + strWhere
	}
	strSql += " order by createtime asc"
	_, err := o.Raw(strSql).QueryRows(&realList)
	if err != nil {
		return nil, errors.New("获取实名制记录列表失败")
	} else {
		return realList, nil
	}
}

//删除一个实名制检查记录
func DelRealnameInfo(id int) (error) {
	o:= orm.NewOrm()
	o.Using("default")
	var strSql = "update admin_realname set isactive = 2 where id = "+strconv.Itoa(id)
	_, err := o.Raw(strSql).Exec()
	if err != nil {
		return errors.New("删除实名制记录内部错误")
	} else {
		return nil
	}
}
