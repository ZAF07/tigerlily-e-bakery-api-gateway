package config

import (
	"github.com/Tiger-Coders/tigerlily-bff/internal/pkg/logger"
	"github.com/Tiger-Coders/tigerlily-inventories/api/rpc"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Inventories          []*rpc.Sku
	PaymentServicePort   string
	InventoryServicePort string
	ServicePort          string
}

type TestConfig struct {
	Name string
	Age  int
}

func InitAppConfig() (in *AppConfig) {
	logger := logger.NewLogger()

	var appConfiguration = &AppConfig{}

	loadFromConfigFile("./config.yml", appConfiguration, *logger)
	loadFromConfigFile("./inventory.yml", appConfiguration, *logger)

	return appConfiguration
}

func loadFromConfigFile(filepath string, config *AppConfig, logger logger.Logger) {
	cfgLoader := viper.New()
	cfgLoader.SetConfigFile(filepath)
	cfgLoader.ReadInConfig()
	cfgLoader.Unmarshal(config)

	cfgLoader.WatchConfig()
	cfgLoader.OnConfigChange(func(e fsnotify.Event) {
		logger.InfoLogger.Printf("[CONFIG] %+v has changed. FilePath: : %s", e.Name, filepath)
		if err := viper.Unmarshal(config); err != nil {
			logger.ErrorLogger.Panicf("[CONFIG] Error unmarshaling %s Config on change : %+v\n", filepath, err)
		}
	})
}
