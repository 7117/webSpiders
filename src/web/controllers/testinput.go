package controllers

import (
	"github.com/astaxie/beego"
)

type TestInputController struct {
	beego.Controller
}

func (c *TestInputController) Get() {
	id:=c.GetString("id");
	c.Ctx.WriteString(id);
}

