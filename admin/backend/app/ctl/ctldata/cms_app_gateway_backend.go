package ctldata

import (
	"cheeradmin/app/db/dbmodel/app_data"
	"cheeradmin/cheerlib"
	"errors"
)

type AppGatewayBackendPageRequest struct {
	BasePageRequest
	BackendName string `json:"backend_name"`
	NodeAddr string `json:"node_addr"`
	ConfigDataId string `json:"config_data_id"`
	Status string `json:"status"`
}


type AppGatewayBackendAddRequest struct {
	Data app_data.AppGatewayBackend `json:"data"`
}

type AppGatewayBackendSaveRequest struct {
	Data app_data.AppGatewayBackend `json:"data"`
}

type AppGatewayBackendMapDataRequest struct {
	ConfigDataId string `json:"config_data_id"`
}

func (this *AppGatewayBackendAddRequest)Check() error  {

	var xError error=nil

	if len(this.Data.ConfigDataId)<1{
		xError=errors.New("请选择配置服务!")
		return xError
	}

	if len(this.Data.BackendName)<1||len(this.Data.BackendName)>100{
		xError=errors.New("节点名称长度1-100!")
		return xError
	}

	if !cheerlib.TextIsNodeAddrList(this.Data.NodeAddr){
		xError=errors.New("节点服务器地址格式不正确!")
		return xError
	}


	return xError

}

func (this *AppGatewayBackendSaveRequest)Check() error  {

	var xError error=nil

	if len(this.Data.ConfigDataId)<1{
		xError=errors.New("请选择配置服务!")
		return xError
	}

	if len(this.Data.BackendName)<1||len(this.Data.BackendName)>100{
		xError=errors.New("节点名称长度1-100!")
		return xError
	}


	if !cheerlib.TextIsNodeAddrList(this.Data.NodeAddr){
		xError=errors.New("节点服务器地址格式不正确!")
		return xError
	}

	return xError

}