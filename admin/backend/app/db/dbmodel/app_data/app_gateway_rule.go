package app_data

import (
	"cheeradmin/cheerlib"
	"fmt"
)

type AppGatewayRule struct {
	AppData    `bson:",inline"`
	ConfigDataId string `bson:"config_data_id" json:"config_data_id"`
	SiteDataId string `bson:"site_data_id" json:"site_data_id"`
	RuleOrderNo int `bson:"rule_order_no" json:"rule_order_no"`
	RuleName string `bson:"rule_name" json:"rule_name"`

	MatchTarget string  `bson:"match_target" json:"match_target"`  //
	MatchOp string  `bson:"match_op" json:"match_op"`       //
	MatchData string  `bson:"match_data" json:"match_data"`

	ActionType string  `bson:"action_type" json:"action_type"`   //|backend|page|content|
	ActionData string   `bson:"action_data" json:"action_data"`
}


func (this *AppGatewayRule) Init(tenantId string) error {

	var xError error

	xError=this.InitData(this.GetTableName(),tenantId)
	this.DataId=cheerlib.EncryptMd5(fmt.Sprintf("%s-%s-%s-%s-%s",this.GetDataId(),this.ConfigDataId,this.SiteDataId,this.RuleName,this.MatchTarget))
	this.Status="enable"

	return xError
}

func (this *AppGatewayRule)GetTableName() string  {
	return "app_gateway_rule"
}

func (this *AppGatewayRule)GetTableIndexFieldList() []string  {

	dataList:= []string{"rule_name","config_data_id","site_data_id","rule_order_no","match_target","match_op","action_type","action_data"}

	xDataList:=this.GetAppDataTableIndexFieldList()

	for _,xDataItem:=range dataList{
		xDataList=append(xDataList,xDataItem)
	}

	return xDataList
}