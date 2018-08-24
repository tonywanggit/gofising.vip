package models

import (
	"gofising.vip/wxapp/repository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// 用户
type User struct {
	UserId          bson.ObjectId `bson:"_id"`
	NickName        string        `bson:"nick"`
	LoginMobile     string        `bson:"mobile"`
	PassWord        string        `bson:"pwd"`
	RegisterArea    string        `bson:"area"`
	RegisterTime    time.Time     `bson:"rtime"`
	UpdateTime      time.Time     `bson:"utime"`
	LastVisitorTime time.Time     `bson:"ltime"`
	ProfileTag      []int         `bson:"tag"`
}

var userCollection *mgo.Collection

func init() {
	userCollection = repository.Mongo.DB("YmtProducts").C("User")
}

func AddUser(user User) (uid bson.ObjectId, err error) {
	uid = bson.NewObjectId()

	user.UserId = uid
	user.RegisterTime = bson.Now()
	user.LastVisitorTime = bson.Now()
	user.UpdateTime = bson.Now()

	err = userCollection.Insert(user)

	return
}
