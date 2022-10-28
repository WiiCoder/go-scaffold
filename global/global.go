package global

import (
	"github.com/olivere/elastic/v7"
	"github.com/wiiCoder/go-scaffold/config"
	"go.uber.org/zap"
	"gopkg.in/tucnak/telebot.v2"
	"gorm.io/gorm"
)

var (
	Config   config.ServerConfig
	Log      *zap.Logger
	EsClient *elastic.Client
	Db       *gorm.DB
	TgBot    *telebot.Bot
)
