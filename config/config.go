package config

import (
	"fmt"

	"github.com/dollarsignteam/go-logger"
	"github.com/dollarsignteam/go-utils"
	"github.com/spf13/viper"
)

var log *logger.Logger

func init() {
	log = logger.NewLogger(logger.LoggerOptions{
		Name:       utils.PackageName(),
		HideCaller: true,
	})
}

func AutoReadConfig(cfgFile string) {
	configFile := "local.env"
	if cfgFile != "" {
		configFile = cfgFile
	}
	if err := ReadConfig(configFile); err == nil {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	} else if cfgFile != "" {
		log.Fatalf("Config file: %s, %s", viper.ConfigFileUsed(), err)
	}
}

func AutoLoadConfig(config interface{}) {
	if err := LoadConfig(config); err != nil {
		log.Fatal(err.Error())
	}
	log.Infof("Config loaded: %+v", config)
}

func ReadConfig(cfgFile string) error {
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}

func LoadConfig(config interface{}) error {
	if err := viper.Unmarshal(config); err != nil {
		return fmt.Errorf("unmarshal config: %s", err)
	}
	return utils.Validate.Struct(config)
}
