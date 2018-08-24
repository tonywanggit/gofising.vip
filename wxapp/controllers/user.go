package controllers

import (
	"encoding/json"
	"gofising.vip/wxapp/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {string} models.User.UserId
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Post() {
	var user models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)

	beego.Debug(&user)

	uid, err := models.AddUser(user)
	if err != nil {
		c.Data["json"] = map[string]string{"err": err.Error()}
	} else {
		c.Data["json"] = map[string]string{"rid": uid.Hex()}
	}

	c.ServeJSON()
}
