package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("unbilivabel sicrete ki of mayne"))

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
