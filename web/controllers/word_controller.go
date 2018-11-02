// file: web/controllers/hello_controller.go

package controllers

import (
	"github.com/cooladdr/xiuxiubeidanci/services"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris"
)

type WordController struct{
	Ctx iris.Context
	Service services.WordService
}



func (c *WordController) GetBy(word string) mvc.Result {


	w:=c.Service.Find(word)


	return mvc.View{
		Name: "word/index.html",
		Data: map[string]interface{}{
			"W":     w,
		},
	}
}
