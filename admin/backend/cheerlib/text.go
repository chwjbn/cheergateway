package cheerlib

import (
	"encoding/json"
	"golang.org/x/text/encoding"
	"net"
	"strconv"
	"strings"
)

func TextStructFromJson(data interface{},dataJson string) error  {

	var xErr error=nil
	xErr=json.Unmarshal([]byte(dataJson),data)
	return xErr

}

func TextStructToJson(data interface{}) string  {

	xData:="{}"

	jonData,jsonErr:= json.Marshal(data)

	if jsonErr!=nil{
		return xData
	}

	xData=string(jonData)

	return xData
}

func TextGetString(data []byte,encoding encoding.Encoding) string  {

	sData:=""

	dataBuffer,xErr:=encoding.NewDecoder().Bytes(data)

	if xErr!=nil{
		return sData
	}

	sData=string(dataBuffer)

	return sData

}

func TextGetMapColumn(dataMapList []map[string]interface{},colName string) []interface{}  {

	dataList:=[]interface{}{}

	for _,dataMapItem:=range dataMapList{

		xDataItem,bFind:=dataMapItem[colName]
		if !bFind{
			continue
		}

		dataList=append(dataList,xDataItem)
	}

	return dataList
}
func TextIsNodeAddr(data string) bool  {

	bRet:=false

	if !strings.Contains(data,":"){
		return bRet
	}

	xStrArr:=strings.Split(data,":")
	if len(xStrArr)!=2{
		return bRet
	}

	xAddr:=net.ParseIP(xStrArr[0])
	if xAddr==nil{
		return bRet
	}

	xPort,_:=strconv.Atoi(xStrArr[1])
	if xPort<1||xPort>65535{
		return bRet
	}

	bRet=true

	return bRet
}

func TextIsNodeAddrList(data string) bool  {

	bRet:=false

	if !strings.Contains(data,":"){
		return bRet
	}

	if !strings.Contains(data,"|"){
		bRet=TextIsNodeAddr(data)
		return bRet
	}

	xDataArr:=strings.Split(data,"|")

	for _,xDataItem:=range xDataArr{
		if !TextIsNodeAddr(xDataItem){
			return bRet
		}
	}

	bRet=true

	return bRet

}
