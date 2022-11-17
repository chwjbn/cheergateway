package app_data

import (
	"errors"
	"fmt"
	"cheeradmin/cheerlib"
	"time"
)

type IAppData interface {
	GetTableName() string
	GetTableIndexFieldList() []string
	Init(tenantId string) error

	GetTenantId() string
	GetDataId() string

	SetTenantId(data string)
	SetDataId(data string)

	GetLastCheckTime() string
	SetLastCheckTime(data string)
	SetLastCheckError(data error)

	CheckAppData() error

}

type AppData struct {
	TenantId string `bson:"tenant_id" json:"tenant_id"`
	DataId string `bson:"data_id" json:"data_id"`
	CreateTime string `bson:"create_time" json:"create_time"`
	Status string `bson:"status" json:"status"`

	LastCheckTime string `bson:"last_check_time" json:"last_check_time"`
	LastCheckError string `bson:"last_check_error" json:"last_check_error"`

	LastSyncTime string `bson:"last_sync_time" json:"last_sync_time"`
	LastSyncStatus string `bson:"last_sync_status" json:"last_sync_status"`
}

func (this *AppData) GetAppDataTableIndexFieldList() []string  {

	xDataList:=[]string{"tenant_id","data_id","create_time","status","last_check_time","last_sync_time","last_sync_status"}

	for _,xData:=range xDataList{
		xDataList=append(xDataList,xData)
	}

	return xDataList

}

func (this *AppData)InitData(tableName string,tenantId string) error  {

	this.TenantId=tenantId

	if len(this.TenantId)<10{
		return errors.New("invalid tenant_id.")
	}

	if len(this.DataId)<1{
		this.DataId=cheerlib.EncryptMd5(fmt.Sprintf("%s-%s-%s",tableName,this.TenantId,cheerlib.EncryptNewId()))
	}


	this.CreateTime=cheerlib.TimeGetNow()
	this.Status="init"

	this.LastCheckTime=cheerlib.TimeGetTime(time.Now().Add(-100*time.Hour))
	this.LastCheckError=""

	this.LastSyncTime=cheerlib.TimeGetNow()
	this.LastSyncStatus="normal"

	return nil
}

func (this *AppData) GetTenantId() string {
	return this.TenantId
}

func (this *AppData) GetDataId() string {
	return this.DataId
}

func (this *AppData)SetTenantId(data string)  {
	this.TenantId=data
}

func (this *AppData)SetDataId(data string)  {
	this.DataId=data
}

func (this *AppData)GetLastCheckTime() string  {
	return this.LastCheckTime
}

func (this *AppData) SetLastCheckTime(data string) {
	this.LastCheckTime=data
}

func (this *AppData)SetLastCheckError(data error)  {

	if data==nil{
		this.LastCheckError=""
		this.LastSyncStatus="normal"
		return
	}

	this.LastCheckError=data.Error()
	this.LastSyncStatus="error"
}


func (this *AppData)CheckAppData() error  {
	var xError error=nil

	if len(this.TenantId)<10{
		xError=errors.New("invalid tenant_id.")
		return xError
	}

	if len(this.DataId)<10{
		xError=errors.New("invalid data_id.")
		return xError
	}


	return xError
}