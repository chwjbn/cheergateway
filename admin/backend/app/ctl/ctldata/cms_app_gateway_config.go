package ctldata

import (
	"cheeradmin/app/db/dbmodel/app_data"
	"errors"
	"fmt"
	"strings"
)

type AppGatewayConfigPageRequest struct {
	BasePageRequest
	ConfigName string `json:"config_name"`
	EnvType string `json:"env_type"`
	Status string `json:"status"`
}

type AppGatewayConfigAddRequest struct {
	Data app_data.AppGatewayConfig `json:"data"`
}

type AppGatewayConfigSaveRequest struct {
	Data app_data.AppGatewayConfig `json:"data"`
}

type AppGatewayConfigMapDataRequest struct {
	EnvType string `json:"env_type"`
	Mode string `json:"mode"`
}


func (this *AppGatewayConfigAddRequest)Check() error  {

	var xError error=nil

	xEnvTypeMatch:="|prod|test|dev|"

	if !strings.Contains(xEnvTypeMatch,fmt.Sprintf("|%s|",this.Data.EnvType)){
		xError=errors.New("请选择正确的配置环境类别!")
		return xError
	}


	if len(this.Data.ConfigName)<1||len(this.Data.ConfigName)>100{
		xError=errors.New("配置名称长度1-100!")
		return xError
	}

	if len(this.Data.ServerAddr)<2||!strings.Contains(this.Data.ServerAddr,":"){
		xError=errors.New("配置服务器Redis地址格式不正确!")
		return xError
	}

	return xError

}


func (this *AppGatewayConfigSaveRequest)Check() error  {

	var xError error=nil

	if len(this.Data.DataId)<1{
		xError=errors.New("提交的数据缺失数据ID!")
		return xError
	}

	xEnvTypeMatch:="|prod|test|dev|"

	if !strings.Contains(xEnvTypeMatch,fmt.Sprintf("|%s|",this.Data.EnvType)){
		xError=errors.New("请选择正确的配置环境类别!")
		return xError
	}


	if len(this.Data.ConfigName)<1||len(this.Data.ConfigName)>100{
		xError=errors.New("配置名称长度1-100!")
		return xError
	}

	if len(this.Data.ServerAddr)<2||!strings.Contains(this.Data.ServerAddr,":"){
		xError=errors.New("配置服务器Redis地址格式不正确!")
		return xError
	}

	return xError

}
