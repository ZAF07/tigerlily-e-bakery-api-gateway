package config

import (
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Inventories struct {
	Inventories []*rpc.Sku
}

func InitInventoryConfig() (in *Inventories) {
	log := logger.NewLogger()
	inventoryItems := &Inventories{}
	viper.AddConfigPath("./")
	viper.SetConfigName("inventory")
	viper.SetConfigType("json")
	viper.ReadInConfig()

	if err := viper.Unmarshal(inventoryItems); err != nil {
		log.ErrorLogger.Println("[CONFIG] Error unmarshaling data from JSON file : ", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.InfoLogger.Println("[CONFIG] Config has changed: ", e.Name)
		if err := viper.Unmarshal(inventoryItems); err != nil {
			log.ErrorLogger.Panicf("[CONFIG] Error unmarshaling onchange : %+v\n", err)
		}
	})

	return inventoryItems
}
