package ctl

import (
	"cheeradmin/app/ctl/ctldata"
	"cheeradmin/app/db/dbmodel/app_data"
	"cheeradmin/cheerlib"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (this *CmsCtl)CtlAppGatewayRulePage(ctx *gin.Context)  {

	xIsLogin,_:=this.checkLogin(ctx)
	if !xIsLogin{
		return
	}

	xReqData:=ctldata.AppGatewayRulePageRequest{}
	xError:=ctx.BindJSON(&xReqData)
	if xError!=nil{
		return
	}

	xWhere:=make(map[string]interface{})
	xSort:=make(map[string]interface{})

	if len(xReqData.ConfigDataId)>0{
		xWhere["config_data_id"]=xReqData.ConfigDataId
	}

	if len(xReqData.RuleName)>0{
		xWhere["rule_name"]=bson.M{"$regex":fmt.Sprintf("%s",xReqData.RuleName),"$options": "im"}
	}

	if len(xReqData.SiteDataId)>0{
		xWhere["site_data_id"]=xReqData.SiteDataId
	}

	if len(xReqData.Status)>0{
		xWhere["status"]=xReqData.Status
	}

	xSort["rule_order_no"]=1

	xPageData:=this.mDbPool.DbAppSvc.GetAppGatewayRulePageData(xReqData.PageNo,xReqData.PageSize,xWhere,xSort)

	this.ReturnPageData(ctx,xPageData)
}


func (this *CmsCtl)CtlAppGatewayRuleAdd(ctx *gin.Context)  {

	xIsLogin,xSessionInfo:=this.checkLogin(ctx)
	if !xIsLogin{
		return
	}

	xReqData:=ctldata.AppGatewayRuleAddRequest{}
	xError:=ctx.BindJSON(&xReqData)
	if xError!=nil{
		return
	}

	xReqData.Data.ActionType="backend"

	xError=xReqData.Check()
	if xError!=nil{
		this.ReturnAppError(ctx,xError.Error())
		return
	}

	xData:=app_data.AppGatewayRule{}
	xData.ConfigDataId=xReqData.Data.ConfigDataId
	xData.SiteDataId=xReqData.Data.SiteDataId
	xData.RuleOrderNo=xReqData.Data.RuleOrderNo
	xData.RuleName=xReqData.Data.RuleName

	xData.MatchTarget=xReqData.Data.MatchTarget
	xData.MatchOp=xReqData.Data.MatchOp
	xData.MatchData=xReqData.Data.MatchData

	xData.ActionType=xReqData.Data.ActionType
	xData.ActionData=xReqData.Data.ActionData

	xError=xData.Init(xSessionInfo.TenantId)
	if xError!=nil{
		this.ReturnAppError(ctx,xError.Error())
		return
	}

	xError=this.mDbPool.DbAppSvc.AddAppData(&xData)
	if xError!=nil{
		this.ReturnAppError(ctx,xError.Error())
		return
	}

	this.mDbPool.DbAppSvc.RenewAppGatewayConfigLastVer(xData.ConfigDataId)

	this.ReturnAppSuccess(ctx,"app.server.msg.common.op.succ")
}



func (this *CmsCtl)CtlAppGatewayRuleGet(ctx *gin.Context)  {

	xIsLogin,_:=this.checkLogin(ctx)
	if !xIsLogin{
		return
	}

	xReqData:=ctldata.BaseGetInfoRequest{}
	xError:=ctx.BindJSON(&xReqData)
	if xError!=nil{
		return
	}



	xError=xReqData.Check()
	if xError!=nil{
		this.ReturnAppError(ctx,xError.Error())
		return
	}

	xData:=app_data.AppGatewayRule{}
	xData.DataId=xReqData.DataId

	xError=this.mDbPool.DbAppSvc.GetAppData(&xData)
	if xError!=nil{
		this.ReturnAppError(ctx,"数据获取失败!")
		return
	}

	this.ReturnMsg(ctx,"0","app.server.msg.common.op.succ",xData)
}



func (this *CmsCtl)CtlAppGatewayRuleSave(ctx *gin.Context)  {

	xIsLogin,_:=this.checkLogin(ctx)
	if !xIsLogin{
		return
	}

	xReqData:=ctldata.AppGatewayRuleSaveRequest{}
	xError:=ctx.BindJSON(&xReqData)
	if xError!=nil{
		return
	}

	xReqData.Data.ActionType="backend"

	xError=xReqData.Check()
	if xError!=nil{
		this.ReturnAppError(ctx,xError.Error())
		return
	}

	xData:=xReqData.Data

	xError=this.mDbPool.DbAppSvc.GetAppData(&xData)
	if xError!=nil{
		this.ReturnAppError(ctx,"操作失败,提交的数据不存在!")
		return
	}


	xData.ConfigDataId=xReqData.Data.ConfigDataId
	xData.SiteDataId=xReqData.Data.SiteDataId
	xData.RuleOrderNo=xReqData.Data.RuleOrderNo
	xData.RuleName=xReqData.Data.RuleName

	xData.MatchTarget=xReqData.Data.MatchTarget
	xData.MatchOp=xReqData.Data.MatchOp
	xData.MatchData=xReqData.Data.MatchData

	xData.ActionType=xReqData.Data.ActionType
	xData.ActionData=xReqData.Data.ActionData

	xData.LastSyncTime=cheerlib.TimeGetNow()
	xData.Status=xReqData.Data.Status

	xError=this.mDbPool.DbAppSvc.UpdateAppData(&xData)
	if xError!=nil{
		this.ReturnAppError(ctx,"操作失败,数据保存失败!")
		return
	}

	this.mDbPool.DbAppSvc.RenewAppGatewayConfigLastVer(xData.ConfigDataId)

	this.ReturnAppSuccess(ctx,"app.server.msg.common.op.succ")
}

func (this *CmsCtl)CtlAppGatewayRuleRemove(ctx *gin.Context)  {

	xIsLogin,_:=this.checkLogin(ctx)
	if !xIsLogin{
		return
	}

	xReqData:=ctldata.BaseGetInfoRequest{}
	xError:=ctx.BindJSON(&xReqData)
	if xError!=nil{
		return
	}

	xError=xReqData.Check()
	if xError!=nil{
		this.ReturnAppError(ctx,xError.Error())
		return
	}

	xError=this.mDbPool.DbAppSvc.RemoveAppGatewayRule(xReqData.DataId)
	if xError!=nil{
		this.ReturnAppError(ctx,xError.Error())
		return
	}

	this.ReturnAppSuccess(ctx,"app.server.msg.common.op.succ")
}