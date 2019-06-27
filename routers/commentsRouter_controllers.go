package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["user-api/controllers:AdminController"] = append(beego.GlobalControllerRouter["user-api/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user-api/controllers:AdminController"] = append(beego.GlobalControllerRouter["user-api/controllers:AdminController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user-api/controllers:AdminController"] = append(beego.GlobalControllerRouter["user-api/controllers:AdminController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user-api/controllers:AdminController"] = append(beego.GlobalControllerRouter["user-api/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user-api/controllers:AdminController"] = append(beego.GlobalControllerRouter["user-api/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user-api/controllers:AdminController"] = append(beego.GlobalControllerRouter["user-api/controllers:AdminController"],
        beego.ControllerComments{
            Method: "GetToken",
            Router: `/getToken`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user-api/controllers:AdminController"] = append(beego.GlobalControllerRouter["user-api/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user-api/controllers:AdminController"] = append(beego.GlobalControllerRouter["user-api/controllers:AdminController"],
        beego.ControllerComments{
            Method: "SendSms",
            Router: `/sendsms`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user-api/controllers:AdminController"] = append(beego.GlobalControllerRouter["user-api/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Status",
            Router: `/status`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
