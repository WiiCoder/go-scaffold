package initialization

import (
	"fmt"

	"github.com/wiiCoder/roicloud-youshu/server/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Mysql() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Asia/Shanghai",
		global.Config.MysqlConfig.Auth.Username,
		global.Config.MysqlConfig.Auth.Password,
		global.Config.MysqlConfig.Host,
		global.Config.MysqlConfig.Port,
		global.Config.MysqlConfig.Ds,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	global.Db = db

}
