package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserContact struct {
	Email  string `json:"email"`
	Number int    `json:"number"`
}
type UserSecurity struct {
	Password string `json:"password"`
	Totp     string `json:"totp"`
}
type User struct {
	Name     string             `json:"name"`
	Username string             `json:"username"`
	Role     string             `json:"role"`
	Contact  UserContact        `json:"contact"`
	Tags     []string           `json:"tags"`
	Bio      string             `json:"bio"`
	Security UserSecurity       `json:"security"`
	Socials  []string           `json:"socials"`
	ID       primitive.ObjectID `json:"_id"`
}

func getAllUsers() (users []User) {
	users_collection := db.Collection("users")

	cursor, err := users_collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}
	return users
}

func getUserByID(userID string) (User, bool) {
	var user User
	users_collection := db.Collection("users")
	objID, _ := primitive.ObjectIDFromHex(userID)

	cursor, err := users_collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		cursor.Decode(&user)
		if user.ID == objID {
			return user, false
		}
	}
	return user, true
}
func userFromLogin(email string, password string) (User, bool) {
	users_collection := db.Collection("users")
	var user User
	cursor, err := users_collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		cursor.Decode(&user)
		// fmt.Println("mapping:", user.Contact.Email, "to", email, user.Contact.Email == email)
		// fmt.Println("mapping:", user.Security.Password, "to", password, user.Security.Password == password)
		if (user.Contact.Email == email) && (user.Security.Password == password) {
			return user, false
		}
		// fmt.Println("match", (user.Contact.Email == email), (user.Security.Password == password), (user.Contact.Email == email) && (user.Security.Password == password))
	}
	return user, true
}
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
	db = client.Database("test")
	return client
}
