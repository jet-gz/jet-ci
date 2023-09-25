package db

import (
	"jet-ci/pkg/config"
	"jet-ci/pkg/env"
)

type Options struct {
	//DB类型
	DBType string `yaml:"db_type"`
	//连接字符串
	ConnectStr string `yaml:"connect_str"`
	// 释放debug
	IsDebug bool `yaml:"is_debug"`
	// 日志
	LogLevel int `yaml:"log_level"`
}

var options Options = Default()

func Default() Options {
	return Options{
		DBType:     "mysql",
		ConnectStr: "root:123456@tcp(121.4.181.166:3306)/gormTest?charset=utf8mb4&parseTime=True&loc=Local",
		IsDebug:    true,
		LogLevel:   2,
	}
}
func GetOptions() Options {
	return options
}

func SetOptions(opts Options) {
	options = opts
}

func init() {
	//首先加载环境变量
	options.FromEnv()
}

func (options *Options) FromEnv() {
	options.DBType = env.Get("TYPE", options.DBType)
	options.ConnectStr = env.Get("CONNECT_STR", options.ConnectStr)
	options.IsDebug = env.GetBool("IS_DEBUG", options.IsDebug)
	options.LogLevel = env.GetInt("LOG_LEVEL", options.LogLevel)
}

// 特别敏感的信息
const Path = "d:/Git/jet-ci/internal/configs/mysql-db.yml"

func Load() error {
	return config.Load(Path, &options)
}

func Store() error {
	return config.Store(Path, &options)
}
