package controllers

import (
	"liteblog/models"
	"liteblog/syserror"
	"errors"
)

type PraiseController struct {
	BaseController
}

func (this *PraiseController) NextPrepare() {
	this.MustLogin()
}

// @router /:type/:key [post]
func (this *PraiseController) Praise() {
	key := this.Ctx.Input.Param(":key")
	ttype := this.Ctx.Input.Param(":type")
	table := "notes"
	switch ttype {
		case "message":
			table = "messages"
		case "note":
			table = "notes"
		default:
			this.Abort500(errors.New("未知类型"))
	}
	pcnt,err := models.UpdatePraise(table,key,int(this.User.ID))
	if err != nil {
		if e2,ok := err.(syserror.HasPraiseError);ok{
			this.Abort500(e2)
		}
		this.Abort500(syserror.New("点赞失败", err))
	}
	this.JSONOkH("点赞成功",H{"praise":pcnt})
}
