package ctl

import (
	"cheeradmin/app/ctl/ctldata"
	"cheeradmin/app/db/dbmodel/app_data"
	"cheeradmin/cheerlib"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (this *CmsCtl)CtlAppGatewayConfigMapData(ctx *gin.Context)  {

	xIsLogin,_:=this.checkLogin(ctx)
	if !xIsLogin{
		return
	}


	xReqData:=ctldata.AppGatewayConfigMapDataRequest{}
	xError:=ctx.BindJSON(&xReqData)
	if xError!=nil{
		return
	}

	xData:=this.mDbPool.DbAppSvc.GetAppGatewayConfigMapData(xReqData.EnvType,xReqData.Mode)

	xRespData:=[]ctldata.DataMapNode{}

	for k,v:=range xData{
		xDataItem:=ctldata.DataMapNode{}
		xDataItem.DataId=k
		xDataItem.DataName=v
		xRespData=append(xRespData,xDataItem)
	}

	this.ReturnAppSuccessData(ctx,xRespData)


}

func (this *CmsCtl)CtlAppGatewayConfigPage(ctx *gin.Context)  {

	xIsLogin,_:=this.checkLogin(ctx)
	if !xIsLogin{
		return
	}

	xReqData:=ctldata.AppGatewayConfigPageRequest{}
	xError:=ctx.BindJSON(&xReqData)
	if xError!=nil{
		return
	}

	xWhere:=make(map[string]interface{})
	xSort:=make(map[string]interface{})

	if len(xReqData.ConfigName)>0{
		xWhere["config_name"]=bson.M{"$regex":fmt.Sprintf("%s",xReqData.ConfigName),"$options": "im"}
	}

	if len(xReqData.EnvType)>0{
		xWhere["env_type"]=xReqData.EnvType
	}

	if len(xReqData.Status)>0{
		xWhere["status"]=xReqData.Status
	}

	xPageData:=this.mDbPool.DbAppSvc.GetAppGatewayConfigPageData(xReqData.PageNo,xReqData.PageSize,xWhere,xSort)

	this.ReturnPageData(ctx,xPageData)
}

func (this *CmsCtl)CtlAppGatewayConfigAdd(ctx *gin.Context)  {

	xIsLogin,xSessionInfo:=this.checkLogin(ctx)
	if !xIsLogin{
		return
	}

	xReqData:=ctldata.AppGatewayConfigAddRequest{}
	xError:=ctx.BindJSON(&xReqData)
	if xError!=nil{
		return
	}

	xError=xReqData.Check()
	if xError!=nil{
		this.ReturnAppError(ctx,xError.Error())
		return
	}

	xData:=app_data.AppGatewayConfig{}
	xData.EnvType=xReqData.Data.EnvType
	xData.ConfigName=xReqData.Data.ConfigName
	xData.ServerAddr=xReqData.Data.ServerAddr
	xData.UserName=xReqData.Data.UserName
	xData.Password=xReqData.Data.Password
	xData.LastPubVer="1971-01-01 00:00:00"
	xData.LastVer=cheerlib.TimeGetNow()

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

	this.ReturnAppSuccess(ctx,"app.server.msg.common.op.succ")
}

func (this *CmsCtl)CtlAppGatewayConfigGet(ctx *gin.Context)  {

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

	xData:=app_data.AppGatewayConfig{}
	xData.DataId=xReqData.DataId

	xError=this.mDbPool.DbAppSvc.GetAppData(&xData)
	if xError!=nil{
		this.ReturnAppError(ctx,"数据获取失败!")
		return
	}

	this.ReturnMsg(ctx,"0","app.server.msg.common.op.succ",xData)
}

func (this *CmsCtl)CtlAppGatewayConfigSave(ctx *gin.Context)  {

	xIsLogin,_:=this.checkLogin(ctx)
	if !xIsLogin{
		return
	}

	xReqData:=ctldata.AppGatewayConfigSaveRequest{}
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

	xData.EnvType=xReqData.Data.EnvType
	xData.ConfigName=xReqData.Data.ConfigName
	xData.ServerAddr=xReqData.Data.ServerAddr
	xData.UserName=xReqData.Data.UserName
	xData.Password=xReqData.Data.Password
	xData.LastSyncTime=cheerlib.TimeGetNow()
	xData.Status=xReqData.Data.Status
	xData.LastVer=cheerlib.TimeGetNow()

	xError=this.mDbPool.DbAppSvc.UpdateAppData(&xData)
	if xError!=nil{
		this.ReturnAppError(ctx,"操作失败,数据保存失败!")
		return
	}

	this.ReturnAppSuccess(ctx,"app.server.msg.common.op.succ")
}

func (this *CmsCtl)CtlAppGatewayConfigRemove(ctx *gin.Context)  {

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

	xError=this.mDbPool.DbAppSvc.RemoveAppGatewayConfig(xReqData.DataId)
	if xError!=nil{
		this.ReturnAppError(ctx,xError.Error())
		return
	}

	this.ReturnAppSuccess(ctx,"app.server.msg.common.op.succ")
}

func (this *CmsCtl)CtlAppGatewayConfigPublish(ctx *gin.Context)  {

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

	xData:=app_data.AppGatewayConfig{}
	xData.DataId=xReqData.DataId

	xError=this.mDbPool.DbAppSvc.GetAppData(&xData)
	if xError!=nil{
		this.ReturnAppError(ctx,"配置数据获取失败,操作失败!")
		return
	}


	xBackendData:=app_data.AppGatewayBackend{}
	xSiteData:=app_data.AppGatewaySite{}
	xRuleData:=app_data.AppGatewayRule{}

	xBackendPubData:=this.mDbPool.DbAppSvc.GetAppGatewayDataListByConfigDataIdForPublish(xBackendData.GetTableName(),xData.DataId,"")
	if len(xBackendPubData)<1{
		this.ReturnAppError(ctx,"此配置下没有有效的节点信息,操作失败!")
		return
	}

	xSitePubData:=this.mDbPool.DbAppSvc.GetAppGatewayDataListByConfigDataIdForPublish(xSiteData.GetTableName(),xData.DataId,"site_order_no")
	if len(xSitePubData)<1{
		this.ReturnAppError(ctx,"此配置下没有有效的站点信息,操作失败!")
		return
	}

	xRulePubData:=this.mDbPool.DbAppSvc.GetAppGatewayDataListByConfigDataIdForPublish(xRuleData.GetTableName(),xData.DataId,"rule_order_no")


	xRedisConn,xRedisErr:=redis.Dial("tcp", xData.ServerAddr,redis.DialConnectTimeout(5*time.Second),redis.DialDatabase(0),redis.DialUsername(xData.UserName),redis.DialPassword(xData.Password))
	if xRedisErr!=nil{
		this.ReturnAppError(ctx,fmt.Sprintf("连接到配置服务器[%s]错误[%s],操作失败!",xData.ServerAddr,xRedisErr.Error()))
		return
	}

	defer func() {
		xRedisConn.Close()
	}()


	// 清理节点历史数据
	xKeys,xKeysErr:=redis.Strings(xRedisConn.Do("KEYS","cheer-gateway-*"))
	if xKeysErr==nil{
		for _,mKey:=range xKeys{
			xRedisConn.Do("DEL",mKey)
		}
	}


	xRedisConn.Do("SET","cheer-gateway-site",cheerlib.TextStructToJson(xSitePubData))

	for i:=0;i<len(xBackendPubData);i++{
		xSetItem:=xBackendPubData[i]
		xSetKey:=fmt.Sprintf("cheer-gateway-backend-%s",xSetItem["data_id"])
		xRedisConn.Do("SET",xSetKey,cheerlib.TextStructToJson(xSetItem))
	}


	xRulePubDataMap:=make(map[string][]map[string]interface{})
	for i:=0;i<len(xRulePubData);i++{
		xDataItem:=xRulePubData[i]
		xSiteDataId:=xDataItem["site_data_id"].(string)
		_,bFind:=xRulePubDataMap[xSiteDataId]
		if !bFind{
			xRulePubDataMap[xSiteDataId]=[]map[string]interface{}{}
		}

		xRulePubDataMap[xSiteDataId]=append(xRulePubDataMap[xSiteDataId],xDataItem)
	}

	for mKey,mVal:=range xRulePubDataMap{
		xSetItem:=mVal
		xSetKey:=fmt.Sprintf("cheer-gateway-rule-%s",mKey)
		xRedisConn.Do("SET",xSetKey,cheerlib.TextStructToJson(xSetItem))
	}

	xData.LastPubVer=xData.LastVer
	this.mDbPool.DbAppSvc.UpdateAppData(&xData)

	this.ReturnMsg(ctx,"0","app.server.msg.common.op.succ",xData)
}