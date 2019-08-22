package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

// 全局的进行声明下
// 连接实例  就是对象


//由于model这个名字叫 UserInfo 那么操作的表其实 user_info
type UserInfo struct{ 
	Id int64
	Username string
	Password string
}



func AddUser(user_info *UserInfo)(int64,error){
	id,err := db.Insert(user_info)
	return id,err
}

// 被调用者使用的是变量
func ReadUserInfo(users *[]UserInfo){
	qb, _:=orm.NewQueryBuilder("mysql")

	qb.Select("*").From("user_info")


	sql := qb.String()
	db.Raw(sql).QueryRows(users)
}

