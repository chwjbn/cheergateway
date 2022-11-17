package dbmodel

import (
	"fmt"
	"cheeradmin/cheerlib"
)

type TenantAccessSession struct {
	TokenId string `bson:"token_id" json:"token_id"`
	TenantId string `bson:"tenant_id" json:"tenant_id"`
	LoginTime string `bson:"login_time" json:"login_time"`
	LoginAddr string `bson:"login_addr" json:"login_addr"`
	LastTime string `bson:"last_time" json:"last_time"`
	LastAddr string `bson:"last_addr" json:"last_addr"`
}

func (this *TenantAccessSession)Init()  {


	this.LoginTime=cheerlib.TimeGetNow()
	this.LastTime=cheerlib.TimeGetNow()

	this.TokenId=cheerlib.EncryptNewId()
	this.TokenId=cheerlib.EncryptMd5(fmt.Sprintf("%s-%s-%s",this.TenantId,this.LoginAddr,this.TokenId))

}