package app

import (
    "bufio"
    "errors"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/klauspost/compress/zip"
    "io/ioutil"
    "cheeradmin/app/ctl"
    "cheeradmin/app/db"
    "cheeradmin/cheerlib"
    "cheeradmin/config"
    "net/http"
    "path"
    "strings"
)

type CheerAdminApp struct {
    mConfig *config.ConfigApp
    mDbPool *db.DbPool

    mCmsCtl *ctl.CmsCtl
}

func RunGateway(cfg *config.ConfigApp) error  {

    var xError error

    xApp:=new(CheerAdminApp)
    xApp.mConfig=cfg

    xApp.mCmsCtl=new(ctl.CmsCtl)
    xError=xApp.mCmsCtl.Init(xApp.mConfig)

    if xError!=nil{
        return xError
    }


    xError=xApp.Run()

    return xError

}

func (this *CheerAdminApp)Run() error  {

    var xError error

    gin.SetMode(gin.DebugMode)

    xRouter := gin.Default()
    xRouter.SetTrustedProxies([]string{"127.0.0.1"})

    xRouter.Any("/*xpath",this.onHttpRequest)

    xServerHostPort:=fmt.Sprintf("%s:%d",this.mConfig.ServerAddr,this.mConfig.ServerPort)

    xError=xRouter.Run(xServerHostPort)

    return xError
}

func (this *CheerAdminApp) onHttpRequest(ctx *gin.Context)  {

    xRouterMap:=make(map[string]gin.HandlerFunc)

    xRouterMap["/xapi/cms/tenant-init"]=this.mCmsCtl.CtlTenantInit

    xRouterMap["/xapi/cms/check-code-image"]=this.mCmsCtl.CtlCheckCodeImage
    xRouterMap["/xapi/cms/tenant-login"]=this.mCmsCtl.CtlTenantLogin
    xRouterMap["/xapi/cms/tenant-current"]=this.mCmsCtl.CtlTenantCurrent
    xRouterMap["/xapi/cms/tenant-update-info"]=this.mCmsCtl.CtlTenantUpdateInfo
    xRouterMap["/xapi/cms/tenant-update-password"]=this.mCmsCtl.CtlTenantUpdatePassword

    xRouterMap["/xapi/cms/app-gateway-config-page"]=this.mCmsCtl.CtlAppGatewayConfigPage
    xRouterMap["/xapi/cms/app-gateway-config-add"]=this.mCmsCtl.CtlAppGatewayConfigAdd
    xRouterMap["/xapi/cms/app-gateway-config-get"]=this.mCmsCtl.CtlAppGatewayConfigGet
    xRouterMap["/xapi/cms/app-gateway-config-save"]=this.mCmsCtl.CtlAppGatewayConfigSave
    xRouterMap["/xapi/cms/app-gateway-config-remove"]=this.mCmsCtl.CtlAppGatewayConfigRemove

    xRouterMap["/xapi/cms/app-gateway-config-pub"]=this.mCmsCtl.CtlAppGatewayConfigPublish

    xRouterMap["/xapi/cms/app-gateway-config-map-data"]=this.mCmsCtl.CtlAppGatewayConfigMapData

    xRouterMap["/xapi/cms/app-gateway-backend-page"]=this.mCmsCtl.CtlAppGatewayBackendPage
    xRouterMap["/xapi/cms/app-gateway-backend-add"]=this.mCmsCtl.CtlAppGatewayBackendAdd
    xRouterMap["/xapi/cms/app-gateway-backend-get"]=this.mCmsCtl.CtlAppGatewayBackendGet
    xRouterMap["/xapi/cms/app-gateway-backend-save"]=this.mCmsCtl.CtlAppGatewayBackendSave
    xRouterMap["/xapi/cms/app-gateway-backend-remove"]=this.mCmsCtl.CtlAppGatewayBackendRemove

    xRouterMap["/xapi/cms/app-gateway-backend-map-data"]=this.mCmsCtl.CtlAppGatewayBackendMapData

    xRouterMap["/xapi/cms/app-gateway-site-page"]=this.mCmsCtl.CtlAppGatewaySitePage
    xRouterMap["/xapi/cms/app-gateway-site-add"]=this.mCmsCtl.CtlAppGatewaySiteAdd
    xRouterMap["/xapi/cms/app-gateway-site-get"]=this.mCmsCtl.CtlAppGatewaySiteGet
    xRouterMap["/xapi/cms/app-gateway-site-save"]=this.mCmsCtl.CtlAppGatewaySiteSave
    xRouterMap["/xapi/cms/app-gateway-site-remove"]=this.mCmsCtl.CtlAppGatewaySiteRemove

    xRouterMap["/xapi/cms/app-gateway-site-map-data"]=this.mCmsCtl.CtlAppGatewaySiteMapData


    xRouterMap["/xapi/cms/app-gateway-rule-page"]=this.mCmsCtl.CtlAppGatewayRulePage
    xRouterMap["/xapi/cms/app-gateway-rule-add"]=this.mCmsCtl.CtlAppGatewayRuleAdd
    xRouterMap["/xapi/cms/app-gateway-rule-get"]=this.mCmsCtl.CtlAppGatewayRuleGet
    xRouterMap["/xapi/cms/app-gateway-rule-save"]=this.mCmsCtl.CtlAppGatewayRuleSave
    xRouterMap["/xapi/cms/app-gateway-rule-remove"]=this.mCmsCtl.CtlAppGatewayRuleRemove

    //先匹配api接口路由
    xReqPath:=ctx.Request.URL.Path
    bIsMatch:=false

    for xRKey,xRVal:=range xRouterMap{

        if strings.EqualFold(xReqPath,xRKey){
            bIsMatch=true
            xRVal(ctx)
            break
        }
    }

    if bIsMatch{
        return
    }

    //读取资源文件
    xFileErr,xFileContent:=this.readStaticFile(xReqPath)
    if xFileErr!=nil{
        ctx.String(http.StatusNotFound,"Resource Error:%s",xFileErr.Error())
        return
    }

    //获取资源文件的MIME
    xContentType:=this.getContentType(xReqPath)

    ctx.Header("Content-Type",xContentType)
    ctx.Writer.Write(xFileContent)

}



func (this *CheerAdminApp)getContentType(urlPath string) string  {

    sData:="text/html; charset=utf-8"

    xFilePath:=path.Join(cheerlib.ApplicationBaseDirectory(),"data","mime.data")
    if !cheerlib.FileExists(xFilePath){
        return sData
    }

    dataContent:=cheerlib.FileReadAllText(xFilePath)

    xScanner := bufio.NewScanner(strings.NewReader(dataContent))


    for xScanner.Scan(){

        xLine:=xScanner.Text()
        xIdex:=strings.Index(xLine,"|")
        if xIdex<1{
            continue
        }

        xDataKey:="."+xLine[0:xIdex]
        xDataVal:=xLine[xIdex+1:]

        if strings.HasSuffix(urlPath,xDataKey){
            sData=fmt.Sprintf("%s; charset=utf-8",xDataVal)
        }

    }

    return sData

}

func (this *CheerAdminApp)readStaticFile(urlPath string)(error,[]byte)  {

    var xError error=nil
    xFileContent:=[]byte{}

    xZipFilePath:=path.Join(cheerlib.ApplicationBaseDirectory(),"data","webui.zip")

    if !cheerlib.FileExists(xZipFilePath){
        xError=errors.New("webui Not Found.")
        return xError,xFileContent
    }

    xZipFile,xZipFileErr:=zip.OpenReader(xZipFilePath)
    if xZipFileErr!=nil{
        xError=xZipFileErr
        return xError,xFileContent
    }

    defer xZipFile.Close()


    xMatchFile:="dist"+urlPath
    xMatchFileWithIndex:=xMatchFile
    if strings.HasSuffix(xMatchFileWithIndex,"/"){
        xMatchFileWithIndex=xMatchFileWithIndex+"index.html"
    }else {
        xMatchFileWithIndex=xMatchFileWithIndex+"/index.html"
    }



    for _,xZipFileItem:=range xZipFile.File{

        if xZipFileItem.FileInfo().IsDir(){
            continue
        }

        xFName:=xZipFileItem.Name
        if strings.EqualFold(xFName,xMatchFile)||strings.EqualFold(xFName,xMatchFileWithIndex){

            xZipItemFile,xZipItemFileErr:=xZipFileItem.Open()
            if xZipItemFileErr!=nil{
                xError=xZipItemFileErr
                return xError,xFileContent
            }

            defer xZipItemFile.Close()

            xDataBuffer,xDataBufferErr:=ioutil.ReadAll(xZipItemFile)
            if xDataBufferErr!=nil{
                xError=xDataBufferErr
                return xError,xFileContent
            }

            xFileContent=xDataBuffer
            return xError,xFileContent
        }

    }

    xError=errors.New(fmt.Sprintf("%s Not Found.",urlPath))
    return xError,xFileContent

}
