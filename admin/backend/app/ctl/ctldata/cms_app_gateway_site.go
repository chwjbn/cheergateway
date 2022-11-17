package ctldata

import (
	"cheeradmin/app/db/dbmodel/app_data"
	"errors"
	"fmt"
	"strings"
)

type AppGatewaySitePageRequest struct {
	BasePageRequest
	SiteName string `json:"site_name"`
	ConfigDataId string `json:"config_data_id"`
	Status string `json:"status"`
}


type AppGatewaySiteAddRequest struct {
	Data app_data.AppGatewaySite `json:"data"`
}

type AppGatewaySiteSaveRequest struct {
	Data app_data.AppGatewaySite `json:"data"`
}

type AppGatewaySiteMapDataRequest struct {
	ConfigDataId string `json:"config_data_id"`
}

func (this *AppGatewaySiteAddRequest)Check() error  {

	var xError error=nil


	if len(this.Data.ConfigDataId)<1{
		xError=errors.New("请选择配置服务!")
		return xError
	}

	if len(this.Data.SiteName)<1||len(this.Data.SiteName)>100{
		xError=errors.New("站点名称长度为1-100!")
		return xError
	}

	xRuleTypeRange:="|eq|regex|contain|in|neq|notregex|notcontain|notin|"
	if !strings.Contains(xRuleTypeRange,fmt.Sprintf("|%s|",this.Data.RuleType)){
		xError=errors.New("请选择正确的站点域名匹配方式!")
		return xError
	}

	if len(this.Data.RuleData)<1||len(this.Data.RuleData)>100{
		xError=errors.New("站点域名匹配内容长度为1-100!")
		return xError
	}

	if len(this.Data.DefaultBackendDataId)<1{
		xError=errors.New("请选择默认后端节点!")
		return xError
	}

	return xError

}

func (this *AppGatewaySiteSaveRequest)Check() error  {

	var xError error=nil


	if len(this.Data.ConfigDataId)<1{
		xError=errors.New("请选择配置服务!")
		return xError
	}

	if len(this.Data.SiteName)<1||len(this.Data.SiteName)>100{
		xError=errors.New("站点名称长度为1-100!")
		return xError
	}

	xRuleTypeRange:="|eq|regex|contain|in|neq|notregex|notcontain|notin|"
	if !strings.Contains(xRuleTypeRange,fmt.Sprintf("|%s|",this.Data.RuleType)){
		xError=errors.New("请选择正确的站点域名匹配方式!")
		return xError
	}

	if len(this.Data.RuleData)<1||len(this.Data.RuleData)>100{
		xError=errors.New("站点域名匹配内容长度为1-100!")
		return xError
	}

	if len(this.Data.DefaultBackendDataId)<1{
		xError=errors.New("请选择默认后端节点!")
		return xError
	}

	return xError

}