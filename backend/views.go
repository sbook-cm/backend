package backend

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func do_users(w http.ResponseWriter, r *http.Request) {
	// Insert a document
	//_, err = collection.InsertOne(context.Background(), Person{"Alice", 30})
	//if err != nil {
	//	log.Fatal(err)
	//}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", FRONTEND)

	err := json.NewEncoder(w).Encode(UsersResponse{
		_Response: _Response{
			Ok:     true,
			Status: 200,
		},
		Users: getAllUsers(),
	})
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func do_current_user(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Asking for current user")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", FRONTEND)

	var users []User
	var user User
	var herr bool
	session, err := getSession(mux.Vars(r)["sessionid"])
	if err != nil {
		json.NewEncoder(w).Encode(UsersResponse{
			_Response: _Response{
				Ok:     false,
				Status: 402,
			},
			Users: users,
		})
		return
	}
	user, herr = getUserByUsername(session.Username)
	if herr {
		json.NewEncoder(w).Encode(UsersResponse{
			_Response: _Response{
				Ok:     false,
				Status: 404,
			},
			Users: users,
		})
		return
	}
	json.NewEncoder(w).Encode(UserResponse{
		_Response: _Response{
			Ok:     true,
			Status: 200,
		},
		User: user,
	})
	fmt.Println(user)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func do_signin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signin in")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", FRONTEND)
	var users []User
	if err := r.ParseForm(); err != nil {
		json.NewEncoder(w).Encode(UsersResponse{
			_Response: _Response{
				Ok:     false,
				Status: 403,
			},
			Users: users,
		})
		return
	}
	var form LoginForm
	err := decoder.Decode(&form, r.Form)
	if err != nil {
		json.NewEncoder(w).Encode(UsersResponse{
			_Response: _Response{
				Ok:     false,
				Status: 403,
			},
			Users: users,
		})
		return
	}
	user, herr := userFromLogin(form.Email, form.Password)
	fmt.Println(" sign in as", user)
	if herr {
		json.NewEncoder(w).Encode(UsersResponse{
			_Response: _Response{
				Ok:     false,
				Status: 404,
			},
			Users: users,
		})
		return
	}
	session := createSession(user.Username)
	if err := saveSession(session); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(SigninResponse{
		UserResponse: UserResponse{
			_Response: _Response{
				Ok:     true,
				Status: 200,
			},
			User: user,
		},
		Sessionid: session.Sessionid,
	})
}

func do_username(w http.ResponseWriter, r *http.Request) {
	userName := mux.Vars(r)["username"]
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", FRONTEND)

	var users []User
	var user User
	var herr bool
	user, herr = getUserByUsername(userName)
	fmt.Println("got:", user)
	if herr {
		json.NewEncoder(w).Encode(UsersResponse{
			_Response: _Response{
				Ok:     false,
				Status: 404,
			},
			Users: users,
		})
		return
	}
	err := json.NewEncoder(w).Encode(UserResponse{
		_Response: _Response{
			Ok:     true,
			Status: 200,
		},
		User: user,
	})
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func do_latest_events(w http.ResponseWriter, r *http.Request) {
	var user User
	var herr bool
	var form LatestEventsForm
	fmt.Println("Asking for latest events user")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", FRONTEND)
	if err := r.ParseForm(); err != nil {
		json.NewEncoder(w).Encode(EventsResponse{
			_Response: _Response{
				Ok:     false,
				Status: 403,
			},
			Events: []Event{},
		})
		return
	}
	session, err := getSession(mux.Vars(r)["sessionid"])
	if err != nil {
		json.NewEncoder(w).Encode(EventsResponse{
			_Response: _Response{
				Ok:     false,
				Status: 402,
			},
			Events: []Event{},
		})
		return
	}
	user, herr = getUserByUsername(session.Username)
	if herr {
		json.NewEncoder(w).Encode(EventsResponse{
			_Response: _Response{
				Ok:     false,
				Status: 404,
			},
			Events: []Event{},
		})
		return
	}
	err = decoder.Decode(&form, r.Form)
	if err != nil {
		json.NewEncoder(w).Encode(EventsResponse{
			_Response: _Response{
				Ok:     false,
				Status: 403,
			},
			Events: []Event{},
		})
		return
	}
	events, eerr := getLatestEventsForUser(user, form.Number)
	if eerr != nil {
		json.NewEncoder(w).Encode(EventsResponse{
			_Response: _Response{
				Ok:     false,
				Status: 500,
			},
			Events: []Event{},
		})
		return
	}
	json.NewEncoder(w).Encode(EventsResponse{
		_Response: _Response{
			Ok:     true,
			Status: 200,
		},
		Events: events,
	})
}

func do_username_profile(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userid"]
	if isValidHex(userId) {
		http.ServeFile(w, r, "./profile/user/"+userId+".png")
	} else {
		http.Error(w, "Wrong user id", http.StatusInternalServerError)
	}
}

func isValidHex(s string) bool {
	if len(s) != 24 {
		return false
	}
	for _, c := range s {
		if !(c >= '0' && c <= '9') && !(c >= 'a' && c <= 'f') {
			return false
		}
	}
	return true
}
