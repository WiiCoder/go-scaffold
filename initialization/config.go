package initialization

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/wiiCoder/go-scaffold/config"
	"github.com/wiiCoder/go-scaffold/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func Config() {
	debug := GetEnvInfo("DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("resources/%s-dev.yml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("resources/%s-pro.yml", configFilePrefix)
	}
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	var config config.ServerConfig
	if err := v.Unmarshal(&config); err != nil {
		panic(err)
	}
	global.Config = config
}
