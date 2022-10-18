package initialization

import (
	"fmt"
	"log"
	"os"

	"github.com/olivere/elastic/v7"

	"github.com/wiiCoder/roicloud-youshu/server/global"
)

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
