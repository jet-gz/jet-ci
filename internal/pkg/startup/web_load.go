package startup

import (
	"jet-ci/pkg/config"
	"jet-ci/pkg/env"
)

type Options struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
	IsDebug bool   `yaml:"is_debug"`
	IsGzip  bool   `yaml:"is_gzip"`
}

var options Options = Default()

func Default() Options {
	return Options{
		Address: "0.0.0.0",
		Port:    5000,
		IsDebug: true,
		IsGzip:  true,
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
	options.Address = env.Get("address", options.Address)
	options.Port = env.GetInt("port", options.Port)
	options.IsDebug = env.GetBool("IS_DEBUG", options.IsDebug)
	options.IsGzip = env.GetBool("is_gzip", options.IsGzip)
}

// 特别敏感的信息配置文件不放项目里面
const Path = "d:/Git/jet-ci/internal/configs/mysql-db.yml"

func Load() error {
	return config.Load(Path, &options)
}

func Store() error {
	return config.Store(Path, &options)
}
