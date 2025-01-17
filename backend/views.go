package backend

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// func do_users(w http.ResponseWriter, r *http.Request) {
// 	// Insert a document
// 	//_, err = collection.InsertOne(context.Background(), Person{"Alice", 30})
// 	//if err != nil {
// 	//	log.Fatal(err)
// 	//}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", FRONTEND)

// 	err := json.NewEncoder(w).Encode(UsersResponse{
// 		_SResponse: _SResponse{
// 			Ok:        true,
// 			Status:    200,
// 			Sessionid: "",
// 		},
// 		Users: getAllUsers(),
// 	})
// 	if err != nil {
// 		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
// 		return
// 	}
// }

// func do_current_user(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Asking for current user")
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", FRONTEND)

// 	var users []User
// 	var user User
// 	oldsessionid := mux.Vars(r)["sessionid"]
// 	session, err := getSession(oldsessionid)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(UsersResponse{
// 			_SResponse: _SResponse{
// 				Ok:        false,
// 				Status:    402,
// 				Sessionid: "",
// 			},
// 			Users: users,
// 		})
// 		return
// 	}
// 	newsessionid, _ := getNewerSessionID(session, oldsessionid)
// 	user, err = getUserByUserID(session.Userid)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(UsersResponse{
// 			_SResponse: _SResponse{
// 				Ok:        false,
// 				Status:    404,
// 				Sessionid: newsessionid,
// 			},
// 			Users: users,
// 		})
// 		return
// 	}
// 	json.NewEncoder(w).Encode(UserResponse{
// 		_SResponse: _SResponse{
// 			Ok:        true,
// 			Status:    200,
// 			Sessionid: newsessionid,
// 		},
// 		User: user,
// 	})
// 	fmt.Println(user)
// 	if err != nil {
// 		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
// 		return
// 	}
// }

// func do_signin(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("signin in")
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", FRONTEND)
// 	var users []User
// 	if err := r.ParseForm(); err != nil {
// 		json.NewEncoder(w).Encode(UsersResponse{
// 			_SResponse: _SResponse{
// 				Ok:        false,
// 				Status:    403,
// 				Sessionid: "",
// 			},
// 			Users: users,
// 		})
// 		return
// 	}
// 	var form LoginForm
// 	var user User
// 	err := decoder.Decode(&form, r.Form)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(UsersResponse{
// 			_SResponse: _SResponse{
// 				Ok:        false,
// 				Status:    403,
// 				Sessionid: "",
// 			},
// 			Users: users,
// 		})
// 		return
// 	}
// 	user, err = userFromLogin(form.Email, form.Password)
// 	fmt.Println(" sign in as", user)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(UsersResponse{
// 			_SResponse: _SResponse{
// 				Ok:        false,
// 				Status:    404,
// 				Sessionid: "",
// 			},
// 			Users: users,
// 		})
// 		return
// 	}
// 	session := createSession(user.Username)
// 	if err := saveSession(session); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(SigninResponse{
// 		_SResponse: _SResponse{
// 			Ok:        true,
// 			Status:    200,
// 			Sessionid: session.Sessionid,
// 		},
// 		User: user,
// 	})
// }

// func do_userid(w http.ResponseWriter, r *http.Request) {
// 	userid := mux.Vars(r)["userid"]
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", FRONTEND)

// 	var users []User
// 	user, err := getUserByUserID(userid)
// 	fmt.Println("got:", user)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(UsersResponse{
// 			_SResponse: _SResponse{
// 				Ok:        false,
// 				Status:    404,
// 				Sessionid: "",
// 			},
// 			Users: users,
// 		})
// 		return
// 	}
// 	err = json.NewEncoder(w).Encode(UserResponse{
// 		_SResponse: _SResponse{
// 			Ok:        true,
// 			Status:    200,
// 			Sessionid: "",
// 		},
// 		User: user,
// 	})
// 	if err != nil {
// 		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
// 		return
// 	}
// }

// func do_latest_events(w http.ResponseWriter, r *http.Request) {
// 	var user User
// 	var form LatestEventsForm
// 	fmt.Println("Asking for latest events user")
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", FRONTEND)
// 	if err := r.ParseForm(); err != nil {
// 		json.NewEncoder(w).Encode(EventsResponse{
// 			_SResponse: _SResponse{
// 				Ok:        false,
// 				Status:    403,
// 				Sessionid: "",
// 			},
// 			Events: []Event{},
// 		})
// 		return
// 	}
// 	oldsessionid := mux.Vars(r)["sessionid"]
// 	session, err := getSession(oldsessionid)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(EventsResponse{
// 			_SResponse: _SResponse{
// 				Ok:        false,
// 				Status:    402,
// 				Sessionid: "",
// 			},
// 			Events: []Event{},
// 		})
// 		return
// 	}
// 	newsessionid, _ := getNewerSessionID(session, oldsessionid)
// 	user, err = getUserByUserID(session.Userid)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(EventsResponse{
// 			_SResponse: _SResponse{
// 				Ok:        false,
// 				Status:    404,
// 				Sessionid: newsessionid,
// 			},
// 			Events: []Event{},
// 		})
// 		return
// 	}
// 	err = decoder.Decode(&form, r.Form)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(EventsResponse{
// 			_SResponse: _SResponse{
// 				Ok:        false,
// 				Status:    403,
// 				Sessionid: newsessionid,
// 			},
// 			Events: []Event{},
// 		})
// 		return
// 	}
// 	events, eerr := getLatestEventsForUser(user, form.Number)
// 	if eerr != nil {
// 		json.NewEncoder(w).Encode(EventsResponse{
// 			_SResponse: _SResponse{
// 				Ok:        false,
// 				Status:    500,
// 				Sessionid: newsessionid,
// 			},
// 			Events: []Event{},
// 		})
// 		return
// 	}
// 	json.NewEncoder(w).Encode(EventsResponse{
// 		_SResponse: _SResponse{
// 			Ok:        true,
// 			Status:    200,
// 			Sessionid: newsessionid,
// 		},
// 		Events: events,
// 	})
// }

// func do_userid_profile(w http.ResponseWriter, r *http.Request) {
// 	userId := mux.Vars(r)["userid"]
// 	if isValidHex(userId) {
// 		http.ServeFile(w, r, "./profile/user/"+userId+".png")
// 	} else {
// 		http.Error(w, "Wrong user id", http.StatusInternalServerError)
// 	}
// }

// func isValidHex(s string) bool {
// 	if len(s) != 24 {
// 		return false
// 	}
// 	for _, c := range s {
// 		if !(c >= '0' && c <= '9') && !(c >= 'a' && c <= 'f') {
// 			return false
// 		}
// 	}
// 	return true
// }

// func do_logout(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", FRONTEND)
// 	session, err := getSession(mux.Vars(r)["sessionid"])
// 	if err != nil {
// 		json.NewEncoder(w).Encode(_Response{
// 			Ok:     false,
// 			Status: 500,
// 		})
// 		return
// 	}
// 	err = deleteSession(session)
// 	if err != nil {
// 		http.Error(w, "Wrong user id", http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(_Response{
// 		Ok:     true,
// 		Status: 200,
// 	})
// }

func do_email(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	w.Header().Set("Content-Type", "application/json")

	user, err := userFromEmail(email)
	if err != nil {
		json.NewEncoder(w).Encode(UserResponse{
			_Response: _Response{
				Ok:     false,
				Status: 404,
			},
			User: user,
		})
		return
	}
	err = json.NewEncoder(w).Encode(UserResponse{
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

func do_session(w http.ResponseWriter, r *http.Request) {
	sessionid := mux.Vars(r)["sessionid"]
	w.Header().Set("Content-Type", "application/json")

	session, err := getSession(sessionid)
	if err != nil {
		json.NewEncoder(w).Encode(SessionResponse{
			_Response: _Response{
				Ok:     false,
				Status: 404,
			},
			Session: session,
			User:    User{},
		})
		return
	}
	user, err := getUserByUserID(session.Userid)
	fmt.Println(err)
	if err != nil {
		json.NewEncoder(w).Encode(SessionResponse{
			_Response: _Response{
				Ok:     false,
				Status: 500,
			},
			Session: session,
			User:    user,
		})
		return
	}
	err = json.NewEncoder(w).Encode(SessionResponse{
		_Response: _Response{
			Ok:     true,
			Status: 200,
		},
		Session: session,
		User:    user,
	})
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}
