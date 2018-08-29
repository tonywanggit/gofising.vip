package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:ObjectController"] = append(beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:ObjectController"] = append(beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:ObjectController"] = append(beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:ObjectController"] = append(beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:ObjectController"] = append(beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:ReportController"] = append(beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:ReportController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:ReportController"] = append(beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:ReportController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:ReportController"] = append(beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:ReportController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:rid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:TokenController"] = append(beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:TokenController"],
		beego.ControllerComments{
			Method:           "GetToken",
			Router:           `/:token`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:TokenController"] = append(beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:TokenController"],
		beego.ControllerComments{
			Method:           "WxLogin",
			Router:           `/wxLogin`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:UserController"] = append(beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:UserController"],
		beego.ControllerComments{
			Method:           "CreateUser",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:UserController"] = append(beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetUser",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:UserController"] = append(beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:UserController"],
		beego.ControllerComments{
			Method:           "ModifyUser",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:UserController"] = append(beego.GlobalControllerRouter["gofising.vip/wxapp/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UpdateUserStat",
			Router:           `/stat/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

}
