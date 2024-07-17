package backend

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type FlashMessage struct {
	Status  int8   `json:"status" bson:"status"`
	Message string `json:"message" bson:"message"`
}

type Session struct {
	Username  string                 `json:"username" bson:"username"`
	Sessionid string                 `json:"sessionid" bson:"sessionid"`
	Time      time.Time              `json:"time" bson:"time"`
	Data      map[string]interface{} `json:"data" bson:"data"`
	Flashes   []FlashMessage         `json:"flashes" bson:"flashes"`
}

func generateUniqueSessionKey(username string) string {
	var key string
	// for {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	key = username + ":" + hex.EncodeToString(b)
	// err := sessions.FindOne(context.Background(), bson.M{
	// 	"sessionid": key,
	// }).Decode(&sess)
	// if err != nil {
	return key
	// }
	// }
}

func createSession(username string) Session {
	sessions := db.Collection("sessions")
	var data map[string]interface{}
	session := Session{
		Username:  username,
		Sessionid: generateUniqueSessionKey(username),
		Time:      time.Now(),
		Data:      data,
	}
	_, err := sessions.InsertOne(context.Background(), session)
	if err != nil {
		log.Fatal(err)
	}
	createEvent("session-creation", EventParams{
		"sessionid": session.Sessionid,
		"username":  username,
	})
	return session
}

func getSession(sessionid string) (Session, error) {
	sessions := db.Collection("sessions")
	var session Session
	err := sessions.FindOne(context.Background(), bson.M{
		"sessionid": sessionid,
	}).Decode(&session)
	return session, err
}

func saveSession(session Session) error {
	sessions := db.Collection("sessions")

	_, err := sessions.UpdateOne(
		context.Background(),
		bson.M{
			"sessionid": session.Sessionid,
		},
		bson.M{
			"$set": bson.M{
				"data": session.Data,
			},
		},
	)

	return err
}
