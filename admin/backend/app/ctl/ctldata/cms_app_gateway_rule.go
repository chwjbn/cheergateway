package ctldata

import (
	"cheeradmin/app/db/dbmodel/app_data"
	"errors"
	"fmt"
	"strings"
)

type AppGatewayRulePageRequest struct {
	BasePageRequest
	RuleName string `json:"rule_name"`
	ConfigDataId string `json:"config_data_id"`
	SiteDataId string `json:"site_data_id"`
	Status string `json:"status"`
}


type AppGatewayRuleAddRequest struct {
	Data app_data.AppGatewayRule `json:"data"`
}

type AppGatewayRuleSaveRequest struct {
	Data app_data.AppGatewayRule `json:"data"`
}


func (this *AppGatewayRuleAddRequest)Check() error  {

	var xError error=nil


	if len(this.Data.ConfigDataId)<1{
		xError=errors.New("请选择配置服务!")
		return xError
	}

	if len(this.Data.RuleName)<1||len(this.Data.RuleName)>100{
		xError=errors.New("规则名称长度为1-100!")
		return xError
	}

	if len(this.Data.MatchTarget)<1||len(this.Data.MatchTarget)>100{
		xError=errors.New("规则匹配对象长度为1-100!")
		return xError
	}

	xMatchOpRange:="|regex|contain|in|notregex|notcontain|notin|eq|neq|lt|gt|lte|gte|"
	if !strings.Contains(xMatchOpRange,fmt.Sprintf("|%s|",this.Data.MatchOp)){
		xError=errors.New("请选择正确的规则匹配操作符!")
		return xError
	}

	if len(this.Data.MatchData)<1||len(this.Data.MatchData)>100{
		xError=errors.New("规则匹配内容长度为1-100!")
		return xError
	}

	xActionTypeRange:="|backend|page|content|"
	if !strings.Contains(xActionTypeRange,fmt.Sprintf("|%s|",this.Data.ActionType)){
		xError=errors.New("请选择正确的动作执行类别!")
		return xError
	}

	if len(this.Data.ActionData)<1{
		xError=errors.New("请选择动作执行内容!")
		return xError
	}

	return xError

}



func (this *AppGatewayRuleSaveRequest)Check() error  {

	var xError error=nil


	if len(this.Data.ConfigDataId)<1{
		xError=errors.New("请选择配置服务!")
		return xError
	}

	if len(this.Data.SiteDataId)<1{
		xError=errors.New("请选择规则站点!")
		return xError
	}

	if len(this.Data.RuleName)<1||len(this.Data.RuleName)>100{
		xError=errors.New("规则名称长度为1-100!")
		return xError
	}

	if len(this.Data.MatchTarget)<1||len(this.Data.MatchTarget)>100{
		xError=errors.New("规则匹配对象长度为1-100!")
		return xError
	}

	xMatchOpRange:="|regex|contain|in|notregex|notcontain|notin|eq|neq|lt|gt|lte|gte|"
	if !strings.Contains(xMatchOpRange,fmt.Sprintf("|%s|",this.Data.MatchOp)){
		xError=errors.New("请选择正确的规则匹配操作符!")
		return xError
	}

	if len(this.Data.MatchData)<1||len(this.Data.MatchData)>100{
		xError=errors.New("规则匹配内容长度为1-100!")
		return xError
	}

	xActionTypeRange:="|backend|page|content|"
	if !strings.Contains(xActionTypeRange,fmt.Sprintf("|%s|",this.Data.ActionType)){
		xError=errors.New("请选择正确的动作执行类别!")
		return xError
	}

	if len(this.Data.ActionData)<1{
		xError=errors.New("请选择动作执行内容!")
		return xError
	}

	return xError

}