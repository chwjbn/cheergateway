package db

import (
	"cheeradmin/app/db/dbmodel"
	"cheeradmin/app/db/dbmodel/app_data"
	"cheeradmin/cheerlib"
	"errors"
	"fmt"
	"strings"
)

//把根据config_data_id映射config_data_name
func (this *DbAppSvc) AppGatewayConfigDataMapFilterHandler(dataMapList []map[string]interface{},srcDataMap map[string]map[string]interface{},idColName string,nameColName string) (error,[]map[string]interface{})  {

	var xError error=nil
	xDataMapList:=dataMapList

	for i:=0;i<len(dataMapList);i++{

		dataItem:=dataMapList[i]

		srcDataIdVal,bSrcDataId:=dataItem[idColName]
		if !bSrcDataId{
			continue
		}

		srcDataId:=fmt.Sprintf("%v",srcDataIdVal)

		srcData,bSrcData:=srcDataMap[srcDataId]
		if !bSrcData{
			continue
		}

		xDataMapList[i][nameColName]=fmt.Sprintf("%s",srcData["config_name"])

	}

	return xError,xDataMapList

}


//把根据backend_data_id映射backend_data_name
func (this *DbAppSvc) AppGatewayBackendDataMapFilterHandler(dataMapList []map[string]interface{},srcDataMap map[string]map[string]interface{},idColName string,nameColName string) (error,[]map[string]interface{})  {

	var xError error=nil
	xDataMapList:=dataMapList

	for i:=0;i<len(dataMapList);i++{

		dataItem:=dataMapList[i]

		srcDataIdVal,bSrcDataId:=dataItem[idColName]
		if !bSrcDataId{
			continue
		}

		srcDataId:=fmt.Sprintf("%v",srcDataIdVal)

		srcData,bSrcData:=srcDataMap[srcDataId]
		if !bSrcData{
			continue
		}

		xDataMapList[i][nameColName]=fmt.Sprintf("%s",srcData["backend_name"])

	}

	return xError,xDataMapList

}

//把根据site_data_id映射site_data_name
func (this *DbAppSvc) AppGatewaySiteDataMapFilterHandler(dataMapList []map[string]interface{},srcDataMap map[string]map[string]interface{},idColName string,nameColName string) (error,[]map[string]interface{})  {

	var xError error=nil
	xDataMapList:=dataMapList

	for i:=0;i<len(dataMapList);i++{

		dataItem:=dataMapList[i]

		srcDataIdVal,bSrcDataId:=dataItem[idColName]
		if !bSrcDataId{
			continue
		}

		srcDataId:=fmt.Sprintf("%v",srcDataIdVal)

		srcData,bSrcData:=srcDataMap[srcDataId]
		if !bSrcData{
			continue
		}

		xDataMapList[i][nameColName]=fmt.Sprintf("%s",srcData["site_name"])

	}

	return xError,xDataMapList

}

func (this *DbAppSvc)GetAppGatewayConfigMapData(envType string,mode string) map[string]string  {

	xModelData:=app_data.AppGatewayConfig{}

	xData:=make(map[string]string)

	xWhere:=make(map[string]interface{})
	xSort:=make(map[string]interface{})

	if len(envType)>0{
		xWhere["env_type"]=envType
	}

	xError,xDataList:=this.GetDataList(xModelData.GetTableName(),xWhere,xSort,-1,-1)

	if xError!=nil{
		return xData
	}

	for _,dItem:=range xDataList{

		xItem:=app_data.AppGatewayConfig{}
		xError=cheerlib.FlatMapToStruct(dItem,&xItem)
		if xError!=nil{
			continue
		}

		if strings.EqualFold(mode,"simple"){
			xData[xItem.DataId]=fmt.Sprintf("%s",xItem.ConfigName)
			continue
		}

		xData[xItem.DataId]=fmt.Sprintf("%s(%s)",xItem.ConfigName,xItem.ServerAddr)

	}

	return xData
}


func (this *DbAppSvc)GetAppGatewayConfigPageData(pageNo int,pageSize int,whereMap map[string]interface{},sortMap map[string]interface{}) dbmodel.PageData {

	xData:=app_data.AppGatewayConfig{}

	xTableName:=xData.GetTableName()

	xPageData:=this.GetDataPageList(xTableName,whereMap,sortMap,pageNo,pageSize)

	return xPageData
}

func (this *DbAppSvc)RenewAppGatewayConfigLastVer(dataId string) error  {

	var xError error=nil

	xData:=app_data.AppGatewayConfig{}
	xData.DataId=dataId

	this.mDbPool.DbAppSvc.GetAppData(&xData)

	if len(xData.Status)<1{
		xError=errors.New("要刷新的配置服务不存在!")
		return xError
	}

	xData.LastVer=cheerlib.TimeGetNow()

	this.mDbPool.DbAppSvc.UpdateAppData(&xData)

	return xError
}

func (this *DbAppSvc)RemoveAppGatewayConfig(dataId string) error  {

	var xError error=nil

	var xErr error=nil

	xBackendData:=app_data.AppGatewayBackend{}
	xBackendWhere:=make(map[string]interface{})
	xBackendWhere["config_data_id"]=dataId
	xErr=this.GetData(xBackendData.GetTableName(),xBackendWhere,&xBackendData)
	if xErr==nil{
		if len(xBackendData.DataId)>0{
			xError=errors.New("此配置服务下包含节点信息,操作失败!")
			return xError
		}
	}

	xSiteData:=app_data.AppGatewaySite{}
	xSiteWhere:=make(map[string]interface{})
	xSiteWhere["config_data_id"]=dataId
	xErr=this.GetData(xSiteData.GetTableName(),xSiteWhere,&xSiteData)
	if xErr==nil{
		if len(xSiteData.DataId)>0{
			xError=errors.New("此配置服务下包含站点信息,操作失败!")
			return xError
		}
	}

	xRuleData:=app_data.AppGatewayRule{}
	xRuleWhere:=make(map[string]interface{})
	xRuleWhere["config_data_id"]=dataId
	xErr=this.GetData(xRuleData.GetTableName(),xRuleWhere,&xRuleData)
	if xErr==nil{
		if len(xRuleData.DataId)>0{
			xError=errors.New("此配置服务下包含规则信息,操作失败!")
			return xError
		}
	}

	xData:=app_data.AppGatewayConfig{}
	xDataWhere:=make(map[string]interface{})
	xDataWhere["data_id"]=dataId

	this.DeleteData(xData.GetTableName(),xDataWhere)

	return xError

}


func (this *DbAppSvc)GetAppGatewayBackendPageData(pageNo int,pageSize int,whereMap map[string]interface{},sortMap map[string]interface{}) dbmodel.PageData {

	xData:=app_data.AppGatewayBackend{}
	xTableName:=xData.GetTableName()
	xPageData:=this.GetDataPageList(xTableName,whereMap,sortMap,pageNo,pageSize)

	if len(xPageData.DataList)<1{
		return xPageData
	}

	//映射config_data_id
	xConfigSrcTableName:=(&app_data.AppGatewayConfig{}).GetTableName()
	xConfigSrcDataIdColName:="config_data_id"
	xConfigDesDataNameColName:="config_data_name"

	xPageData.DataList=this.FillDataMapNameByDataId(xPageData.DataList,xConfigSrcDataIdColName,xConfigDesDataNameColName,xConfigSrcTableName,this.AppGatewayConfigDataMapFilterHandler)


	return xPageData
}

func (this *DbAppSvc)GetAppGatewayBackendMapData(configDataId string) map[string]string  {

	xModelData:=app_data.AppGatewayBackend{}

	xData:=make(map[string]string)

	xWhere:=make(map[string]interface{})
	xSort:=make(map[string]interface{})

	xWhere["config_data_id"]=configDataId

	xError,xDataList:=this.GetDataList(xModelData.GetTableName(),xWhere,xSort,-1,-1)

	if xError!=nil{
		return xData
	}

	for _,dItem:=range xDataList{

		xItem:=app_data.AppGatewayBackend{}
		xError=cheerlib.FlatMapToStruct(dItem,&xItem)
		if xError!=nil{
			continue
		}

		xData[xItem.DataId]=fmt.Sprintf("%s(%s)",xItem.BackendName,xItem.NodeAddr)
	}

	return xData
}

func (this *DbAppSvc)RemoveAppGatewayBackend(dataId string) error  {

	var xError error=nil

	var xErr error=nil


	xSiteData:=app_data.AppGatewaySite{}
	xSiteWhere:=make(map[string]interface{})
	xSiteWhere["default_backend_data_id"]=dataId
	xErr=this.GetData(xSiteData.GetTableName(),xSiteWhere,&xSiteData)
	if xErr==nil{
		if len(xSiteData.DataId)>0{
			xError=errors.New(fmt.Sprintf("此节点信息下包含站点信息[%s],操作失败!",xSiteData.SiteName))
			return xError
		}
	}

	xRuleData:=app_data.AppGatewayRule{}
	xRuleWhere:=make(map[string]interface{})
	xRuleWhere["action_data"]=dataId
	xErr=this.GetData(xRuleData.GetTableName(),xRuleWhere,&xRuleData)
	if xErr==nil{
		if len(xRuleData.DataId)>0{
			xError=errors.New(fmt.Sprintf("此节点信息下包含规则信息[%s],操作失败!",xRuleData.RuleName))
			return xError
		}
	}

	xData:=app_data.AppGatewayBackend{}

	xData.DataId=dataId
	this.GetAppData(&xData)
	this.RenewAppGatewayConfigLastVer(xData.ConfigDataId)


	xDataWhere:=make(map[string]interface{})
	xDataWhere["data_id"]=dataId

	this.DeleteData(xData.GetTableName(),xDataWhere)

	return xError

}

func (this *DbAppSvc)GetAppGatewaySitePageData(pageNo int,pageSize int,whereMap map[string]interface{},sortMap map[string]interface{}) dbmodel.PageData {

	xData:=app_data.AppGatewaySite{}
	xTableName:=xData.GetTableName()
	xPageData:=this.GetDataPageList(xTableName,whereMap,sortMap,pageNo,pageSize)

	if len(xPageData.DataList)<1{
		return xPageData
	}

	//映射config_data_id
	xConfigSrcTableName:=(&app_data.AppGatewayConfig{}).GetTableName()
	xConfigSrcDataIdColName:="config_data_id"
	xConfigDesDataNameColName:="config_data_name"

	xPageData.DataList=this.FillDataMapNameByDataId(xPageData.DataList,xConfigSrcDataIdColName,xConfigDesDataNameColName,xConfigSrcTableName,this.AppGatewayConfigDataMapFilterHandler)


	return xPageData
}


func (this *DbAppSvc)GetAppGatewaySiteMapData(configDataId string) map[string]string  {

	xModelData:=app_data.AppGatewaySite{}

	xData:=make(map[string]string)

	xWhere:=make(map[string]interface{})
	xSort:=make(map[string]interface{})

	xWhere["config_data_id"]=configDataId

	xError,xDataList:=this.GetDataList(xModelData.GetTableName(),xWhere,xSort,-1,-1)

	if xError!=nil{
		return xData
	}

	for _,dItem:=range xDataList{

		xItem:=app_data.AppGatewaySite{}
		xError=cheerlib.FlatMapToStruct(dItem,&xItem)
		if xError!=nil{
			continue
		}

		xData[xItem.DataId]=fmt.Sprintf("%s",xItem.SiteName)
	}

	return xData
}

func (this *DbAppSvc)RemoveAppGatewaySite(dataId string) error  {

	var xError error=nil

	var xErr error=nil


	xRuleData:=app_data.AppGatewayRule{}

	xRuleWhere:=make(map[string]interface{})
	xRuleWhere["site_data_id"]=dataId
	xErr=this.GetData(xRuleData.GetTableName(),xRuleWhere,&xRuleData)
	if xErr==nil{
		if len(xRuleData.DataId)>0{
			xError=errors.New(fmt.Sprintf("此站点信息下包含规则信息[%s],操作失败!",xRuleData.RuleName))
			return xError
		}
	}

	xData:=app_data.AppGatewaySite{}

	xData.DataId=dataId
	this.GetAppData(&xData)
	this.RenewAppGatewayConfigLastVer(xData.ConfigDataId)

	xDataWhere:=make(map[string]interface{})
	xDataWhere["data_id"]=dataId

	this.DeleteData(xData.GetTableName(),xDataWhere)

	return xError

}

func (this *DbAppSvc)GetAppGatewayRulePageData(pageNo int,pageSize int,whereMap map[string]interface{},sortMap map[string]interface{}) dbmodel.PageData {

	xData:=app_data.AppGatewayRule{}
	xTableName:=xData.GetTableName()
	xPageData:=this.GetDataPageList(xTableName,whereMap,sortMap,pageNo,pageSize)

	if len(xPageData.DataList)<1{
		return xPageData
	}

	//映射config_data_id
	xConfigSrcTableName:=(&app_data.AppGatewayConfig{}).GetTableName()
	xConfigSrcDataIdColName:="config_data_id"
	xConfigDesDataNameColName:="config_data_name"

	xPageData.DataList=this.FillDataMapNameByDataId(xPageData.DataList,xConfigSrcDataIdColName,xConfigDesDataNameColName,xConfigSrcTableName,this.AppGatewayConfigDataMapFilterHandler)


	//映射site_data_id
	xSiteSrcTableName:=(&app_data.AppGatewaySite{}).GetTableName()
	xSiteSrcDataIdColName:="site_data_id"
	xSiteDesDataNameColName:="site_data_name"

	xPageData.DataList=this.FillDataMapNameByDataId(xPageData.DataList,xSiteSrcDataIdColName,xSiteDesDataNameColName,xSiteSrcTableName,this.AppGatewaySiteDataMapFilterHandler)

	//映射action_data
	xActionSrcTableName:=(&app_data.AppGatewayBackend{}).GetTableName()
	xActionSrcDataIdColName:="action_data"
	xActionDesDataNameColName:="action_data_name"

	xPageData.DataList=this.FillDataMapNameByDataId(xPageData.DataList,xActionSrcDataIdColName,xActionDesDataNameColName,xActionSrcTableName,this.AppGatewayBackendDataMapFilterHandler)

	return xPageData
}

func (this *DbAppSvc)RemoveAppGatewayRule(dataId string) error  {

	var xError error=nil

	xData:=app_data.AppGatewayRule{}

	xData.DataId=dataId
	this.GetAppData(&xData)
	this.RenewAppGatewayConfigLastVer(xData.ConfigDataId)


	xDataWhere:=make(map[string]interface{})
	xDataWhere["data_id"]=dataId

	this.DeleteData(xData.GetTableName(),xDataWhere)



	return xError

}

func (this *DbAppSvc)GetAppGatewayDataListByConfigDataIdForPublish(tableName string,configDataId string,sortFieldName string) []map[string]interface{} {

	xDataList:=[]map[string]interface{}{}

	xWhere:=make(map[string]interface{})
	xSort:=make(map[string]interface{})

	xWhere["config_data_id"]=configDataId
	xWhere["status"]="enable"

	if len(sortFieldName)>0{
		xSort[sortFieldName]=1
	}

	xError,xTempDataList:=this.GetDataList(tableName,xWhere,xSort,-1,-1)

	if xError!=nil{
		return xDataList
	}

	// 去掉不必要的字段
	xSkipFields:="|_id|tenant_id|create_time|status|last_check_time|last_check_error|last_sync_time|last_sync_status|config_data_id|backend_name|site_name|rule_name|"

	for i:=0;i<len(xTempDataList);i++{

		xTempData:=xTempDataList[i]

		xDataItem:=make(map[string]interface{})

		for xKey,xVal:=range xTempData{

			if strings.Contains(xSkipFields,fmt.Sprintf("|%s|",xKey)){
				continue
			}

			xDataItem[xKey]=xVal
		}

		xDataList=append(xDataList,xDataItem)
	}

	return xDataList

}