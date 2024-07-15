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
		url = "mongodb://mongo:PadssscQGFKrBYnwYnSkaLElshJgSUFM@monorail.proxy.rlwy.net:36478"
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func main() {
	client := connectDatabase()
	backend.SetDatabase(client.Database("test"))
	defer client.Disconnect(context.TODO())

	// r.HandleFunc("/users/{userid}", func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	userID := vars["userid"]
	// 	fmt.Fprintf(w, "User ID: %s", userID)
	// })

	http.Handle("/", backend.Route())

	log.Println("Server is running on http://localhost:8765")
	log.Fatal(http.ListenAndServe(":8765", nil))
}
