package backend

import "github.com/gorilla/mux"

var router *mux.Router
var FRONTEND string = "http://localhost:5173"

func Route() *mux.Router {
	router = mux.NewRouter()
	router.HandleFunc("/users.json", do_users)
	router.HandleFunc("/signin", do_signin)
	router.HandleFunc("/user/{username}.json", do_username)
	router.HandleFunc("/user/{userid}/profile", do_username_profile)
	router.HandleFunc("/{sessionid}/user.json", do_current_user)
	router.HandleFunc("/{sessionid}/events.json", do_latest_events)
	return router
}
