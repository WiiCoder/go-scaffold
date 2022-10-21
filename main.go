package main

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	// 初始化配置 配置文件 > Mysql > Redis > Elasticsearch > tg bot > i18n > gin
	//initialization.Global()
	bot := TelegramConfigExample()
	bot.Send([]string{"1241046678", "-1001854008739"}, "*bold text*\n_italic text_\n[inline URL](http://www.example.com/)\n[inline mention of a user](tg://user?id=123456789)\n`inline fixedwidth code`\n```\npreformatted fixedwidth code block\n```\n```python\npreformatted fixedwidth code block written in the Python programming language\n```")
}

type TelegramConfig struct {
	Api   string `mapstructure:"api"`
	Token string `mapstructure:"token"`
}

func TelegramConfigExample() *TelegramConfig {
	return &TelegramConfig{
		"https://api.telegram.org",
		"2138454577:AAG7DRXBkEJNjzsmjXl2F_neTum9f8Ajxvk",
	}
}

func (cfg *TelegramConfig) Send(targets []string, message string) {
	if cfg.Api == "" {
		cfg.Api = "https://api.telegram.org"
	}

	uri, _ := url.Parse("http://127.0.0.1:7890")
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(uri),
		},
	}
	bot, err := tb.NewBot(tb.Settings{
		URL:    cfg.Api,
		Token:  cfg.Token,
		Poller: &tb.LongPoller{Timeout: 5 * time.Second},
		Client: &client,
	})
	if err != nil {
		logrus.Errorf("Telegram alarm send failed: %s", err)
		return
	}

	opt := &tb.SendOptions{ParseMode: tb.ModeMarkdownV2}
	for _, t := range targets {
		to, err := strconv.Atoi(t)
		if err != nil {
			logrus.Errorf("Telegram alarm send failed [%s]: %s", t, err)
			continue
		}
		_, err = bot.Send(tb.ChatID(to), "*bold \\*text*\n_italic \\*text_\n__underline__\n~strikethrough~\n||spoiler||\n*bold _italic bold ~italic bold strikethrough ||italic bold strikethrough spoiler||~ __underline italic bold___ bold*\n[inline URL](http://www.example.com/)\n[inline mention of a user](tg://user?id=123456789)\n`inline fixed-width code`\n```\npre-formatted fixed-width code block\n```\n```python\npre-formatted fixed-width code block written in the Python programming language\n```", opt)
		if err != nil {
			logrus.Errorf("Telegram alarm send failed [%s]: %s", t, err)
			continue
		}
	}
}
