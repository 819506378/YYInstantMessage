package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	snowflake "yyInstantMessage/common-library/snowFlake"
	"yyInstantMessage/common-library/xorm-mysql/xormMysqlConfig"
	instant_message_main "yyInstantMessage/instant-message/instant-message-main"
)

func init() {}

func main() {
	configDir := os.Getenv("YYINSTANTMESSAGECONFIG")
	imObj := instant_message_main.NewInstantMessageMain()

	/* 设置雪花ID的机器号 */
	snowflake.SetMachineId(1)
	/* 初始化日志库 */
	InitLog()
	/* 初始化数据库 */
	imObj.InitXormOptMysql(GetXormConfig(configDir))
	/* 初始化redis */

}

//InitLog
func InitLog() {}

//GetXormConfig 读取mysql配置文件
func GetXormConfig(configDir string) *xormMysqlConfig.XormMysqlParam {
	mysqlConfBytes, err := ioutil.ReadFile(configDir + "/mysqlConf.json")
	if nil != err {
		fmt.Println(runtime.Caller(0))
		panic(err)
	}
	xormMysqlParam := &xormMysqlConfig.XormMysqlParam{}
	err = json.Unmarshal(mysqlConfBytes, xormMysqlParam)
	if nil != err {
		fmt.Println(runtime.Caller(0))
		panic(err)
	}

	return xormMysqlParam
}
