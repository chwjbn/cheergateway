package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"cheeradmin/cheerlib"
	"sync"
	"time"
)

type MongodbSvc struct {
	mDbUri string
	mDbClient *mongo.Client
	mLocker *sync.RWMutex
}

func NewMongodbSvc(dbUri string) (error,*MongodbSvc)  {

	var xError error=nil

	pThis:=new(MongodbSvc)
	pThis.mDbUri=dbUri

	pThis.mLocker=new(sync.RWMutex)

	pThis.checkClient()

	return xError,pThis

}

func (this *MongodbSvc)GetDbHandle(dbName string) *mongo.Database  {

	this.checkClient()

	return this.mDbClient.Database(dbName)

}

func (this *MongodbSvc)checkClient()  {

	this.mLocker.Lock()
	defer this.mLocker.Unlock()

	var xError error

	if this.mDbClient!=nil{

		xError=this.mDbClient.Ping(context.TODO(),nil)

		if xError!=nil{
			cheerlib.LogError("MongodbSvc.CheckClient mDbClient.Ping Error="+xError.Error())
			this.mDbClient.Disconnect(context.TODO())
			this.mDbClient=nil
		}

		// cheerlib.LogInfo("MongodbSvc.CheckClient mDbClient.Ping Success")
	}

	if this.mDbClient!=nil{
		return
	}


	clientOptions := options.Client().ApplyURI(this.mDbUri).SetConnectTimeout(10*time.Second).SetSocketTimeout(10*time.Second)

	cheerlib.LogInfo(fmt.Sprintf("MongodbSvc begin connect to [%s]",this.mDbUri))

	for{

		this.mDbClient,xError=mongo.Connect(context.TODO(), clientOptions)
		if xError==nil{
			break
		}

		cheerlib.LogError("MongodbSvc.CheckClient mongo.Connect Error="+xError.Error())

		time.Sleep(3*time.Second)
	}


	cheerlib.LogInfo(fmt.Sprintf("MongodbSvc successfully connected to [%s]",this.mDbUri))


}

