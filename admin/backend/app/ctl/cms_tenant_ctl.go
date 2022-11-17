package ctl

import (
	"cheeradmin/app/ctl/ctldata"
	"cheeradmin/app/db/dbmodel"
	"cheeradmin/cheerlib"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"strings"
)

// 获取图形验证码
func (this *CmsCtl)CtlCheckCodeImage(ctx *gin.Context)  {

	xDriver:=base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	xCaptcha:=base64Captcha.NewCaptcha(xDriver,base64Captcha.DefaultMemStore)

	xCheckCodeImageId,xCheckCodeImageData,xCheckCodeImageErr:=xCaptcha.Generate()

	if xCheckCodeImageErr!=nil{
		this.ReturnIntenalError(ctx)
		return
	}

	xImageData:=make(map[string]interface{})
	xImageData["code_image_id"]=xCheckCodeImageId
	xImageData["code_image_data"]=xCheckCodeImageData

	this.ReturnMsg(ctx,"0","succ.",xImageData)
}

func (this *CmsCtl)CtlTenantInit(ctx *gin.Context)  {

	xReqIp:=ctx.ClientIP()

	if !strings.EqualFold(xReqIp,"127.0.0.1"){
		this.ReturnAppError(ctx,"app.server.msg.tenant.init.disallow")
		return
	}

	xTenantInfo:=dbmodel.TenantInfo{}

	xPwd:="Cheergo112233"

	xTenantInfo.UserName="gadmin"
	xTenantInfo.Email="publics@foxmail.com"
	xTenantInfo.Password=xPwd
	xTenantInfo.CreateIp=xReqIp
	xTenantInfo.UserImgUrl=fmt.Sprintf("https://github.com/identicons/%s.png",cheerlib.EncryptMd5(xTenantInfo.UserName))
	xTenantInfo.Init()

	xTenantInfo.Status="active"

	xTenantInfoData:=this.mDbPool.DbTenantSvc.GetTenantInfoByUserName(xTenantInfo.UserName)
	if len(xTenantInfoData.UserName)>1{
		this.ReturnAppError(ctx,"app.server.msg.tenant.init.exists.error")
		return
	}

	xErr:=this.mDbPool.DbTenantSvc.CreateTenantInfo(xTenantInfo)

	if xErr!=nil{
		this.ReturnAppError(ctx,xErr.Error())
		return
	}

	this.ReturnMsg(ctx,"0","app.server.msg.tenant.init.succ",fmt.Sprintf("UserName:[%s],Password:[%s]",xTenantInfo.UserName,xPwd))
}


//登录
func (this *CmsCtl)CtlTenantLogin(ctx *gin.Context)  {

	xRequestData:=ctldata.TenantLoginRequest{}
	xError:=ctx.BindJSON(&xRequestData)
	if xError!=nil {
		return
	}

	xError=xRequestData.CheckInput()

	if xError!=nil{
		this.ReturnMsg(ctx,"250",xError.Error(),nil)
		return
	}

	if !base64Captcha.DefaultMemStore.Verify(xRequestData.CheckCodeImgId,xRequestData.CheckCodeImgData,true){
		this.ReturnMsg(ctx,"250","app.server.msg.tenant.checkimgcode.error",nil)
		return
	}

	xTenantInfo:=this.mDbPool.DbTenantSvc.GetTenantInfoByUserName(xRequestData.UserName)
	if len(xTenantInfo.TenantId)<1{
		this.ReturnMsg(ctx,"250","app.server.msg.tenant.account.error",nil)
		return
	}

	xCheckPassword:=xTenantInfo.EncryptPassword(xRequestData.Password,xTenantInfo.PwdSalt)
	if !strings.EqualFold(xCheckPassword,xTenantInfo.Password){
		this.ReturnMsg(ctx,"250","app.server.msg.tenant.account.error",nil)
		return
	}

	if !strings.EqualFold(xTenantInfo.Status,"active"){
		this.ReturnMsg(ctx,"250","app.server.msg.tenant.account.status.noactive",nil)
		return
	}

	xTenantAccessSession:=dbmodel.TenantAccessSession{}
	xTenantAccessSession.TenantId=xTenantInfo.TenantId
	xTenantAccessSession.LoginAddr=this.getClientIp(ctx)
	xTenantAccessSession.LastAddr=this.getClientIp(ctx)

	xTenantAccessSession.Init()

	xError=this.mDbPool.DbTenantSvc.CreateAccessSession(xTenantAccessSession)
	if xError!=nil{
		this.ReturnMsg(ctx,"500","app.server.msg.common.server.error",nil)
		cheerlib.LogError(fmt.Sprintf("DbTenantSvc.CreateAccessSession Error:%s",xError.Error()))
		return
	}

	xTenantLoginRespData:=ctldata.TenantLoginRespData{}
	xTenantLoginRespData.TokenId=xTenantAccessSession.TokenId
	xTenantLoginRespData.TenantId=xTenantAccessSession.TenantId
	xTenantLoginRespData.UserName=xTenantInfo.UserName
	xTenantLoginRespData.UserImgUrl=xTenantInfo.UserImgUrl
	xTenantLoginRespData.Role="admin"

	this.ReturnMsg(ctx,"0","app.server.msg.tenant.login.succ",xTenantLoginRespData)
}

func (this *CmsCtl)CtlTenantCurrent(ctx *gin.Context)  {

	xIsLogin,xSessData:=this.checkLogin(ctx)
	if !xIsLogin{
		return
	}

	xData:=this.mDbPool.DbTenantSvc.GetTenantInfo(xSessData.TenantId)

	this.ReturnMsg(ctx,"0","app.server.msg.common.op.succ",xData)
}


func (this *CmsCtl)CtlTenantUpdateInfo(ctx *gin.Context)  {

	xIsLogin,_:=this.checkLogin(ctx)
	if !xIsLogin{
		return
	}

	xReqData:=ctldata.TenantInfoRequest{}
	xError:=ctx.BindJSON(&xReqData)
	if xError!=nil{
		return
	}

	xError=xReqData.CheckInputForUpdateInfo()
	if xError!=nil{
		this.ReturnAppError(ctx,xError.Error())
		return
	}

	xData:=this.mDbPool.DbTenantSvc.GetTenantInfo(xReqData.Data.TenantId)
	if len(xData.Status)<1{
		this.ReturnAppError(ctx,"此账号信息不存在,操作失败!")
		return
	}

	xData.Email=xReqData.Data.Email
	this.mDbPool.DbTenantSvc.SaveTenantInfo(xData)

	this.ReturnMsg(ctx,"0","app.server.msg.common.op.succ",xData)
}

func (this *CmsCtl)CtlTenantUpdatePassword(ctx *gin.Context)  {

	xIsLogin,_:=this.checkLogin(ctx)
	if !xIsLogin{
		return
	}

	xReqData:=ctldata.TenantInfoRequest{}
	xError:=ctx.BindJSON(&xReqData)
	if xError!=nil{
		return
	}

	xError=xReqData.CheckInputForUpdatePassword()
	if xError!=nil{
		this.ReturnAppError(ctx,xError.Error())
		return
	}

	xData:=this.mDbPool.DbTenantSvc.GetTenantInfo(xReqData.Data.TenantId)
	if len(xData.Status)<1{
		this.ReturnAppError(ctx,"此账号信息不存在,操作失败!")
		return
	}

	if strings.EqualFold(xData.Password,xReqData.Data.Password){
		this.ReturnMsg(ctx,"0","app.server.msg.common.op.succ",nil)
		return
	}

	xData.Password=xData.EncryptPassword(xReqData.Data.Password,xData.PwdSalt)
	this.mDbPool.DbTenantSvc.SaveTenantInfo(xData)

	this.ReturnMsg(ctx,"0","app.server.msg.common.op.succ",nil)
}
