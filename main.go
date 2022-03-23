package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/test/models"
	_ "github.com/test/routers"
)

func main() {
	fmt.Println("hello scorecard...")
	models.LinkedDome()
	beego.Run()
}
