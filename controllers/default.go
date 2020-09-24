package controllers

import (
	"Beego9.24/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Post(){
	var person models.Person
	dataBytes, err := ioutil.ReadAll(c.Ctx.Request.Body)
	json.Unmarshal(dataBytes,&person)
	if err != nil{
		c.Ctx.WriteString("数据接收错误，请重试。")
	}
	err = json.Unmarshal(dataBytes,&person)
	if err != nil{
		c.Ctx.WriteString("数据解析失败，请重试。")
	}
	fmt.Println("姓名：",person.Name)
	fmt.Println("生日：",person.Birthday)
	fmt.Println("地址：",person.Address)
	fmt.Println("小名：",person.Nick)
	c.Ctx.WriteString("数据解析成功。")
}
