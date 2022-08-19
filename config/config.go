package config

import (
	"fmt"

	"github.com/Tiger-Coders/tigerlily-bff/internal/pkg/logger"
	"github.com/Tiger-Coders/tigerlily-inventories/api/rpc"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Inventories          []*rpc.Sku
	PaymentServicePort   string
	InventoryServicePort string
}

var inventoryItems = &AppConfig{}

func InitAppConfig() (in *AppConfig) {
	log := logger.NewLogger()
	viper.SetConfigFile("./config.yml")
	viper.ReadInConfig()

	if err := viper.Unmarshal(inventoryItems); err != nil {
		log.ErrorLogger.Println("[CONFIG] Error unmarshaling data from JSON file : ", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.InfoLogger.Println("[CONFIG] Config has changed: ", e.Name)
		if err := viper.Unmarshal(inventoryItems); err != nil {
			log.ErrorLogger.Panicf("[CONFIG] Error unmarshaling on change : %+v\n", err)
		}
	})
	initInventoryConfig()
	return inventoryItems
}

func initInventoryConfig() {
	log := logger.NewLogger()
	viper.SetConfigFile("./inventory.yml")
	viper.ReadInConfig()

	err := viper.Unmarshal(inventoryItems)
	if err != nil {
		log.ErrorLogger.Fatalf("Error reading second file : %+v", err)
		fmt.Println(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.InfoLogger.Println("[CONFIG] Inventory Config has changed: ", e.Name)
		if err := viper.Unmarshal(inventoryItems); err != nil {
			log.ErrorLogger.Panicf("[CONFIG] Error unmarshaling Inventory Config on change : %+v\n", err)
		}
	})
}
