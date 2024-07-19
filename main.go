package main

import (
	"log"
	"os"

	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/sbook-cm/backend/backend"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func connectDatabase() *mongo.Client {
	var url string
	if os.Getenv("ENV") == "dev" {
		url = "mongodb://localhost:27017"
	} else {
		url = "mongodb://mongo:PadssscQGFKrBYnwYnSkaLElshJgSUFM@monorail.proxy.rlwy.net:36478"cd
		backend.FRONTEND = "https://sbook-cm.web.app"
	}
	backend.FRONTEND = "http://localhost:5173"
	url = "mongodb+srv://ken-morel:amemimy114865009@sbook.cxtildr.mongodb.net/?retryWrites=true&w=majority&appName=sbook"
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func main() {
	client := connectDatabase()
	backend.SetDatabase(client.Database("test"))
	defer client.Disconnect(context.TODO())

	http.Handle("/", backend.Route())

	log.Println("Server is running on http://localhost:1234")
	log.Fatal(http.ListenAndServe(":1234", nil))
}
