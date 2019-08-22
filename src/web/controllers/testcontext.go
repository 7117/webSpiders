package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
)

type TestContextController struct {
	beego.Controller
}

func (c *TestContextController) Get() {
	// int转为string
	c.Ctx.WriteString(c.Ctx.Input.IP() + ":" + strconv.Itoa(c.Ctx.Input.Port()))
	// $_REQUEST["name"]
	c.Ctx.WriteString(c.Ctx.Input.Query("name"))
	// output
	m := make(map[string]float64)
	m["a"] = 99
	c.Ctx.Output.JSON(m, false, false)
}
