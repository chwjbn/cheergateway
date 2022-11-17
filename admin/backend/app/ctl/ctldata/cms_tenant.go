package ctldata

import (
	"cheeradmin/app/db/dbmodel"
	"errors"
)


type TenantLoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	CheckCodeImgId string `json:"check_code_img_id"`
	CheckCodeImgData string `json:"check_code_img_data"`
}

type TenantLoginRespData struct {
	UserName string `json:"user_name"`
	TokenId string `json:"token_id"`
	TenantId string `json:"tenant_id"`
	UserImgUrl string `json:"user_img_url"`
	Role string `json:"role"`
}

type TenantInfoRequest struct {
	Data dbmodel.TenantInfo `json:"data"`
}

func (this *TenantInfoRequest)CheckInputForUpdateInfo() error  {

	var xError error=nil

	if len(this.Data.TenantId)<1{
		xError=errors.New("账号ID信息缺失!")
		return xError
	}

	if len(this.Data.Email)<1{
		xError=errors.New("请输入正确的账号邮箱地址!")
		return xError
	}


	return xError
}

func (this *TenantInfoRequest)CheckInputForUpdatePassword() error  {

	var xError error=nil

	if len(this.Data.TenantId)<1{
		xError=errors.New("账号ID信息缺失!")
		return xError
	}

	if len(this.Data.Password)<1{
		xError=errors.New("请输入正确的密码!")
		return xError
	}

	return xError
}

func (this *TenantLoginRequest)CheckInput() error  {

	var xError error=nil

	if len(this.UserName)<1{
		xError=errors.New("app.server.msg.tenant.username.invalid")
		return xError
	}

	if len(this.Password)<6||len(this.Password)>16{
		xError=errors.New("app.server.msg.tenant.password.invalid")
		return xError
	}

	if len(this.CheckCodeImgId)<6||len(this.CheckCodeImgData)<1{
		xError=errors.New("app.server.msg.tenant.checkimgcode.invalid")
		return xError
	}

	return xError

}