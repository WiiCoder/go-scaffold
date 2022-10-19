package main

import (
	"github.com/wiiCoder/go-scaffold/initialization"
)

func main() {
	// 初始化配置 配置文件 -> Mysql -> Redis -> Elasticsearch -> i18n -> gin
	initialization.Global()

}
