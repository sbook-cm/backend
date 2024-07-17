package backend

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type EventParams = map[string]interface{}

type Event struct {
	Eventtype string             `json:"eventtype" bson:"eventtype"`
	Time      time.Time          `json:"time" bson:"time"`
	Params    EventParams        `json:"params" bson:"params"`
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
}

func getLatestEventsForUser(user User, number int) ([]Event, error) {
	var events []Event

	collection := db.Collection("events")
	cursor, err := collection.Find(
		context.Background(),
		bson.M{},
		// options.Find().SetSort(bson.M{
		// 	"$natural": -1,
		// }),
	)
	if err != nil {
		fmt.Println("error:", err)
		return events, err
	}
	defer cursor.Close(context.Background())
	for i := 0; cursor.Next(context.TODO()) && i < number; i++ {
		var event Event
		cursor.Decode(&event)
		events = append(events, event)
	}
	fmt.Println(number)
	return events, err
}

func createEvent(eventtype string, params EventParams) (Event, error) {
	fmt.Println("  create event", eventtype)
	events := db.Collection("events")
	event := Event{
		Eventtype: eventtype,
		Time:      time.Now(),
		Params:    params,
	}
	obj, err := events.InsertOne(context.Background(), event)
	if err != nil {
		fmt.Println("error:", err)
		return event, err
	}
	fmt.Println(obj)
	return event, err
}
