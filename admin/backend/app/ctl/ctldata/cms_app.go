package ctldata

import (
	"errors"
)

type BasePageRequest struct {
	PageNo int `json:"current"`
	PageSize int `json:"pageSize"`
}

type BaseGetInfoRequest struct {
	DataId string `json:"data_id"`
}

type DataMapNode struct {
	DataId string `json:"data_id"`
	DataName string `json:"data_name"`
}


func (this *BaseGetInfoRequest)Check() error  {

	var xError error=nil

	if len(this.DataId)<1{
		xError=errors.New("请求数据ID缺失!")
		return xError
	}

	return xError
}
