//
//  main.go
//  golang
//
//  Created by zhangzhe on 2014-1-25
//

package mongo

import (
	"errors"
	"labix.org/v2/mgo"
)

/*******************************************
/*描述：
/*    DataBase是数据库实例的描述类，用于存取数据
/*  仓内的集合.
/*
/*****************************************/
type DataBase struct {
	DBName       string                       //公司名称
	DBInstance   *mgo.Database                //该公司对应的数据库实例
	DBCollection map[string](*mgo.Collection) //Key:集合名称 Value:集合对象
}

func (this *DataBase) Init(dbName string) {
	this.DBName = dbName
	this.DBInstance = nil
	this.DBCollection = make(map[string](*mgo.Collection))

}

/******************************************
*           返回一个集合
*******************************************/
func (this *DataBase) Collection(collectionName string) (result *mgo.Collection) {
	if this.DBInstance == nil {
		return nil
	}
	_, ok := this.DBCollection[collectionName]

	if !ok {
		this.newCollection(collectionName)
	}

	result = this.DBCollection[collectionName]

	return result
}

/******************************************
*           新增一个集合
*******************************************/
func (this *DataBase) newCollection(collectionName string) {
	if this.DBInstance != nil {
		collection := this.DBInstance.C(collectionName)
		this.DBCollection[collectionName] = collection
	}

}

/*******************************************
/*描述：
/*   MongoDBControl是数据库连接描述类，用于存取数据
/*  库.
/*
/*****************************************/
type Handler struct {
	DBsession *mgo.Session        //数据库连接会话
	Databases map[string]DataBase //Key:数据库名陈 Value:DataBase对象
}

func NewHandler(address string, userName string, passWord string) (handler *Handler, err error) {
	handler = new(Handler)
	err = handler.Connect(address, userName, passWord)

	return handler, err
}

/******************************************
*           初始化连接
*******************************************/
func (this *Handler) init() {
	if this.DBsession != nil {
		this.DBsession.Close()
	}
	this.DBsession = nil
	this.Databases = make(map[string]DataBase)
}

/******************************************
*           建立一个数据库连接会话
*******************************************/
func (this *Handler) Connect(address string, userName string, passWord string) (err error) {
	//建立连接
	//mongodb://myuser:mypass@localhost:40001,otherhost:40001/mydb
	//var url string = "mongodb://" + aUserName + ":" + aUserPwd + "@" +
	//	aAddr + "/" + aDataBaseName
	//var url string = "mongodb://" + address + "/" + aDataBaseName

	this.init()

	var url string = "mongodb://" + address
	this.DBsession, err = mgo.Dial(url)
	if err != nil {
		errors.New("errro occurred!")
		return err
	}
	//设置模式
	this.DBsession.SetMode(mgo.Monotonic, true)
	return err
}

/******************************************
*           断开连接
*******************************************/
func (this *Handler) DisConnect() {
	if this.DBsession != nil {
		this.DBsession.Close()
	}
}

/******************************************
*           选择一个数据库
*******************************************/
func (this *Handler) DataBase(DBName string) (result *DataBase) {

	if this.DBsession == nil {
		return nil
	}
	dataBase, ok := this.Databases[DBName]

	if ok {
		result = &dataBase
	} else {
		result = this.newDataBase(DBName)
	}

	return result
}

/******************************************
*          添加一个数据库连接
*******************************************/
func (this *Handler) newDataBase(dbName string) *DataBase {
	if this.DBsession == nil {
		return nil
	}
	dataBase := DataBase{}
	dataBase.Init(dbName)
	dataBase.DBInstance = this.DBsession.DB(dbName)
	this.Databases[dbName] = dataBase
	return &dataBase
}

/******************************************
*           选择一个集合
*******************************************/
func (this *Handler) Collection(DBName string, CollectionName string) (result *mgo.Collection) {
	if this.DBsession == nil {
		return nil
	}
	dataBase := this.DataBase(DBName)
	result = dataBase.Collection(CollectionName)
	return result
}
