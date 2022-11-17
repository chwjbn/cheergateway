package cheerlib

import (
"fmt"
"io"
"log"
"os"
"path"
"path/filepath"
"runtime"
"strings"
"sync"
"time"
)

type Log struct {
	mLocker *sync.RWMutex
	mLogDate string

	mInfoLogger *log.Logger
	mErrorLogger *log.Logger

	mInfoLogFile *os.File
	mErrorLogFile *os.File
}

func CreateLog() *Log  {

	pThis:=new(Log)
	pThis.mLocker=new(sync.RWMutex)

	pThis.CheckLog()

	return pThis

}

var mLog=CreateLog()

func LogInfo(logContent string)  {

	_,file,line,_ := runtime.Caller(1)
	file=filepath.Base(file)

	logContent=fmt.Sprintf("%s:%d ",file,line)+logContent

	mLog.CheckLog()
	mLog.mInfoLogger.Println(logContent)

}

func LogInfoSkipStack(logContent string,skipStack int)  {

	_,file,line,_ := runtime.Caller(skipStack)
	file=filepath.Base(file)

	logContent=fmt.Sprintf("%s:%d ",file,line)+logContent

	mLog.CheckLog()
	mLog.mInfoLogger.Println(logContent)
}

func LogError(logContent string)  {

	_,file,line,_ := runtime.Caller(1)
	file=filepath.Base(file)

	logContent=fmt.Sprintf("%s:%d ",file,line)+logContent

	mLog.CheckLog()
	mLog.mErrorLogger.Println(logContent)

}

func LogErrorSkipStack(logContent string,skipStack int)  {

	_,file,line,_ := runtime.Caller(skipStack)
	file=filepath.Base(file)

	logContent=fmt.Sprintf("%s:%d ",file,line)+logContent

	mLog.CheckLog()
	mLog.mErrorLogger.Println(logContent)

}

func LogTag(logContent string)  {

	logContent="===================="+logContent+"===================="

	_,file,line,_ := runtime.Caller(1)
	file=filepath.Base(file)

	logContent=fmt.Sprintf("%s:%d ",file,line)+logContent

	mLog.CheckLog()
	mLog.mInfoLogger.Println(logContent)
}


func (this *Log)CheckLog()  {

	lastLogDate:=this.readLogDate()
	logDate:=time.Now().Format("20060102")

	if strings.EqualFold(lastLogDate,logDate){
		return
	}

	this.setLogDate(logDate)
}

func (this *Log)readLogDate() string  {

	this.mLocker.RLock()
	defer this.mLocker.RUnlock()

	return this.mLogDate

}

func (this *Log)setLogDate(logDate string)  {

	this.mLocker.Lock()
	defer this.mLocker.Unlock()

	logFileDir:=path.Join(ApplicationBaseDirectory(),"log",path.Base(ApplicationFileName()))
	if !DirectoryExists(logFileDir){
		DirectoryCreateDirectory(logFileDir)
	}


	if this.mInfoLogFile!=nil{
		this.mInfoLogFile.Close()
		this.mInfoLogFile=nil
	}

	if this.mErrorLogFile!=nil{
		this.mErrorLogFile.Close()
		this.mErrorLogFile=nil
	}


	var mInfoLogFileErr error
	this.mInfoLogFile,mInfoLogFileErr=os.OpenFile(path.Join(logFileDir,fmt.Sprintf("info-%s.log",logDate)),os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if mInfoLogFileErr!=nil{
		os.Stderr.WriteString(fmt.Sprintf("cheerlib.Log.setLogDate.mInfoLogFileErr=%s\n",mInfoLogFileErr.Error()))
		return
	}

	var mErrorLogFileErr error
	this.mErrorLogFile,mErrorLogFileErr=os.OpenFile(path.Join(logFileDir,fmt.Sprintf("error-%s.log",logDate)),os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if mErrorLogFileErr!=nil{
		os.Stderr.WriteString(fmt.Sprintf("cheerlib.Log.setLogDate.mErrorLogFileErr=%s\n",mErrorLogFileErr.Error()))
		return
	}

	if this.mInfoLogger!=nil{
		this.mInfoLogger.SetOutput(io.MultiWriter(os.Stdout,this.mInfoLogFile))
	}else{
		this.mInfoLogger=log.New(io.MultiWriter(os.Stdout,this.mInfoLogFile),"[Info]",log.Ldate | log.Ltime)
	}

	if this.mErrorLogger!=nil{
		this.mErrorLogger.SetOutput(io.MultiWriter(os.Stderr,this.mErrorLogFile))
	}else{
		this.mErrorLogger=log.New(io.MultiWriter(os.Stderr,this.mErrorLogFile),"[Error]",log.Ldate | log.Ltime)
	}

	this.mLogDate=logDate
}


