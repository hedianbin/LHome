package routers

import (
	"loveHome/controllers"
	"github.com/astaxie/beego"
	"strings"
	"net/http"
	"github.com/astaxie/beego/context"
	"loveHome/utils"
)

func init() {
	ignoreStaticPath()
	ns := beego.NewNamespace("/api",beego.NSCond(func(ctx *context.Context) bool {
		if ctx.Input.Domain()==utils.G_server_addr{
			return true
		}
		beego.Debug("now domain is ", ctx.Input.Domain(), " not "+utils.G_server_addr)
		return false
	}),
		beego.NSNamespace("/v1.0",
		//beego.Router("/", &controllers.MainController{}),
		/*
		请求地理区域信息
		Request URL: http://localhost:8899/api/v1.0/areas
		Request Method: GET
		*/
		beego.NSRouter("/areas", &controllers.AreaController{}, "get:GetArea"),
		/*
		创建session和退出登录
			Request URL: http://localhost:8899/api/v1.0/session
			Request Method: GET
		*/

		beego.NSRouter("/session", &controllers.SessionController{}, "get:GetSessionData;delete:DeleteSessionData"),
		/*
		注册
			Request URL: http://localhost:8899/api/v1.0/users
			Request Method: POST
		*/
		beego.NSRouter("/users", &controllers.UserController{}, "post:Reg"),
		/*
		登录
			Request URL: http://localhost:8899/api/v1.0/sessions
			Request Method: POST
		*/
		beego.NSRouter("/sessions", &controllers.SessionController{}, "post:Login"),
		/*
		上传头像
				Request URL: http://localhost:8899/api/v1.0/user/avatar
				Request Method: POST
			*/
		beego.NSRouter("/user/avatar", &controllers.UserController{}, "post:PostAvatar"),
		/*
		个人信息
			1. Request URL: http://localhost:8899/api/v1.0/user
			2. Request Method: GET
		*/
		beego.NSRouter("/user", &controllers.UserController{}, "get:GetUserData"),
		/*
		更新用户名
			1. Request URL: http://localhost:8899/api/v1.0/user/name
			2. Request Method: PUT
		*/
		beego.NSRouter("/user/name", &controllers.UserController{}, "put:UpdateName"),
		/*
		实名认证GET,POST
			1. Request URL: http://localhost:8899/api/v1.0/user/auth
			2. Request Method: GET
		*/
		beego.NSRouter("/user/auth", &controllers.UserController{}, "get:GetUserData;post:PostRealName"),
		/*
		请求当前用户已发布房源
			1. Request URL: http://localhost:8899/api/v1.0/user/houses
			2. Request Method: GET
		*/
		beego.NSRouter("/user/houses", &controllers.HouseController{}, "get:GetHouseData"),
		/*
		发布房源信息post:PostHouseData
			1. Request URL: http://localhost:8899/api/v1.0/houses
			2. Request Method: POST
		获取用户搜索房源信息get:GetHouseSearchData
			Request URL: http://10.0.151.242:9999/api/v1.0/houses?aid=1&sd=2018-06-27&ed=2018-06-28&sk=new&p=1
			Request Method: GET
		*/
		beego.NSRouter("/houses", &controllers.HouseController{}, "post:PostHouseData;get:GetHouseSearchData"),
		/*
		房源详细信息
			1. Request URL: http://localhost:8899/api/v1.0/houses/2
			2. Request Method: GET
		*/
		beego.NSRouter("/houses/?:id", &controllers.HouseController{}, "get:GetDetailHouseData"),
		/*
		房源图片上传
			1. Request URL: http://10.0.151.242:8899/api/v1.0/houses/8/images
			2. Request Method: POST
		*/
		beego.NSRouter("/houses/?:id/images", &controllers.HouseController{}, "post:UploadHouseImage"),
		/*
		   用户请求房源首页列表信息
		   Request URL: http://localhost:8899/api/v1.0/houses/index
		   Request Method: GET
		   */
		beego.NSRouter("/houses/index", &controllers.HouseController{},"get:GetHouseIndex"),
		/*
		提交订单
			Request URL: http://localhost:8899/api/v1.0/orders
			Request Method: POST
		*/
		beego.NSRouter("/orders", &controllers.OrderController{}, "post:PostOrderHouseData"),
		/*
		   我的订单，租客订单
			1. Request URL: http://10.0.151.242:8899/api/v1.0/user/orders?role=custom
			2. Request Method: GET
		   */
		beego.NSRouter("/user/orders", &controllers.OrderController{}, "get:GetOrderData"),
		/*
		   房东处理订单
			1. Request URL: http://10.0.151.242:8899/api/v1.0/orders/4/status
			2. Request Method: PUT
		   */
		beego.NSRouter("/orders/:id/status", &controllers.OrderController{}, "put:OrderStatus"),
		/*
		用户发送订单评价信息
			1. Request URL: http://10.0.151.242:8899/api/v1.0/orders/6/comment
			2. Request Method: PUT
		*/
		beego.NSRouter("/orders/:id/comment", &controllers.OrderController{}, "put:OrderComment"),
		),//要地方换行就需要逗号
	)
	//注册namespace
	beego.AddNamespace(ns)
}


func ignoreStaticPath() {
	//透明static
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}
func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url: ", orpath)
	//如果请求url还有api字段，说明是指令应该取消静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)
}
