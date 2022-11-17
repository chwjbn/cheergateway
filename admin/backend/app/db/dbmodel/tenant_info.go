package dbmodel

import (
	"errors"
	"fmt"
	"cheeradmin/cheerlib"
	"strings"
)

type TenantInfo struct {
	TenantId string `bson:"tenant_id" json:"tenant_id"`
	UserName string `bson:"user_name" json:"user_name"`
	Password string `bson:"password" json:"password"`
	PwdSalt string `bson:"pwd_salt" json:"pwd_salt"`
	Email string `bson:"email" json:"email"`
	UserImgUrl string `bson:"user_img_url" json:"user_img_url"`
	CreateTime string `bson:"create_time" json:"create_time"`
	CreateIp string `bson:"create_ip" json:"create_ip"`
	Status string `bson:"status" json:"status"`   //void/active/disable
}

func (this *TenantInfo) Check() error  {

	var xError error=nil

	if len(this.UserName)<1{
		xError=errors.New("app.server.msg.tenant.username.invalid")
		return xError
	}

	if len(this.Password)<6||len(this.Password)>16{
		xError=errors.New("app.server.msg.tenant.password.invalid")
		return xError
	}


	if !strings.Contains(this.Email,"@"){
		xError=errors.New("app.server.msg.tenant.email.invalid")
		return xError
	}


	return xError
}

func (this *TenantInfo)EncryptPassword(pwd string,salt string) string  {

	return cheerlib.EncryptMd5(fmt.Sprintf("%s_%s",cheerlib.EncryptMd5(pwd),salt))
}

func (this *TenantInfo)Init()  {

	this.TenantId=cheerlib.EncryptMd5(fmt.Sprintf("%s_%s",this.UserName,this.Email))
	this.TenantId=cheerlib.EncryptMd5(fmt.Sprintf("%s_%s",this.TenantId,cheerlib.EncryptNewId()))

	this.PwdSalt=cheerlib.EncryptNewId()
	this.CreateTime=cheerlib.TimeGetNow()
	this.Status="void"

	this.Password=this.EncryptPassword(this.Password,this.PwdSalt)

}
