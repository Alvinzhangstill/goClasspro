package main

import (
	_ "class/models"
	_ "class/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
