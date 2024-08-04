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

func generateUniqueSessionKey(userid string) string {
	var key string
	// for {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	key = userid + ":" + hex.EncodeToString(b)
	// err := sessions.FindOne(context.Background(), bson.M{
	// 	"sessionid": key,
	// }).Decode(&sess)
	// if err != nil {
	return key
	// }
	// }
}

func createSession(userid string) Session {
	sessions := db.Collection("sessions")
	var data map[string]interface{}
	sessionid := generateUniqueSessionKey(userid)
	session := Session{
		Userid:    userid,
		Sessionid: sessionid,
		Time:      time.Now(),
		Data:      data,
	}
	_, err := sessions.InsertOne(context.Background(), session)
	if err != nil {
		log.Fatal(err)
	}
	createEvent("session-creation", EventParams{
		"sessionid": session.Sessionid,
		"userid":    userid,
	})
	return session
}

func deleteSession(session Session) error {
	sessions := db.Collection("sessions")
	filter := bson.M{"_id": "myid"}
	_, err := sessions.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	createEvent("session-deletion", EventParams{
		"sessionid": session.Sessionid,
		"userid":    session.Userid,
	})
	return nil
}

func getSession(sessionid string) (Session, error) {
	sessions := db.Collection("sessions")
	var session Session
	err := sessions.FindOne(context.Background(), bson.M{
		"sessionid": sessionid,
	}).Decode(&session)
	if err != nil {
		err = sessions.FindOne(context.Background(), bson.M{
			"oldsessionid": sessionid,
		}).Decode(&session)
	}
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
			"$set": session,
		},
	)

	return err
}

func getNewerSessionID(session Session, oldid string) (string, error) {
	if session.Sessionid == oldid {
		session.Sessionid = generateUniqueSessionKey(session.Userid)
		err := saveSession(session)
		return session.Sessionid, err
	} else {
		return session.Sessionid, nil
	}
}
