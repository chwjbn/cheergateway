package db

import (
	"cheeradmin/app/db/dbmodel/app_data"
	"cheeradmin/cheerlib"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

type DbAppSvc struct {
	DbModelSvc
}

type DataMapFilterHandler func([]map[string]interface{},map[string]map[string]interface{},string,string) (error,[]map[string]interface{})


func (this *DbAppSvc)AddAppData(data app_data.IAppData) error  {

	var xError error=nil

	xError=data.CheckAppData()
	if xError!=nil{
		return xError
	}

	xTableName:=data.GetTableName()

	xError=this.AddData(xTableName,data)

	if xError!=nil{
		return xError
	}

	xIndexList:=data.GetTableIndexFieldList()
	for _,xIndexItem:=range xIndexList{

		if strings.EqualFold(xIndexItem,"data_id"){
			this.MakeSureIndexUnique(xTableName,[]string{xIndexItem})
			continue
		}
		this.MakeSureIndex(xTableName,[]string{xIndexItem})
	}

	return xError
}

func (this *DbAppSvc)UpdateAppData(data app_data.IAppData) error  {

	var xError error=nil

	defer func() {
		if xError!=nil{
			cheerlib.LogError(fmt.Sprintf("DbAppSvc.UpdateAppData Error=[%s]",xError.Error()))
		}

	}()

	xError=data.CheckAppData()
	if xError!=nil{
		return xError
	}

	xTableName:=data.GetTableName()
	xWhere:=make(map[string]interface{})
	xWhere["data_id"]=data.GetDataId()

	xError=this.SaveData(xTableName,xWhere,data)

	return xError
}

func (this *DbAppSvc)GetAppData(data app_data.IAppData) error  {

	var xError error=nil

	xTableName:=data.GetTableName()
	xWhere:=make(map[string]interface{})
	xWhere["data_id"]=data.GetDataId()

	xError=this.GetData(xTableName,xWhere,data)

	return xError
}

func (this *DbAppSvc)GetAppDataListByDataIds(tableName string,dataIdList []interface{})(error,map[string]map[string]interface{})  {

	var xError error=nil
	xDataMap:=make(map[string]map[string]interface{})

	whereMap:=make(map[string]interface{})
	whereMap["data_id"]=bson.M{"$in":dataIdList}

	sortMap:=make(map[string]interface{})

	xError,xDataMapList:=this.GetDataList(tableName,whereMap,sortMap,-1,-1)

	for _,dItem:=range xDataMapList{
		_,bFind:=dItem["data_id"]
		if !bFind{
			continue
		}

		xDataMap[fmt.Sprintf("%s",dItem["data_id"])]=dItem
	}

	return xError,xDataMap
}

func (this *DbAppSvc)FillDataMapNameByDataId(dataMapList []map[string]interface{},dataIdColName string,dataNameColName string,srcDataTableName string,dataFilter DataMapFilterHandler) []map[string]interface{}  {

	xDataMapList:=dataMapList

	xIdList:=cheerlib.TextGetMapColumn(dataMapList,dataIdColName)

	if len(xIdList)<1{
		return xDataMapList
	}

    xError,xSrcDataMapList:=this.GetAppDataListByDataIds(srcDataTableName,xIdList)
    if xError!=nil{
    	return xDataMapList
	}

	xError,xDataMapList=dataFilter(dataMapList,xSrcDataMapList,dataIdColName,dataNameColName)

	return xDataMapList

}