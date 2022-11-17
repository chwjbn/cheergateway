package ctl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"cheeradmin/app/db"
	"cheeradmin/app/db/dbmodel"
	"cheeradmin/cheerlib"
	"cheeradmin/config"
	"net/http"
)

type BaseCtl struct {
	mConfig *config.ConfigApp
	mDbPool *db.DbPool
}


func (this *BaseCtl)Init(cfg *config.ConfigApp) error  {


	var xError error=nil

	this.mConfig=cfg

	for {

		xError,this.mDbPool=db.NewDbPool(this.mConfig.DbUri)

		cheerlib.LogInfo(fmt.Sprintf("BaseCtl Init With Config=[%v]",this.mConfig))

		if xError==nil{
			break
		}

		cheerlib.LogError(fmt.Sprintf("db.NewDbPool Error:%s",xError.Error()))
	}

	return xError
}


func (this *BaseCtl)ReturnMsg(ctx *gin.Context,errorCode string,errorMessage string,data interface{})  {

	xRetData := map[string]interface{}{
		"success":true,
		"errorCode":errorCode,
		"errorMessage":errorMessage,
		"data":data,
	}

	ctx.AsciiJSON(http.StatusOK,xRetData)
}



func (this *BaseCtl)ReturnIntenalError(ctx *gin.Context)  {
	this.ReturnMsg(ctx,"500","app.server.msg.system.upgrade",nil)
}

func (this *BaseCtl)ReturnAppError(ctx *gin.Context,errorMessage string)  {
	this.ReturnMsg(ctx,"250",errorMessage,nil)
}

func (this *BaseCtl)ReturnAppSuccess(ctx *gin.Context,errorMessage string)  {
	this.ReturnMsg(ctx,"0",errorMessage,nil)
}

func (this *BaseCtl)ReturnAppSuccessData(ctx *gin.Context,data interface{})  {
	this.ReturnMsg(ctx,"0","app.server.msg.common.request.succ",data)
}

func (this *BaseCtl)ReturnPageData(ctx *gin.Context,pageData dbmodel.PageData)  {

	xRetData := map[string]interface{}{
		"data":pageData.DataList,
		"success":true,
		"total":pageData.TotalCount,
	}

	ctx.AsciiJSON(http.StatusOK,xRetData)
}
