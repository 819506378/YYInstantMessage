package instant_message_main

import (
	"fmt"
	"runtime"
	"time"
	xorm_mysql "yyInstantMessage/common-library/xorm-mysql"
	"yyInstantMessage/common-library/xorm-mysql/xormMysqlConfig"
)

type instantMessageObj struct {
	orm *xorm_mysql.XormOptMysql
}

func NewInstantMessageMain() *instantMessageObj {
	return &instantMessageObj{}
}

//InitXormOptMysql 初始化xorm操作mysql数据库、设置连接池属性、开启ping数据库协程
func (im *instantMessageObj) InitXormOptMysql(param *xormMysqlConfig.XormMysqlParam) {
	im.orm = xorm_mysql.NewXormEngine(param)
	/* 设置最大连接数 */
	im.orm.SetXormMysqlMaxOpenConnections(param.MaxConnectionNum)
	/* 设置最大空闲连接数 */
	im.orm.SetXormMysqlMaxIdleConnections(param.MaxIdleConnectionNum)
	/* 设置连接生命周期 */
	im.orm.SetXomrMysqlConnectionMaxLifeTime(time.Minute * param.ConnectionLifeTime)

	/* 每隔一段时间ping一次数据库，保证数据库崩溃时，程序会有日志输出，方便查找问题 */
	//function.Must(nil, im.orm.Orm.Ping())
	go func() {
		ticker := time.NewTicker(time.Second * param.PingIntervalTime)
		defer ticker.Stop()
		for range ticker.C {
			if err := im.orm.Orm.Ping(); nil != err {
				fmt.Println(runtime.Caller(0))
				panic(err)
			}
		}
	}()
}
