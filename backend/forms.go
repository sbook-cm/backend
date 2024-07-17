package backend

import "github.com/gorilla/schema"

type LoginForm struct {
	Email    string `schema:"email,required"`
	Password string `schema:"password,required"` // default:21 ; default ama|st
}

type LatestEventsForm struct {
	Number int `json:"number"`
}

var decoder = schema.NewDecoder()
