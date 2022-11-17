package config

import (
	"cheeradmin/cheerlib"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Config interface {
	Check() error
}

var configFileName string

func ParseConfigData(configData []byte,cfg Config) error  {
	var xError error=nil
	xError=yaml.Unmarshal(configData,cfg)
	return xError

}

func ReadConfigFile(filePath string,cfg Config) error  {

	var xError error=nil
	var xFilePath =""

	var xConfigData []byte

	xFilePath,xError=filepath.Abs(filePath)
	if xError!=nil{
		return xError
	}

	xConfigData,xError=ioutil.ReadFile(xFilePath)

	if xError!=nil{
		return xError
	}

	configFileName=xFilePath

	xError=ParseConfigData(xConfigData,cfg)

	if xError==nil{
		cheerlib.LogInfo(fmt.Sprintf("ReadConfigFile From [%s] Values(%v)",xFilePath,cfg))
	}

	return xError

}

func SaveConfigFile(cfg Config) error {

	xConfigData, xError := yaml.Marshal(cfg)
	if xError != nil {
		return xError
	}

	xError = ioutil.WriteFile(configFileName, xConfigData, 0755)
	if xError != nil {
		return xError
	}


	cheerlib.LogInfo(fmt.Sprintf("SaveConfigFile To [%s] Values(%s)",configFileName,cfg))

	return nil
}

