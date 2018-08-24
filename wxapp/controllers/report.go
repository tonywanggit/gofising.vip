package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"gofising.vip/wxapp/models"
	"runtime"
)

type ReportController struct {
	beego.Controller
}

// @router / [post]
func (c *ReportController) Post() {
	var report models.Report
	json.Unmarshal(c.Ctx.Input.RequestBody, &report)

	beego.Debug(&report)

	rid, err := models.AddReport(report)
	if err != nil {
		c.Data["json"] = map[string]string{"err": err.Error()}
	} else {
		c.Data["json"] = map[string]string{"rid": rid.Hex()}
	}

	c.ServeJSON()
}

// @router / [get]
func (c *ReportController) GetAll() {
	reports, err := models.GetAllReport()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = reports
	}
	c.ServeJSON()
}

// @router /:rid [get]
func (c *ReportController) Get() {
	defer func() {
		if err := recover(); err != nil {
			var stack string
			for i := 1; ; i++ {
				_, file, line, ok := runtime.Caller(i)
				if !ok {
					break
				}
				logs.Critical(fmt.Sprintf("%s:%d", file, line))
				stack = stack + fmt.Sprintln(fmt.Sprintf("%s:%d", file, line))
			}

			beego.Debug(err)
			beego.Debug(stack)
			c.Data["json"] = "system error"
			c.ServeJSON()
		}
	}()

	rid := c.GetString(":rid")
	if rid != "" {
		user, err := models.GetReport(rid)
		if err != nil {
			c.Data["json"] = err.Error()
		} else {
			c.Data["json"] = user
		}
	}
	c.ServeJSON()
}
