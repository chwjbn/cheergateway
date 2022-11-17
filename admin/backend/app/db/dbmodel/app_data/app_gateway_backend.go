package app_data

import (
	"cheeradmin/cheerlib"
	"fmt"
)

type AppGatewayBackend struct {
	AppData    `bson:",inline"`
	ConfigDataId string `bson:"config_data_id" json:"config_data_id"`
	BackendName string `bson:"backend_name" json:"backend_name"`
	NodeAddr string `bson:"node_addr" json:"node_addr"`
}

func (this *AppGatewayBackend) Init(tenantId string) error {

	var xError error

	xError=this.InitData(this.GetTableName(),tenantId)
	this.DataId=cheerlib.EncryptMd5(fmt.Sprintf("%s-%s-%s",this.GetDataId(),this.ConfigDataId,this.BackendName))
	this.Status="enable"

	return xError
}

func (this *AppGatewayBackend)GetTableName() string  {
	return "app_gateway_backend"
}

func (this *AppGatewayBackend)GetTableIndexFieldList() []string  {

	dataList:= []string{"backend_name","config_data_id"}

	xDataList:=this.GetAppDataTableIndexFieldList()

	for _,xDataItem:=range dataList{
		xDataList=append(xDataList,xDataItem)
	}

	return xDataList
}