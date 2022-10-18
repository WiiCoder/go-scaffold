package global

import (
	"github.com/olivere/elastic/v7"
	"github.com/wiiCoder/roicloud-youshu/server/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config   config.ServerConfig
	Log      *zap.Logger
	EsClient *elastic.Client
	Db       *gorm.DB
)
