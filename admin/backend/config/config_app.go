package config

import "errors"

type ConfigApp struct {

	ServerAddr string `yaml:"server_addr"`
	ServerPort int `yaml:"server_port"`

	DbUri      string `yaml:"db_uri"`

}

func (this *ConfigApp)Check() error  {

	var xError error

	if len(this.DbUri)<1{
		xError=errors.New("invalid config node=[db_uri]")
		return xError
	}

	return xError

}