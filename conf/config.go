package conf

import (
	"strings"

	"github.com/sanyewudezhuzi/memo/model"

	"github.com/go-ini/ini"
)

var (
	// 服务模块
	AppMode  string
	HttpPort string
)

var (
	// MySQL 模块
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func Conf() {
	file := load_ini()

	read_ini(file)

	// mysql 连接
	dsn := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	model.ConnectMySQL(dsn)

	// 数据库迁移
	model.SyncSchema()
}

// 加载 .ini
func load_ini() *ini.File {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		panic("Profile missing.")
	}
	return file
}

// 读取 .ini
func read_ini(f *ini.File) {
	// service
	AppMode = f.Section("service").Key("AppMode").String()
	HttpPort = f.Section("service").Key("HttpPort").String()

	// mysql
	Db = f.Section("mysql").Key("Db").String()
	DbHost = f.Section("mysql").Key("DbHost").String()
	DbPort = f.Section("mysql").Key("DbPort").String()
	DbUser = f.Section("mysql").Key("DbUser").String()
	DbPassWord = f.Section("mysql").Key("DbPassWord").String()
	DbName = f.Section("mysql").Key("DbName").String()
}
