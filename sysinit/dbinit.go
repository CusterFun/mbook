package sysinit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/custergo/mbook/models" // 初始化models里注册的类需要先执行models里的init()函数
)

// 建立数据库连接,参数alias可能代表w主库写库
func dbinit(alias string) {
	dbAlias := alias // default
	if "w" == alias || "default" == alias || len(alias) <= 0 {
		dbAlias = "default" // 默认库的链接
		alias = "w"
	}

	// 拼接数据库链接
	// 数据库名称
	dbName := beego.AppConfig.String("db_" + alias + "_database")
	// 数据库用户名
	dbUser := beego.AppConfig.String("db_" + alias + "_username")
	// 数据库密码
	dbPwd := beego.AppConfig.String("db_" + alias + "_password")
	// 数据库IP
	dbHost := beego.AppConfig.String("db_" + alias + "_host")
	// 数据库端口
	dbPort := beego.AppConfig.String("db_" + alias + "_port")
	// root:root1234@tcp(127.0.0.1:3306)/mbook?charset=urf8
	orm.RegisterDataBase(dbAlias, "mysql",
		dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8", 30)
	// 如果是开发模式 isDev 为true
	isDev := ("dev" == beego.AppConfig.String("runmode"))
	// 主库自动建表
	if "w" == alias {
		// false: 发生错误时，不继续执行下一条sql
		// true: 打印日志详细信息
		orm.RunSyncdb("default", false, isDev)
	}
	if isDev {
		orm.Debug = isDev
	}
}
