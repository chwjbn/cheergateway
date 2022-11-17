package dbmodel

import (
	"fmt"
	"math/rand"
	"cheeradmin/cheerlib"
	"strings"
	"time"
)

type TenantCaptcha struct {

	CaptchaId string `bson:"captcha_id" json:"captcha_id"`
	CaptchaType string `bson:"captcha_type" json:"captcha_type"`
	CaptchaTarget string `bson:"captcha_target" json:"captcha_target"`
	CaptchaCode string `bson:"captcha_code" json:"captcha_code"`
	CreateTime string `bson:"create_time" json:"create_time"`
	Status string `bson:"status" json:"status"`
}

func (this *TenantCaptcha)Init()  {


	this.CaptchaCode=this.GenCode(6)

	this.CaptchaId=cheerlib.EncryptMd5(fmt.Sprintf("%s-%s-%s",this.CaptchaType,this.CaptchaTarget,this.CaptchaCode))
	this.CaptchaId=cheerlib.EncryptMd5(this.CaptchaId+cheerlib.EncryptNewId())

	this.CreateTime=cheerlib.TimeGetNow()
	this.Status="active"

}


func (this *TenantCaptcha)GenCode(width int) string  {

	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ rand.Intn(r) ])
	}

	return sb.String()
}