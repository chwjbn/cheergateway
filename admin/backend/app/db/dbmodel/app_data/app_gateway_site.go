package app_data

import (
	"cheeradmin/cheerlib"
	"fmt"
)

type AppGatewaySite struct {
	AppData    `bson:",inline"`
	ConfigDataId string `bson:"config_data_id" json:"config_data_id"`
	SiteOrderNo int `bson:"site_order_no" json:"site_order_no"`
	SiteName string `bson:"site_name" json:"site_name"`
	RuleType string `bson:"rule_type" json:"rule_type"`
	RuleData string `bson:"rule_data" json:"rule_data"`
	DefaultBackendDataId string `bson:"default_backend_data_id" json:"default_backend_data_id"`
}



func (this *AppGatewaySite) Init(tenantId string) error {

	var xError error

	xError=this.InitData(this.GetTableName(),tenantId)
	this.DataId=cheerlib.EncryptMd5(fmt.Sprintf("%s-%s-%s-%s-%s",this.GetDataId(),this.ConfigDataId,this.SiteName,this.RuleType,this.RuleData))
	this.Status="enable"

	return xError
}

func (this *AppGatewaySite)GetTableName() string  {
	return "app_gateway_site"
}

func (this *AppGatewaySite)GetTableIndexFieldList() []string  {

	dataList:= []string{"site_name","site_order_no","rule_type"}

	xDataList:=this.GetAppDataTableIndexFieldList()

	for _,xDataItem:=range dataList{
		xDataList=append(xDataList,xDataItem)
	}

	return xDataList
}