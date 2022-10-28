package initialization

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/wiiCoder/go-scaffold/global"

	"github.com/olivere/elastic/v7"
	tb "gopkg.in/tucnak/telebot.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Global() {
	Config()      // 配置文件
	Mysql()       // Mysql初始化
	Redis()       // Redis初始化
	Elastic()     // Es初始化
	TelegramBot() // TG 机器人初始化
}

func Mysql() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.Config.MysqlConfig.Auth.Username,
		global.Config.MysqlConfig.Auth.Password,
		global.Config.MysqlConfig.Host,
		global.Config.MysqlConfig.Port,
		global.Config.MysqlConfig.Ds,
	)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	global.Db = db

}

func Redis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Config.RedisConfig.Host + ":" + string(global.Config.RedisConfig.Post),
		Password: global.Config.RedisConfig.Password,
		DB:       global.Config.RedisConfig.DB, // 默认使用 DB0
	})

	global.Rd = rdb
}

func Elastic() {
	logger := log.New(os.Stdout, global.Config.Base.Name, log.LstdFlags)

	var err error
	// 初始化变量
	esUrl := fmt.Sprintf("%s:%d", global.Config.EsConfig.Host, global.Config.EsConfig.Post)
	global.EsClient, err = elastic.NewClient(
		elastic.SetURL(esUrl), // 连接地址
		elastic.SetBasicAuth(
			global.Config.EsConfig.Auth.Username,
			global.Config.EsConfig.Auth.Password,
		), // 认证信息
		elastic.SetTraceLog(logger), // 增加日志输出
		elastic.SetSniff(false),     // 嗅探器开关，需要关闭
	)

	if err != nil {
		panic(err)
	}
}

// bot.Send([]string{"1241046678", "-1001854008739"}, "*bold text*\n_italic text_\n[inline URL](http://www.example.com/)\n[inline mention of a user](tg://user?id=123456789)\n`inline fixedwidth code`\n```\npreformatted fixedwidth code block\n```\n```python\npreformatted fixedwidth code block written in the Python programming language\n```")

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
