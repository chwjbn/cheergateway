package main

import (
	"errors"
	"fmt"
	"cheeradmin/app"
	"cheeradmin/cheerlib"
	"cheeradmin/common"
	"cheeradmin/config"
	"os"
	"path"
	"runtime"
	"strings"
)

func AppWork() error {

	var xError error

	defer func() {

		if xError!=nil{
			cheerlib.LogError(fmt.Sprintf("AppWork Error=[%s]",xError.Error()))
		}

	}()

	configFilePath:=path.Join(cheerlib.ApplicationBaseDirectory(),"config","app.yml")
	if !cheerlib.FileExists(configFilePath){
		xError=errors.New("Lost Config File:"+configFilePath)
		return xError
	}

	var cfg config.ConfigApp

	xError=config.ReadConfigFile(configFilePath,&cfg)
	if xError!=nil{
		return xError
	}

	xError=cfg.Check()
	if xError!=nil{
		return xError
	}

	//读取环境变量
	xEnvDbUri:=os.Getenv("db_uri")
	if len(xEnvDbUri)>0{
		cfg.DbUri=xEnvDbUri
	}

	xError=app.RunGateway(&cfg)
	if xError!=nil{
		return xError
	}

	return xError
}

func RunApp()  {

	runtime.GOMAXPROCS(runtime.NumCPU())

	xServiceMgr,xServiceMgrErr:=common.CreateServiceMgr("cheeradmin","cheeradmin App Service","cheeradmin App Service",AppWork)
	if xServiceMgrErr!=nil{
		cheerlib.LogError("xServiceMgrErr="+xServiceMgrErr.Error())
	}


	xRunArgs:=os.Args

	if len(xRunArgs)>1{
		xRunArg:=xRunArgs[1]
		if strings.Contains(xRunArg,"install"){
			xServiceMgr.Install()
			return
		}

		if strings.Contains(xRunArg,"remove"){
			xServiceMgr.StopService()
			xServiceMgr.Uninstall()
			return
		}
	}


	xServiceMgr.RunService()

}

func main() {

	cheerlib.LogTag("cheeradmin App Begin")
	RunApp()
	cheerlib.LogTag("cheeradmin App End")

}
