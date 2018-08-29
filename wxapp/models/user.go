package models

import (
	"errors"
	"github.com/astaxie/beego"
	"gofising.vip/wxapp/repository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
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
	Level           int           `bson:"level"`
	UserStat        UserStat      `bson:"us"`
}

// 用户统计信息
type UserStat struct {
	ReportNum     int `bson:"rn"`
	FansNum       int `bson:"fn"`
	AttentionNum  int `bson:"an"`
	CollectionNum int `bson:"cn"`
	ViewNum       int `bson:"vn"`
}

var UserStatBson map[string]string

func init() {
	UserStatBson = make(map[string]string, 5)

	usType := reflect.TypeOf(UserStat{})
	for i := 0; i < usType.NumField(); i++ {
		field := usType.Field(i)
		UserStatBson[field.Name] = "us." + field.Tag.Get("bson")
	}

	beego.Debug(UserStatBson)
}

var userCollection *mgo.Collection

func init() {
	userCollection = repository.Mongo.DB("YmtProducts").C("User")
}

func GetUser(uid string) (user *User, err error) {
	user = &User{}
	err = userCollection.FindId(bson.ObjectIdHex(uid)).One(user)

	return
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

func UpdateUser(user User) (err error) {
	userPo := &User{}
	err = userCollection.FindId(user.UserId).One(userPo)
	beego.Debug(userPo)

	if user.NickName != "" {
		userPo.NickName = user.NickName
	}

	if user.ProfileTag != nil {
		userPo.ProfileTag = user.ProfileTag
	}

	if user.Level > 1 {
		userPo.Level = user.Level
	}

	userPo.UpdateTime = bson.Now()
	userCollection.UpsertId(userPo.UserId, userPo)

	return errors.New("User Not Exist")
}

func UpdateUserStat(uid string, dataMap map[string]int) (err error) {
	selector := bson.M{"_id": bson.ObjectIdHex(uid)}

	updateField := bson.M{}
	for k, v := range dataMap {
		updateField[UserStatBson[k]] = v
	}

	beego.Debug(updateField)

	data := bson.M{"$set": updateField}

	err = userCollection.Update(selector, data)
	return err
}
