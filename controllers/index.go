package controllers

import (
	"liteblog/models"
	"liteblog/syserror"
)

type IndexController struct {
	BaseController
}

// @router / [get]
func (this *IndexController) Get() {
	limit := 10
	// page
	page, err := this.GetInt("page", 1)
	if err != nil || page <= 0 {
		page = 1
	}

	title := this.GetString("title")
	//得到页面的数据
	notes, err := models.QueryNotesByPage(title,page, limit)
	if err != nil {
		this.Abort500(err)
	}
	this.Data["notes"] = notes

	//这页数
	count, err := models.QueryNotesCount(title)
	if err != nil {
		this.Abort500(err)
	}
	totpage := count / limit
	if count%limit != 0 {
		totpage = totpage + 1
	}
	this.Data["totpage"] = totpage
	this.Data["page"] = page
	this.Data["title"] = title
	this.TplName = "index.html"
}

//@router /message [get]
func (this *IndexController) GetMessage() {
	this.TplName = "message.html"
}

// @router /about [get]
func (this *IndexController) GetAbout() {
	this.TplName = "about.html"
}

// @router /user [get]
func (this *IndexController) GetUser() {
	this.TplName = "user.html"
}

// @router /reg [get]
func (this *IndexController) GetReg() {
	this.TplName = "reg.html"
}

//@router /comment/:key
func (this *IndexController) GetComment(){
	key:= this.Ctx.Input.Param(":key")
	note,err:=models.QueryNoteByKey(key)
	if err!=nil{
		this.Abort500(syserror.New("文章不存在",err))
	}
	this.Data["note"] = note
	this.TplName="comment.html"
}

//@router /details/:key [get]
func (this *IndexController) GetDetails() {
	key:= this.Ctx.Input.Param(":key")
	note,err:=models.QueryNoteByKey(key)
	if err!=nil{
		this.Abort500(syserror.New("文章不存在",err))
	}
	this.Data["note"] = note

	ms,err:=models.QueryMessagesByNoteKey(key)
	if err!=nil{
		this.Abort500(syserror.New("文章不存在",err))
	}
	this.Data["messages"] = ms
	this.TplName="details.html"
}