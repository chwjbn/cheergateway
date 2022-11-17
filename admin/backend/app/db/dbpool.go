package db

type DbPool struct {
	mDbUri string
	mMongodbSvc *MongodbSvc

	DbTenantSvc *DbTenantSvc
	DbAppSvc *DbAppSvc
}

func NewDbPool(dbUri string) (error,*DbPool)  {

	var xError error=nil
	pThis:=new(DbPool)

	pThis.mDbUri=dbUri

	xError,pThis.mMongodbSvc=NewMongodbSvc(pThis.mDbUri)

	if xError!=nil{
		return xError,pThis
	}

	//租户服务
	pThis.DbTenantSvc=new(DbTenantSvc)
	pThis.DbTenantSvc.InitDbModel(pThis,"db_cheergateway")

	//应用数据
	pThis.DbAppSvc=new(DbAppSvc)
	pThis.DbAppSvc.InitDbModel(pThis,"db_cheergateway")


	return xError,pThis

}

func (this *DbPool) GetMongodbSvc() *MongodbSvc  {
	return this.mMongodbSvc
}



