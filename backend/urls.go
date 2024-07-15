package backend

import "github.com/gorilla/mux"

func Route() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users.json", do_users)
	router.HandleFunc("/user.json", do_user)
	router.HandleFunc("/signin", do_signin)
	return router
}
