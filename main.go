package main

import (
	"encoding/json"
	"log"

	"context"
	"fmt"
	"net/http"

	//"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"github.com/gorilla/schema"
	"go.mongodb.org/mongo-driver/mongo"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type _Response struct {
	Ok     bool `json:"ok"`
	Status int  `json:"status"`
}

type UsersResponse struct {
	_Response

	Users []User `json:"users"`
}
type UserResponse struct {
	_Response

	User User `json:"user"`
}

var store = sessions.NewCookieStore([]byte("unbilivabel sicrete ki of mayne"))

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
var db *mongo.Database

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func do_users(w http.ResponseWriter, r *http.Request) {
	// Insert a document
	//_, err = collection.InsertOne(context.Background(), Person{"Alice", 30})
	//if err != nil {
	//	log.Fatal(err)
	//}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

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

func do_user(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Asking for user")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

	var users []User
	var user User
	var herr bool
	session, e := store.Get(r, "auth")
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(session.Values)
	if userID, ok := session.Values["user_id"].(string); ok {
		user, herr = getUserByID(userID)
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
	} else {
		json.NewEncoder(w).Encode(UsersResponse{
			_Response: _Response{
				Ok:     false,
				Status: 403,
			},
			Users: users,
		})
		return
	}
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

var decoder = schema.NewDecoder()

type LoginForm struct {
	Email    string `schema:"email,required"`
	Password string `schema:"password,required"` // default:21 ; default ama|st
}

func do_signin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signin in")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	session, e := store.Get(r, "auth")
	if e != nil {
		log.Fatal(e)
	}
	var users []User
	if err := r.ParseForm(); err != nil {
		fmt.Println("  Error parsing form", err)
		json.NewEncoder(w).Encode(UsersResponse{
			_Response: _Response{
				Ok:     false,
				Status: 403,
			},
			Users: users,
		})
		return
	}
	fmt.Println("  parsed form", r.Form)
	var form LoginForm

	err := decoder.Decode(&form, r.Form)
	fmt.Println("  decoded form", form)
	if err != nil {
		fmt.Println("error decoding form")
		json.NewEncoder(w).Encode(UsersResponse{
			_Response: _Response{
				Ok:     false,
				Status: 403,
			},
			Users: users,
		})
		return
	}
	fmt.Println("  getting user", form.Email, form.Password)
	user, herr := userFromLogin(form.Email, form.Password)
	fmt.Println("  got user", user)
	if herr {
		fmt.Println("  No user matching query:", form.Email, form.Password)
		json.NewEncoder(w).Encode(UsersResponse{
			_Response: _Response{
				Ok:     false,
				Status: 404,
			},
			Users: users,
		})
		return
	}
	fmt.Println("go user, session:", session, session.Values)
	session.Values["user_id"] = user.ID.Hex()
	fmt.Println("  Set session user_id", user.ID, user.ID.Hex())
	err = sessions.Save(r, w)
	//store.Save(r, w, session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	fmt.Println(session.Values)
	json.NewEncoder(w).Encode(UserResponse{
		_Response: _Response{
			Ok:     true,
			Status: 200,
		},
		User: user,
	})
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func main() {
	client := connectDatabase()
	defer client.Disconnect(context.TODO())

	router := mux.NewRouter()

	router.HandleFunc("/users.json", do_users)
	router.HandleFunc("/user.json", do_user)
	router.HandleFunc("/signin", do_signin)
	// r.HandleFunc("/users/{userid}", func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	userID := vars["userid"]
	// 	fmt.Fprintf(w, "User ID: %s", userID)
	// })

	http.Handle("/", router)

	log.Println("Server is running on http://localhost:8765")
	log.Fatal(http.ListenAndServe(":8765", nil))
}
