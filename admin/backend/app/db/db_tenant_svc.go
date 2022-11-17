package db

import (
	"cheeradmin/app/db/dbmodel"
)

type DbTenantSvc struct {
	DbModelSvc
}

func (this *DbTenantSvc)CreateTenantCaptcha(data dbmodel.TenantCaptcha) error  {

	var xError error=nil

	xTableName:="tenant_captcha"


	xError=this.AddData(xTableName,data)
	if xError!=nil{
		return xError
	}

	this.MakeSureIndex(xTableName,[]string{"captcha_id"})
	this.MakeSureIndex(xTableName,[]string{"captcha_type"})
	this.MakeSureIndex(xTableName,[]string{"captcha_target"})

	this.MakeSureIndex(xTableName,[]string{"create_time"})
	this.MakeSureIndex(xTableName,[]string{"status"})


	return xError
}

func (this *DbTenantSvc)GetTenantCaptcha(captchaId string)  dbmodel.TenantCaptcha {

	var xData dbmodel.TenantCaptcha

	xTableName:="tenant_captcha"

	xWhere:=make(map[string]interface{})
	xWhere["captcha_id"]=captchaId

	xError:=this.GetData(xTableName,xWhere,&xData)
	if xError!=nil{
		xData=dbmodel.TenantCaptcha{}
	}

	return xData
}

func (this *DbTenantSvc)CreateAccessSession(data dbmodel.TenantAccessSession) error  {

	var xError error=nil

	xTableName:="tenant_token"


	xError=this.AddData(xTableName,data)
	if xError!=nil{
		return xError
	}

	this.MakeSureIndex(xTableName,[]string{"token_id"})
	this.MakeSureIndex(xTableName,[]string{"tenant_id"})
	this.MakeSureIndex(xTableName,[]string{"last_time"})


	return xError

}

func (this *DbTenantSvc)GetAccessSession(tokenId string) dbmodel.TenantAccessSession {

	var xData dbmodel.TenantAccessSession

	xTableName:="tenant_token"

	xWhere:=make(map[string]interface{})
	xWhere["token_id"]=tokenId

	xError:=this.GetData(xTableName,xWhere,&xData)
	if xError!=nil{
		xData=dbmodel.TenantAccessSession{}
	}

	return xData

}

func (this *DbTenantSvc)UpdateAccessSession(data dbmodel.TenantAccessSession) error  {

	var xError error=nil

	xTableName:="tenant_token"

	xWhere:=make(map[string]interface{})
	xWhere["token_id"]=data.TokenId

	xError=this.SaveData(xTableName,xWhere,data)

	return xError

}


func (this *DbTenantSvc)GetTenantInfo(tenantId string)dbmodel.TenantInfo  {

	xData:=dbmodel.TenantInfo{}

	xTableName:="tenant_info"

	xWhere:=make(map[string]interface{})
	xWhere["tenant_id"]=tenantId

	xError:=this.GetData(xTableName,xWhere,&xData)
	if xError!=nil{
		xData=dbmodel.TenantInfo{}
	}


	return xData

}

func (this *DbTenantSvc)GetTenantInfoByUserName(userName string)dbmodel.TenantInfo  {

	xData:=dbmodel.TenantInfo{}


	xTableName:="tenant_info"

	xWhere:=make(map[string]interface{})
	xWhere["user_name"]=userName

	xError:=this.GetData(xTableName,xWhere,&xData)
	if xError!=nil{
		xData=dbmodel.TenantInfo{}
	}


	return xData

}

func (this *DbTenantSvc)GetTenantInfoByEmail(email string)dbmodel.TenantInfo  {

	xData:=dbmodel.TenantInfo{}


	xTableName:="tenant_info"

	xWhere:=make(map[string]interface{})
	xWhere["email"]=email

	xError:=this.GetData(xTableName,xWhere,&xData)
	if xError!=nil{
		xData=dbmodel.TenantInfo{}
	}


	return xData

}

func (this *DbTenantSvc)CreateTenantInfo(data dbmodel.TenantInfo) error {

	var xError error=nil

	xTableName:="tenant_info"

	xError=this.AddData(xTableName,data)
    if xError!=nil{
    	return xError
	}

	this.MakeSureIndex(xTableName,[]string{"tenant_id"})
	this.MakeSureIndex(xTableName,[]string{"user_name"})
	this.MakeSureIndex(xTableName,[]string{"email"})
	this.MakeSureIndex(xTableName,[]string{"create_time"})
	this.MakeSureIndex(xTableName,[]string{"status"})

	return xError

}

func (this *DbTenantSvc)SaveTenantInfo(data dbmodel.TenantInfo) error {

	var xError error=nil

	xTableName:="tenant_info"

	xWhere:=make(map[string]interface{})
	xWhere["tenant_id"]=data.TenantId

	xError=this.SaveData(xTableName,xWhere,data)


	return xError

}