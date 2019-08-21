package controllers

import (
	"github.com/astaxie/beego"
)

type TestInputController struct {
	beego.Controller
}

type User struct{
	Username string
	Password string
}

func (c *TestInputController) Get() {
	c.Ctx.WriteString(`<html><form action="http://127.0.0.1:8080/testinput" method="post"> 
	<input type="text" name="Username"/>
	<input type="password" name="Password"/>
	<input type="submit" value="提交"/>
	</form></html>`)
}

func (c *TestInputController) Post(){
	u := User{}
	if err:=c.ParseForm(&u) ; err != nil{

	}
	c.Ctx.WriteString("Username:" + u.Username + " Password:" + u.Password)
}

