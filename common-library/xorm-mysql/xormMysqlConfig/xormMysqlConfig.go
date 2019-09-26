package xormMysqlConfig

import "time"

//XormMysqlParam mysql数据库连接配置
type XormMysqlParam struct {
	UserName             string        `json:"user_name"`               // mysql 用户名
	PassWord             string        `json:"pass_word"`               // mysql 密码
	DatabaseName         string        `json:"database_name"`           // mysql 数据库名
	Ip                   string        `json:"ip"`                      // mysql IP地址
	Port                 int16         `json:"port"`                    // mysql 端口
	Charset              string        `json:"charset"`                 // mysql 字符集设置 （对于utf8的支持不够， utf8mb4才能完全支持utf8）
	MaxConnectionNum     int           `json:"max_connection_num"`      // mysql 最大连接数
	MaxIdleConnectionNum int           `json:"max_idle_connection_num"` // mysql 最大空闲连接数
	ConnectionLifeTime   time.Duration `json:"connection_life_time"`    // mysql 连接生命周期（单位分钟)
	PingIntervalTime     time.Duration `json:"ping_interval_time"`      // ping数据库时间间隔
}
