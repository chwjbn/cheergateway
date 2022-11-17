package ctl

import (
	"cheeradmin/app/ctl/ctldata"
	"cheeradmin/app/db/dbmodel/app_data"
	"cheeradmin/cheerlib"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)


func (this *CmsCtl)CtlAppGatewayBackendMapData(ctx *gin.Context)  {

	xIsLogin,_:=this.checkLogin(ctx)
	if !xIsLogin{
		return
	}


	xReqData:=ctldata.AppGatewayBackendMapDataRequest{}
	xError:=ctx.BindJSON(&xReqData)
	if xError!=nil{
		return
	}


	xData:=this.mDbPool.DbAppSvc.GetAppGatewayBackendMapData(xReqData.ConfigDataId)

	xRespData:=[]ctldata.DataMapNode{}

	for k,v:=range xData{
		xDataItem:=ctldata.DataMapNode{}
		xDataItem.DataId=k
		xDataItem.DataName=v
		xRespData=append(xRespData,xDataItem)
	}

	this.ReturnAppSuccessData(ctx,xRespData)


}

func (this *CmsCtl)CtlAppGatewayBackendPage(ctx *gin.Context)  {

	xIsLogin,_:=this.checkLogin(ctx)
	if !xIsLogin{
		return
	}

	xReqData:=ctldata.AppGatewayBackendPageRequest{}
	xError:=ctx.BindJSON(&xReqData)
	if xError!=nil{
		return
	}

	xWhere:=make(map[string]interface{})
	xSort:=make(map[string]interface{})

	if len(xReqData.ConfigDataId)>0{
		xWhere["config_data_id"]=xReqData.ConfigDataId
	}


	if len(xReqData.BackendName)>0{
		xWhere["backend_name"]=bson.M{"$regex":fmt.Sprintf("%s",xReqData.BackendName),"$options": "im"}
	}

	if len(xReqData.NodeAddr)>0{
		xWhere["node_addr"]=bson.M{"$regex":fmt.Sprintf("%s",xReqData.NodeAddr),"$options": "im"}
	}

	if len(xReqData.Status)>0{
		xWhere["status"]=xReqData.Status
	}

	xPageData:=this.mDbPool.DbAppSvc.GetAppGatewayBackendPageData(xReqData.PageNo,xReqData.PageSize,xWhere,xSort)

	this.ReturnPageData(ctx,xPageData)
}



func (this *CmsCtl)CtlAppGatewayBackendAdd(ctx *gin.Context)  {

	xIsLogin,xSessionInfo:=this.checkLogin(ctx)
	if !xIsLogin{
		return
	}

	xReqData:=ctldata.AppGatewayBackendAddRequest{}
	xError:=ctx.BindJSON(&xReqData)
	if xError!=nil{
		return
	}

	xError=xReqData.Check()
	if xError!=nil{
		this.ReturnAppError(ctx,xError.Error())
		return
	}

	xData:=app_data.AppGatewayBackend{}
	xData.ConfigDataId=xReqData.Data.ConfigDataId
	xData.BackendName=xReqData.Data.BackendName
	xData.NodeAddr=xReqData.Data.NodeAddr

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



func (this *CmsCtl)CtlAppGatewayBackendGet(ctx *gin.Context)  {

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

	xData:=app_data.AppGatewayBackend{}
	xData.DataId=xReqData.DataId

	xError=this.mDbPool.DbAppSvc.GetAppData(&xData)
	if xError!=nil{
		this.ReturnAppError(ctx,"数据获取失败!")
		return
	}

	this.ReturnMsg(ctx,"0","app.server.msg.common.op.succ",xData)
}


func (this *CmsCtl)CtlAppGatewayBackendSave(ctx *gin.Context)  {

	xIsLogin,_:=this.checkLogin(ctx)
	if !xIsLogin{
		return
	}

	xReqData:=ctldata.AppGatewayBackendSaveRequest{}
	xError:=ctx.BindJSON(&xReqData)
	if xError!=nil{
		return
	}

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
	xData.BackendName=xReqData.Data.BackendName
	xData.NodeAddr=xReqData.Data.NodeAddr
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


func (this *CmsCtl)CtlAppGatewayBackendRemove(ctx *gin.Context)  {

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

	xError=this.mDbPool.DbAppSvc.RemoveAppGatewayBackend(xReqData.DataId)
	if xError!=nil{
		this.ReturnAppError(ctx,xError.Error())
		return
	}

	this.ReturnAppSuccess(ctx,"app.server.msg.common.op.succ")
}