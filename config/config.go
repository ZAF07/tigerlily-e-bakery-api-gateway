package config

import (
	"github.com/Tiger-Coders/tigerlily-bff/internal/pkg/logger"
	"github.com/Tiger-Coders/tigerlily-inventories/api/rpc"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type AppConfig struct {
	// Inventories          []*rpc.Sku `mapstructure:"inventories"`
	InventoryConfig      *InventoryConfig `mapstructure:"inventory_config"`
	PaymentServicePort   string           `mapstructure:"payment_service_port"`
	InventoryServicePort string           `mapstructure:"inventory_service_port"`
	ServicePort          string           `mapstructure:"service_port"`
}

type InventoryConfig struct {
	FilePath    string     `mapstructure:"inventory_file_path"`
	Inventories []*rpc.Sku `mapstructure:"inventories"`
}

func InitAppConfig() (in *AppConfig) {
	logger := logger.NewLogger()

	var appConfiguration = &AppConfig{}

	loadFromConfigFile("./config.yml", appConfiguration, *logger)
	loadFromConfigFile(appConfiguration.InventoryConfig.FilePath, appConfiguration, *logger)

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
