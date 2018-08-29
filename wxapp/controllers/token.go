package controllers

import (
	"github.com/astaxie/beego"
	"gofising.vip/wxapp/models"
)

type TokenController struct {
	beego.Controller
}

// @Title WxLogin
// @Description Logs user into the system
// @Param	wxcode		query 	string	true		"The username for login"
// @Success 200 {string} login success
// @router /wxLogin [get]
func (c *TokenController) WxLogin() {
	wxcode := c.GetString("code")

	c.Data["json"] = models.SuccessResponse(wxcode)
	c.ServeJSON()
}

// @router /:token [get]
func (c *TokenController) GetToken() {
	token := c.GetString(":token")
	c.Data["json"] = models.FailResponse(token)
	c.ServeJSON()
}
