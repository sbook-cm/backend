package backend

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func getUserByUserID(userId string) (User, error) {
	var user User
	users_collection := db.Collection("users")
	objectID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return user, err
	}
	err = users_collection.FindOne(context.TODO(), bson.M{
		"_id": objectID,
	}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, err
		} else {
			log.Fatal(err)
		}
	}
	return user, nil
}

func userFromEmail(email string) (User, error) {
	users_collection := db.Collection("users")
	var user User
	cursor, err := users_collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return user, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		//cursor.Decode(&user)
		cursor.Decode(&user)
		fmt.Println(user)
		if user.Contact.Email == email {
			return user, nil
		}
	}
	return user, mongo.ErrNoDocuments
}

func SetDatabase(db2 *mongo.Database) {
	db = db2
}
