package main

import (
	"github.com/wiiCoder/go-scaffold/initialization"
	"github.com/wiiCoder/go-scaffold/service/telegram"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	// 初始化配置 配置文件 > Mysql > Redis > Elasticsearch > tg bot > gin-i18n > gin
	initialization.Global()

	h := telegram.New()
	h.Send([]string{}, "", tb.ModeMarkdownV2)

	// TODO: 优雅关闭应用
}
