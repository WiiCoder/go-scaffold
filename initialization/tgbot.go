package initialization

import (
	"net/http"
	"net/url"
	"time"

	"github.com/wiiCoder/go-scaffold/global"
	tb "gopkg.in/tucnak/telebot.v2"
)

//bot.Send([]string{"1241046678", "-1001854008739"}, "*bold text*\n_italic text_\n[inline URL](http://www.example.com/)\n[inline mention of a user](tg://user?id=123456789)\n`inline fixedwidth code`\n```\npreformatted fixedwidth code block\n```\n```python\npreformatted fixedwidth code block written in the Python programming language\n```")

func TelegramBot() {
	// 部署的服务器在国内需要设置代理，否则无法正常使用
	uri, _ := url.Parse(global.Config.TgConfig.Proxy)
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(uri),
		},
	}
	bot, err := tb.NewBot(tb.Settings{
		URL:    global.Config.TgConfig.Api,
		Token:  global.Config.TgConfig.Token,
		Poller: &tb.LongPoller{Timeout: 5 * time.Second},
		Client: &client,
	})
	if err != nil {
		panic(err)
	}

	global.TgBot = bot
}
