package controllers

import (
	"encoding/json"
	"gofising.vip/wxapp/models"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @router /:uid [get]
func (c *UserController) GetUser() {
	rid := c.GetString(":uid")
	if rid != "" {
		user, err := models.GetUser(rid)
		if err != nil {
			c.Data["json"] = models.FailResponse("user not found")
		} else {
			c.Data["json"] = models.SuccessResponse(user)
		}
	}
	c.ServeJSON()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {string} models.User.UserId
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) CreateUser() {
	var user models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)

	beego.Debug(&user)

	uid, err := models.AddUser(user)
	if err != nil {
		c.Data["json"] = map[string]string{"err": err.Error()}
	} else {
		c.Data["json"] = map[string]string{"uid": uid.Hex()}
	}

	c.ServeJSON()
}

// @Title ModifyUser
// @Description update the user
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {string} models.User.UserId
// @Failure 403 body is empty
// @router /:uid [put]
func (c *UserController) ModifyUser() {
	uid := c.GetString(":uid")

	var user models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	user.UserId = bson.ObjectIdHex(uid)

	err := models.UpdateUser(user)
	if err != nil {
		c.Data["json"] = map[string]string{"err": err.Error()}
	} else {
		c.Data["json"] = map[string]string{"uid": user.UserId.Hex()}
	}

	c.ServeJSON()
}

// @Title UpdateUserStat
// @router /stat/:uid [put]
func (c *UserController) UpdateUserStat() {
	var data map[string]int
	json.Unmarshal(c.Ctx.Input.RequestBody, &data)

	beego.Debug(data)

	uid := c.GetString(":uid")
	//
	//var userStat models.UserStat
	//json.Unmarshal(c.Ctx.Input.RequestBody, &userStat)

	err := models.UpdateUserStat(uid, data)
	if err != nil {
		c.Data["json"] = map[string]string{"err": err.Error()}
	} else {
		c.Data["json"] = map[string]string{"uid": uid}
	}

	c.ServeJSON()
}
