package sysinit

import (
	"github.com/astaxie/beego"
	"path/filepath"
	"strings"
)

func sysinit() {
	// 上传静态资源路径
	uploads := filepath.Join("./", "uploads") // 上传路径
	// 访问域名？uploads即访问到静态资源
	beego.BConfig.WebConfig.StaticDir["/uploads"] = uploads

	// 注册前端使用函数
	registerFunctions()
}

// 静态资源文件views访问函数，需要先在这里声明
func registerFunctions() {
	beego.AddFuncMap("cdnjs", func(p string) string {
		cdn := beego.AppConfig.DefaultString("cdnjs", "") // 配置中有没有cdnjs的参数
		// 有这个参数并以"/"开头和结尾
		if strings.HasPrefix(p, "/") && strings.HasSuffix(cdn, "/") {
			return cdn + string(p[1:]) // 进行拼接域名和参数
		}
	})
}
