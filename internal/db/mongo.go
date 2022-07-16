package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
💡
 NOT USED BUT COULD BE USEFUL TO STORE ALL APPS CONFIG IN MONGODB AS A DOCUMENT AND UPON APP START-UP, WE UNMARSHAL CONFIGURATIONS INTO CONFIG STRUCTS FOR THE APP TO READ
*/
func InitMongoDB() (client *mongo.Client, err error) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI("mongodb+srv://tigerguy:<password>@tigerlily-cluster.x7ahvoi.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return
}
