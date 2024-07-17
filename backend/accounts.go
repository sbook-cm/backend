package backend

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database

func getAllUsers() []User {
	var users []User
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

func getUserByUsername(userName string) (User, bool) {
	users_collection := db.Collection("users")
	var user User
	err := users_collection.FindOne(context.TODO(), bson.M{
		"username": userName,
	}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, true
		} else {
			log.Fatal(err)
		}
	}
	return user, false
}

func userFromLogin(email string, password string) (User, bool) {
	users_collection := db.Collection("users")
	var user User
	cursor, err := users_collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return user, true
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		//cursor.Decode(&user)
		cursor.Decode(&user)
		fmt.Println(user)
		if user.Security.Password == password && user.Contact.Email == email {
			return user, false
		}
	}
	return user, true
}

func SetDatabase(db2 *mongo.Database) {
	db = db2
}
