package repository

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

var Mongo *mgo.Session

func init() {
	mongoDBUrl := beego.AppConfig.String("MongoDBUrl")
	beego.Debug(mongoDBUrl)
	session, err := mgo.Dial(mongoDBUrl)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	Mongo = session
}
