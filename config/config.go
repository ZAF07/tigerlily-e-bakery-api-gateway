package config

import (
	"fmt"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Inventories          []*rpc.Sku
	PaymentServicePort   string
	InventoryServicePort string
}

func InitInventoryConfig() (in *AppConfig) {
	log := logger.NewLogger()
	inventoryItems := &AppConfig{}
	viper.AddConfigPath("./")
	viper.SetConfigName("inventory")
	viper.SetConfigType("json")
	viper.ReadInConfig()
	a := viper.Get("Payment_service_port")
	fmt.Println("HERE _ >", a)
	if err := viper.Unmarshal(inventoryItems); err != nil {
		log.ErrorLogger.Println("[CONFIG] Error unmarshaling data from JSON file : ", err)
	}
	fmt.Printf("CONFIG : %+v", inventoryItems)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.InfoLogger.Println("[CONFIG] Config has changed: ", e.Name)
		if err := viper.Unmarshal(inventoryItems); err != nil {
			log.ErrorLogger.Panicf("[CONFIG] Error unmarshaling onchange : %+v\n", err)
		}
	})

	return inventoryItems
}
