package models

import (
	"gofising.vip/wxapp/repository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// 战报1
type Report struct {
	ReportId   bson.ObjectId `bson:"_id"`
	UserId     string        `bson:"uid"`
	Remark     string        `bson:"remark"`
	CreateTime time.Time     `bson:"ctime"`
}

var reportCollection *mgo.Collection

func init() {
	reportCollection = repository.Mongo.DB("YmtProducts").C("Report")
}

func AddReport(report Report) (rid bson.ObjectId, err error) {
	rid = bson.NewObjectId()

	report.ReportId = rid
	report.CreateTime = bson.Now()
	err = reportCollection.Insert(report)

	return
}

func GetReport(rid string) (report *Report, err error) {

	report = &Report{}
	err = reportCollection.FindId(bson.ObjectIdHex(rid)).One(report)

	return
}

func GetAllReport() (reports []*Report, err error) {
	reports = make([]*Report, 1)

	err = reportCollection.Find(nil).All(&reports)

	return reports, err

}
