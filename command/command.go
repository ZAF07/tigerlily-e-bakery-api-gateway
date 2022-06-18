package command

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"

	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/config"
	"github.com/ZAF07/tigerlily-e-bakery-api-gateway/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
)

var InjectInventoriesCmd = &cobra.Command{
	Use:   "inventories",
	Short: "Init the app passing the inventory items",
	Long:  "CLI command to initialise the app with names of the inventory items you would like to populate the app upon start",
	Run:   InjectInventoryItems,
}

func InjectInventoryItems(cmd *cobra.Command, args []string) {
	log := logger.NewLogger()
	if len(args) < 1 {
		log.WarnLogger.Println("[COMMAND] No arguments entered in CLI for start-up")
		return
	}
	log.InfoLogger.Printf("[COMMAND] These are the args enterd in the cli : %+v\n", args)
	inventories := config.Inventories{}
	for _, v := range args {
		singleItem := &rpc.Sku{
			Name: v,
		}
		inventories.Inventories = append(inventories.Inventories, singleItem)
	}

	b, err := json.Marshal(inventories)
	if err != nil {
		log.ErrorLogger.Fatalf("[COMMAND] Error Marshaling inventories : %+v\n", err)
	}
	if osErr := os.WriteFile("inventory.json", b, 0644); osErr != nil {
		log.ErrorLogger.Fatalf("[COMMAND] Error writing inventories into JSON file : %+v\n", osErr)
	}
}
