package xorm_mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"runtime"
	"time"
	"yyInstantMessage/common-library/xorm-mysql/xormMysqlConfig"
)

type XormOptMysql struct {
	Orm *xorm.Engine
}

//NewXormEngine 创建xorm关于mysql数据的引擎
func NewXormEngine(xormCfg *xormMysqlConfig.XormMysqlParam) *XormOptMysql {
	engine, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s",
		xormCfg.UserName, xormCfg.PassWord, xormCfg.Ip, xormCfg.Port, xormCfg.Charset))
	if nil != err {
		fmt.Println("mysql数据库连接失败,config: %v", xormCfg)
		fmt.Println(runtime.Caller(0))
		panic(err)
	}

	fmt.Println("mysql数据库连接成功")

	return &XormOptMysql{
		Orm: engine,
	}
}

//SetXormMysqlMaxIdleConnections 设置xorm操作mysql最大空闲连接数
func (xpm *XormOptMysql) SetXormMysqlMaxIdleConnections(idleConn int) {
	xpm.Orm.SetMaxIdleConns(idleConn)
}

//SetXormMysqlMaxOpenConnections 设置xorm操作mysql最大连接数
func (xpm *XormOptMysql) SetXormMysqlMaxOpenConnections(openConn int) {
	xpm.Orm.SetMaxOpenConns(openConn)
}

//SetXomrMysqlConnectionMaxLifeTime 设置xorm操作mysql连接生命周期
func (xpm *XormOptMysql) SetXomrMysqlConnectionMaxLifeTime(connLifeTime time.Duration) {
	xpm.Orm.SetConnMaxLifetime(connLifeTime)
}
