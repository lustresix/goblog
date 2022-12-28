package config

import (
	"fmt"
	"goblog/model"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	AppMode  string
	HttpPort string
	JWTKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	AccessKey string
	SecretKey string
	Bucket    string
	Sever     string
	Zone      int
)

func init() {
	file, err := ini.Load("./config/conf.ini") // 加载文件
	if err != nil {
		fmt.Println("配置文件读取失败", err)
	}

	LoadSever(file)
	LoadMysql(file)
	LoadUpload(file)

	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	// 引用数据库
	model.Database(path)
}

func LoadSever(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
	JWTKey = file.Section("service").Key("JWTKey").String()
}

func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db ").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
func LoadUpload(file *ini.File) {
	AccessKey = file.Section("load").Key("AccessKey").String()
	SecretKey = file.Section("load").Key("SecretKey").String()
	Bucket = file.Section("load").Key("Bucket").String()
	Sever = file.Section("load").Key("Sever").String()
	Zone = file.Section("load").Key("Zone").MustInt(1)
}
