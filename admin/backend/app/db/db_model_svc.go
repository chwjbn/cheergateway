package db

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"cheeradmin/app/db/dbmodel"
	"cheeradmin/cheerlib"
	"strings"
)

type DbModelSvc struct {
	mDbPool *DbPool
	mDbName string
}

func (this *DbModelSvc)InitDbModel(dbPool *DbPool,dbName string)  {

	this.mDbPool=dbPool
	this.mDbName=dbName

}

func (this *DbModelSvc)GetDb() *mongo.Database  {
	return this.mDbPool.GetMongodbSvc().GetDbHandle(this.mDbName)
}

func (this *DbModelSvc)GetTable(tableName string)*mongo.Collection  {

	xDb:=this.GetDb()
	return xDb.Collection(tableName)
}

func (this *DbModelSvc)AddData(tableName string,data interface{}) error  {

	var xError error

	xTable:=this.GetTable(tableName)

	_,xError=xTable.InsertOne(context.TODO(),data)

	return xError

}

func (this *DbModelSvc)AddDataList(tableName string,dataList[]interface{}) error  {

	var xError error

	xTable:=this.GetTable(tableName)

	_,xError=xTable.InsertMany(context.TODO(),dataList)

	return xError
}


func (this *DbModelSvc)SaveData(tableName string,whereMap map[string]interface{},data interface{}) error  {

	var xError error


	xWhere:=bson.M{}

	for dataK,dataV:=range whereMap{
		xWhere[dataK]=dataV
	}

	xTable:=this.GetTable(tableName)

	_,xError=xTable.UpdateMany(context.TODO(),xWhere,bson.M{"$set":data})

	return xError

}


func (this *DbModelSvc)DeleteData(tableName string,whereMap map[string]interface{}) error  {

	var xError error


	xWhere:=bson.M{}

	for dataK,dataV:=range whereMap{
		xWhere[dataK]=dataV
	}

	xTable:=this.GetTable(tableName)

	_,xError=xTable.DeleteMany(context.TODO(),xWhere)

	return xError

}

func (this *DbModelSvc)GetData(tableName string,whereMap map[string]interface{},data interface{}) error  {

	var xError error=nil

	xTable:=this.GetTable(tableName)

	xWhere:=bson.M{}

	for dataK,dataV:=range whereMap{
		xWhere[dataK]=dataV
	}

	xFindOneRet:=xTable.FindOne(context.TODO(),xWhere)

	xError=xFindOneRet.Err()
	if xError!=nil{
		return xError
	}

	xError=xFindOneRet.Decode(data)
	if xError!=nil{
		return xError
	}

	return xError
}


func (this *DbModelSvc)GetDataWhereAndOrder(tableName string,whereMap map[string]interface{},sortMap map[string]interface{},data interface{}) error  {

	var xError error=nil

	xTable:=this.GetTable(tableName)

	xWhere:=bson.M{}

	for dataK,dataV:=range whereMap{
		xWhere[dataK]=dataV
	}

	xFindOneOpt:=options.FindOne()

	xSort:=bson.M{}
	for dataK,dataV:=range sortMap{
		xSort[dataK]=dataV
	}

	xFindOneOpt.SetSort(xSort)

	xFindOneRet:=xTable.FindOne(context.TODO(),xWhere,xFindOneOpt)

	xError=xFindOneRet.Err()
	if xError!=nil{
		return xError
	}

	xError=xFindOneRet.Decode(data)
	if xError!=nil{
		return xError
	}

	return xError
}

func (this *DbModelSvc)GetDataCount(tableName string,whereMap map[string]interface{}) int64  {

	var xCount int64=0

	xTable:=this.GetTable(tableName)

	xWhere:=bson.M{}

	for dataK,dataV:=range whereMap{
		xWhere[dataK]=dataV
	}

	xCount,_=xTable.CountDocuments(context.TODO(),xWhere)

	return xCount

}


func (this *DbModelSvc)GetDataList(tableName string,whereMap map[string]interface{},sortMap map[string]interface{},fromIndex int64,count int64) (error,[]map[string]interface{})  {

	var xError error=nil

	xDataList:=[]map[string]interface{}{}

	xTable:=this.GetTable(tableName)

	xWhere:=bson.M{}

	for dataK,dataV:=range whereMap{
		xWhere[dataK]=dataV
	}


	xFindOpt:=options.Find()

	xSort:=bson.M{}
	for dataK,dataV:=range sortMap{
		xSort[dataK]=dataV
	}

	xFindOpt.SetSort(xSort)

	if fromIndex>0{
		xFindOpt.SetSkip(fromIndex)
	}

	if count>0{
		xFindOpt.SetLimit(count)
	}


	var xFindRs *mongo.Cursor

	xFindRs,xError=xTable.Find(context.TODO(),xWhere,xFindOpt)
	if xError!=nil{
		return xError,xDataList
	}

	defer func() {
		xFindRs.Close(context.TODO())
	}()

	for xFindRs.Next(context.TODO()){

		xData:=make(map[string]interface{})
		if xFindRs.Decode(&xData)!=nil{
			break
		}

		xDataList=append(xDataList,xData)
	}

	return xError,xDataList
}

func (this *DbModelSvc)GetDataPageList(tableName string,whereMap map[string]interface{},sortMap map[string]interface{},pageNo int,pageSize int) dbmodel.PageData {

	xPageData:=dbmodel.PageData{}

	xPageData.TotalCount=this.GetDataCount(tableName,whereMap)
	xPageData.PageNo=int64(pageNo)
	xPageData.PageSize=int64(pageSize)
	xPageData.Calc()

	xFromIndex:=(xPageData.PageNo-1)*xPageData.PageSize
	xLimitCount:=xPageData.PageSize

	xError,xDataList:=this.GetDataList(tableName,whereMap,sortMap,xFromIndex,xLimitCount)
	if xError!=nil{
		return xPageData
	}

	for _,dataItem:=range xDataList{
        xPageData.DataList=append(xPageData.DataList,dataItem)
	}

	return xPageData
}

func (this *DbModelSvc)MakeSureIndex(tableName string,fieldList[]string)  {
	this.MakeSureTableIndex(tableName,fieldList,false)
}

func (this *DbModelSvc)MakeSureIndexUnique(tableName string,fieldList[]string)  {
	this.MakeSureTableIndex(tableName,fieldList,true)
}

func (this *DbModelSvc)MakeSureTableIndex(tableName string,fieldList[]string,uniqueFlag bool)  {

	if len(fieldList)<1{
		return
	}

	if this.checkInit()!=nil{
		return
	}


	xIndexName:=fmt.Sprintf("idex_%s",strings.Join(fieldList,"_"))

	xTable := this.GetTable(tableName)

	xCur,xErr:=xTable.Indexes().List(context.TODO())
	if xErr!=nil{
		cheerlib.LogError("DbModelSvc.MakeSureIndex Error:"+xErr.Error())
		return
	}

	defer func() {
		xCur.Close(context.TODO())
	}()

	var xIndexResult []bson.M
	xErr=xCur.All(context.TODO(),&xIndexResult)
	if xErr!=nil{
		cheerlib.LogError("DbModelSvc.MakeSureIndex Error:"+xErr.Error())
		return
	}

	xIsHave:=false
	for _,xIndex:=range xIndexResult{

		xName:=xIndex["name"].(string)
		if xName==xIndexName{
			xIsHave=true
			break
		}

	}

	if xIsHave{
		return
	}


	xIndexOpt:=new(options.IndexOptions)
	xIndexOpt.SetName(xIndexName)
	xIndexOpt.SetBackground(true)
	xIndexOpt.SetUnique(uniqueFlag)

	xIndexData:=mongo.IndexModel{Options: xIndexOpt,Keys: bsonx.Doc{}}
	for _,fItem:=range fieldList{
		xIndexData.Keys=xIndexData.Keys.(bsonx.Doc).Append(fItem,bsonx.Int32(1))
	}


	xCreateRet,xCreateErr:=xTable.Indexes().CreateOne(context.Background(),xIndexData)
	if xCreateErr!=nil{
		cheerlib.LogError(fmt.Sprintf("DbModelSvc.MakeSureIndex tableName=%s.%s xCreateErr=%s",this.mDbName,tableName,xCreateErr.Error()))
		return
	}

	cheerlib.LogInfo(fmt.Sprintf("DbModelSvc.MakeSureIndex tableName=%s.%s xCreateRet=%s",this.mDbName,tableName,xCreateRet))

}

func (this *DbModelSvc) checkInit() error  {

	var xError error=nil

	if len(this.mDbName)<1||this.mDbPool==nil{
		xError=errors.New("DbModelSvc not initialized.")
		cheerlib.LogError("DbModelSvc checkInit Error:"+xError.Error())
		return xError
	}

	return xError

}




