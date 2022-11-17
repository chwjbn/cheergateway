package cheerlib

import (
	"os"
	"path/filepath"
)

func ApplicationBaseDirectory() string {

	sRet:=""

	xFilePath,xFilePathErr:=filepath.Abs(os.Args[0])
	if xFilePathErr!=nil{
		return sRet
	}

	sRet=filepath.Dir(xFilePath)

	return sRet

}

func ApplicationFileName() string  {

	sRet:=""

	xFilePath,xFilePathErr:=filepath.Abs(os.Args[0])
	if xFilePathErr!=nil{
		return sRet
	}

	sRet=filepath.Base(xFilePath)

	return sRet

}

func ApplicationFullPath() string {

	sRet:=""

	xFilePath,xFilePathErr:=filepath.Abs(os.Args[0])
	if xFilePathErr!=nil{
		return sRet
	}

	sRet=xFilePath

	return sRet

}
