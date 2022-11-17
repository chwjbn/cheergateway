package ctl

import (
	"cheeradmin/cheerlib"
	"github.com/gin-gonic/gin"
	"cheeradmin/app/db/dbmodel"
)

type CmsCtl struct {
	BaseCtl
}

func (this *CmsCtl)getClientIp(ctx *gin.Context) string  {

	xData:=ctx.GetHeader("X-Cheer-Client-IP")
	if len(xData)>0{
		return xData
	}

	xData=ctx.GetHeader("x-cheer-client-ip")
	if len(xData)>0{
		return xData
	}

	xData=ctx.GetHeader("X-Real-IP")
	if len(xData)>0{
		return xData
	}

	xData=ctx.GetHeader("x-real-ip")
	if len(xData)>0{
		return xData
	}

	xData=ctx.ClientIP()

	return xData

}

func (this *CmsCtl)checkLogin(ctx *gin.Context)(bool, dbmodel.TenantAccessSession)  {

	xIsLogin:=false
	xSession:= dbmodel.TenantAccessSession{}


	xTokenId:=ctx.Request.Header.Get("Authorization")

	if len(xTokenId)<1{
		xTokenId=ctx.DefaultQuery("token_id","")
	}

	if len(xTokenId)>10{

		xSession=this.mDbPool.DbTenantSvc.GetAccessSession(xTokenId)
		if len(xSession.TenantId)>10{
			xIsLogin=true

			xSession.LastAddr=this.getClientIp(ctx)
			xSession.LastTime=cheerlib.TimeGetNow()
			this.mDbPool.DbTenantSvc.UpdateAccessSession(xSession)
		}

	}

	if !xIsLogin{
		this.ReturnMsg(ctx,"401","app.server.msg.tenant.account.notlogin",nil)
	}

	return xIsLogin,xSession

}
