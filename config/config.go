package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	config Config
)

type Config struct {
	App AppConfig `yaml:"app"`
}

type AppConfig struct {
	Port int    `yaml:"port"`
	Echo string `yaml:"echo"`
}

func InitConfig(cfgPath string) error {
	if len(cfgPath) == 0 {
		fmt.Println("No config file provided, using default settings")
		defaultSettings()
		return nil
	}

	fmt.Printf("Parse config file: %v\n", cfgPath)
	return parseConfig(cfgPath)
}

func defaultSettings() {
	config.App.Port = 8080
	config.App.Echo = "hello"
}

func parseConfig(cfgPath string) error {
	viper.SetConfigFile(cfgPath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.Unmarshal(&config)
}

func GetConfig() *Config {
	return &config
}
