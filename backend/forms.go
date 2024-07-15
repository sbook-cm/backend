package backend

import "github.com/gorilla/schema"

type LoginForm struct {
	Email    string `schema:"email,required"`
	Password string `schema:"password,required"` // default:21 ; default ama|st
}

var decoder = schema.NewDecoder()
