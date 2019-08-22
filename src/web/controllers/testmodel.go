package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type TestModelController struct {
	beego.Controller
}

//由于model这个名字叫 UserInfo 那么操作的表其实 user_info
type UserInfo struct{ 
	Id int64
	Username string
	Password string
}

func (c *TestModelController) Get() {
	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	orm.RegisterModel(new(UserInfo))

	o := orm.NewOrm()
	// //下面是插入
	// user := UserInfo{Username:"aaaa", Password:"123456"}
	// id,_ := o.Insert(&user)
	// c.Ctx.WriteString(fmt.Sprintf("id:%v\n", id))

	// //下面是更新
	// user2 := UserInfo{Id:2,Username:"ttttt", Password:"1"}
	// id,_:=o.Update(&user2)
	// c.Ctx.WriteString(fmt.Sprintf("id:%v\n", id))

	// //下面是读取
	// user3 := UserInfo{Id:2}
	// o.Read(&user3)
	// c.Ctx.WriteString(fmt.Sprintf("user info:%v\n",user3.Id))

	// // 下面是删除
	// u:=UserInfo{Id:4}
	// o.Delete(&u);

	// 原生SQL
	// var maps []orm.Params
	// o.Raw("select * from user_info where Id=1").Values(&maps);

	// for _,v:=range maps{
	// 	c.Ctx.WriteString(fmt.Sprintf("userinfo:%v\n",v));
	// }

	// 原生SQL
	// var usrs []UserInfo
	// o.Raw("select * from user_info").QueryRows(&usrs);
	// c.Ctx.WriteString(fmt.Sprintf("user:%v", usrs))

	// querybuild
	var users []UserInfo
	qb, _:=orm.NewQueryBuilder("mysql")
	qb.Select("password").From("user_info").Where("username= ?").And("id=1").Limit(1)
	sql := qb.String()  	//获取sql
	o.Raw(sql,"1").QueryRows(&users)	//绑定参数
	c.Ctx.WriteString(fmt.Sprintf("user info:%v", users))

}