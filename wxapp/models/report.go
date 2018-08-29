package models

import (
	"gofising.vip/wxapp/repository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// 战报
type Report struct {
	ReportId      bson.ObjectId `bson:"_id"`
	UserId        int           `bson:"uid"`
	Title         string        `bson:"title"`
	Location      Coordinate    `bson:"loc"`
	Address       string        `bson:"addr"`
	ViewNum       int           `bson:"vn"`
	StartNum      int           `bson:"sn"`
	CollectionNum int           `bson:"cn"`
	CreateTime    time.Time     `bson:"ctime"`
	UpdateTime    time.Time     `bson:"utime"`
}

// 坐标
type Coordinate struct {
	Longitude float64 `bson:"lon"`
	Latitude  float64 `bson:"lat"`
}

var reportCollection *mgo.Collection

func init() {
	reportCollection = repository.Mongo.DB("YmtProducts").C("Report")
	err := reportCollection.EnsureIndexKey("$2d:loc")

	if err != nil {
		panic(err)
	}
}

func AddReport(report Report) (rid bson.ObjectId, err error) {
	rid = bson.NewObjectId()

	report.ReportId = rid
	report.CreateTime = bson.Now()
	report.UpdateTime = bson.Now()
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
