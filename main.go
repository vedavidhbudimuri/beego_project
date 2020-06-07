package main

import (
	"beego_project/models"
	_ "beego_project/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.RegisterModels()
	beego.Run()
}
