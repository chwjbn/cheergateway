package app_data

import (
	"cheeradmin/cheerlib"
	"fmt"
)

type AppGatewayConfig struct {
	AppData    `bson:",inline"`
	EnvType string `bson:"env_type" json:"env_type"`
	ConfigName string `bson:"config_name" json:"config_name"`
	ServerAddr string `bson:"server_addr" json:"server_addr"`
	UserName string `bson:"user_name" json:"user_name"`
	Password string `bson:"password" json:"password"`
	LastPubVer string `bson:"last_pub_ver" json:"last_pub_ver"`
	LastVer string  `bson:"last_ver" json:"last_ver"`
}


func (this *AppGatewayConfig) Init(tenantId string) error {

	var xError error

	xError=this.InitData(this.GetTableName(),tenantId)
	this.DataId=cheerlib.EncryptMd5(fmt.Sprintf("%s-%s-%s",this.GetDataId(),this.EnvType,this.ConfigName))
	this.Status="enable"

	return xError
}

func (this *AppGatewayConfig)GetTableName() string  {
	return "app_gateway_config"
}

func (this *AppGatewayConfig)GetTableIndexFieldList() []string  {

	dataList:= []string{"config_name","env_type","last_pub_ver","last_ver"}

	xDataList:=this.GetAppDataTableIndexFieldList()

	for _,xDataItem:=range dataList{
		xDataList=append(xDataList,xDataItem)
	}

	return xDataList
}