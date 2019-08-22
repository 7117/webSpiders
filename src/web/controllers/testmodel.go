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
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	orm.RegisterModel(new(UserInfo))

	o := orm.NewOrm()
	//下面是插入
	user := UserInfo{Username:"aaaa", Password:"123456"}
	id,_ := o.Insert(&user)
	c.Ctx.WriteString(fmt.Sprintf("id:%v\n", id))

	//下面是更新
	user2 := UserInfo{Username:"112", Password:"1"}
	user2.Id = 2
	o.Update(&user2)
	
	//下面是读取
	user3 := UserInfo{Id:2}
	o.Read(&user3)
	c.Ctx.WriteString(fmt.Sprintf("user info:%v\n",user3.Id))

	u:=UserInfo{Id:4}
	o.Delete(&u);


}