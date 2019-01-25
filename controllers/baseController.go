package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type (
	baseController struct {
		beego.Controller
		i18n.Locale
	}
)

func (this *baseController) Prepare() {
	lang := this.Ctx.Input.Header("language")
	if lang == "zh" {
		this.Lang = "zh-CN"
	} else {
		this.Lang = "en-US"
	}
}
