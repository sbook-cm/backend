package backend

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database

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

func SetDatabase(db2 *mongo.Database) {
	db = db2
}
