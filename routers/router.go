// @APIVersion 1.0.0
// @Title File Upload API
// @Description AuXpI 图床提供的 API 上传的方法
// @Contact aimerforreimu#gmail.com (#->@)
package routers

import (
	"auxpi/controllers"
	"auxpi/controllers/api/base"
	"auxpi/controllers/api/v1/user"
	"auxpi/middleware"
	"auxpi/routers/api/auth"
	"auxpi/routers/api/v1"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	//正式环境不使用控制器内环境，调试时使用控制器内反射路由
	beego.Router("/", &page.PagesController{}, "get:IndexShow")
	beego.Router("/Sina", &page.PagesController{}, "get:SinaShow")
	beego.Router("/Smms", &page.PagesController{}, "get:SmmsShow")
	beego.Router("/about", &page.PagesController{}, "get:AboutShow")

	if beego.BConfig.RunMode == "dev" {
		//auth
		auth.RegisterAuth()
		//Vue 调试的时候需要跨域 ()
		setCors()
		//部分需要调试的路由
		testRouter()
	}

	//v1 版本路由注册
	v1.RegisterControlApiV1()
	v1.RegisterOpenApiV1()
	//v2 版本路由注册

}

//测试路由，不要随便开启
func testRouter() {
	beego.InsertFilter("/api/v1/test/user_info", beego.BeforeRouter, middleware.JWT)
	beego.Router("/test", &base.ApiController{}, "post:Test")
	beego.Router("/auth/login", &base.ApiController{}, "post:LoginTest")
	beego.Router("/api/v1/test/user_info",&user.User{},"get:GetFakerUserInfo")

}

//跨域设置
func setCors() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://localhost:9527"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type","X-Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	//Vue 跨域请求，需要允许跨域
	beego.Router("*", &base.ApiController{}, "options:Options")
}
