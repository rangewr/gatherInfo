package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

//用户表
type UserInfo struct {
	Id            int    `orm:"column(id);pk,auto"`          //编号
	Loginname     string `orm:"column(loginname)"`           //登录名
	Loginpwd      string `orm:"column(loginpwd)"`            //登录口令:存储的是MD5后的密码
	Mobileno      string `orm:"column(mobileno)"`            //手机号码
	Communityid   int    `orm:"column(communityid)"`         //所属社区编号
	Isactive      int    `orm:"column(isactive);default(1)"` //可用标志:1:默认值,可用;2:代表已删除
	Registertime  string `orm:"column(registertime)"`        //注册时间
	Lastlogintime string `orm:"column(lastlogintime);null"`  //最新登陆时间
	Usertype      int    `orm:"column(usertype);default(1)"` //用户类型:1:管理员;2:普通用户
	Realname      string `orm:"column(realname)"`            //真实姓名
}

// 自定义表名（系统自动调用）
func (u *UserInfo) TableName() string {
	return "sys_userinfo"
}
func init() {
	orm.RegisterModel(new(UserInfo)) //把LogInfo注册到orm中
}

//登录验证
func Login(strLoginName string, strLoginPwd string) (UserInfo, error) {
	user := UserInfo{Loginname: strLoginName, Loginpwd: strLoginPwd, Isactive: 1}
	if strLoginName == "" || strLoginPwd == "" {
		return user, errors.New("用户名和密码不允许为空")
	}
	o := orm.NewOrm()
	o.Using("default")
	var strSql string = "select * from sys_userinfo where isactive=1 and loginname='" + strLoginName + "' and loginpwd='" + strLoginPwd + "' "
	err := o.Raw(strSql).QueryRow(&user)
	//err:=o.Read(&user) //这种方法只适合主键查询
	if err != nil {
		return user, errors.New("用户名或口令错误")
	} else {
		return user, nil
	}
}

//获取调查员列表
func GetVisiterList() ([]UserInfo, error) {
	var visitList []UserInfo
	o := orm.NewOrm()
	o.Using("default")
	var strSql string = "select id,realname,communityid from sys_userinfo where isactive=1 and usertype = 2"
	_, err := o.Raw(strSql).QueryRows(&visitList)
	if err != nil {
		return nil, errors.New("获取调查员列表失败")
	} else {
		return visitList, nil
	}
}
