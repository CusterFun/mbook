package routers

import (
	"github.com/astaxie/beego"
	"github.com/custergo/mbook/controllers"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	// 首页&分类&详情
	beego.Router("/", &controllers.HomeController{}, "get:Index")               // 首页
	beego.Router("/explore", &controllers.ExploreController{}, "get:Index")     // 分类
	beego.Router("/books/:key", &controllers.DocumentController{}, "get:Index") // 图书目录&详情

	// 读书
	beego.Router("/read/:key/:id", &controllers.DocumentController{}, "*:Read")         // 图书章节阅读
	beego.Router("/read/:key/search", &controllers.DocumentController{}, "post:Search") // 图书章节内搜索

	// 编辑
	beego.Router("/api/:key/edit/?:id", &controllers.DocumentController{}, "*:Edit")       // 文档编辑
	beego.Router("/api/:key/content/?:id", &controllers.DocumentController{}, "*:Content") // 保存文档
	beego.Router("/api/upload", &controllers.DocumentController{}, "post:Upload")          // 上传图片
	beego.Router("/api/:key/create", &controllers.DocumentController{}, "post:Create")     // 创建章节
	beego.Router("/api/:key/delete", &controllers.DocumentController{}, "post:Delete")     // 删除章节

	// 搜索
	beego.Router("/search", &controllers.SearchController{}, "get:Search")        // 首页搜索
	beego.Router("/search/result", &controllers.SearchController{}, "get:Result") // 搜索结果页

	// 登录
	beego.Router("/login", &controllers.AccountController{}, "*:Login")          // 登录页面
	beego.Router("/regist", &controllers.AccountController{}, "*:Regist")        // 注册页面
	beego.Router("/logout", &controllers.AccountController{}, "*:Logout")        // 登出页面
	beego.Router("/doregist", &controllers.AccountController{}, "post:DoRegist") // 注册提交的post请求

	//用户图书管理
	beego.Router("/book", &controllers.BookController{}, "*:Index")                         //我的图书
	beego.Router("/book/create", &controllers.BookController{}, "post:Create")              //创建图书
	beego.Router("/book/:key/setting", &controllers.BookController{}, "*:Setting")          //图书设置
	beego.Router("/book/setting/upload", &controllers.BookController{}, "post:UploadCover") //图书封面
	beego.Router("/book/star/:id", &controllers.BookController{}, "*:Collection")           //收藏图书
	beego.Router("/book/setting/save", &controllers.BookController{}, "post:SaveBook")      //保存
	beego.Router("/book/:key/release", &controllers.BookController{}, "post:Release")       //发布
	beego.Router("/book/setting/token", &controllers.BookController{}, "post:CreateToken")  //创建Token

	//个人中心
	beego.Router("/user/:username", &controllers.UserController{}, "get:Index")                 //分享
	beego.Router("/user/:username/collection", &controllers.UserController{}, "get:Collection") //收藏
	beego.Router("/user/:username/follow", &controllers.UserController{}, "get:Follow")         //关注
	beego.Router("/user/:username/fans", &controllers.UserController{}, "get:Fans")             //粉丝
	beego.Router("/follow/:uid", &controllers.BaseController{}, "get:SetFollow")                //关注或取消关注
	beego.Router("/book/score/:id", &controllers.BookController{}, "*:Score")                   //评分
	beego.Router("/book/comment/:id", &controllers.BookController{}, "post:Comment")            //评论

	//个人设置
	beego.Router("/setting", &controllers.SettingController{}, "*:Index")         // get请求设置页面，post请求更新设置
	beego.Router("/setting/upload", &controllers.SettingController{}, "*:Upload") // 上传头像

	//管理后台
	beego.Router("/manager/category", &controllers.ManagerController{}, "post,get:Category")    // 分类管理
	beego.Router("/manager/update-cate", &controllers.ManagerController{}, "get:UpdateCate")    // 更新分类
	beego.Router("/manager/del-cate", &controllers.ManagerController{}, "get:DelCate")          // 删除分类
	beego.Router("/manager/icon-cate", &controllers.ManagerController{}, "post:UpdateCateIcon") // 更新分类图标
}
