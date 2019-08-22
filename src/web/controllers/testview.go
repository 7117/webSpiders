package controllers

import (
	"github.com/astaxie/beego"
	"web/models"
)

type TestViewController struct {
	beego.Controller
}

func (c *TestViewController) Get() {
	// 构建数组
	var users []models.UserInfo
	// 调用方法  调用者给的是地址  被调用者给的是变量
	models.ReadUserInfo(&users)

	// data分配数据
	c.Data["Title"] = "aaa"
	c.Data["IsDisplay"] = 1;
	c.Data["Content"] = "dd";
	c.Data["Users"] = users
	c.Data["len"] = len(users)

	// 指定模板
	c.TplName = "test_view.tpl"
}
