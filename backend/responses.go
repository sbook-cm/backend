package backend

type _Response struct {
	Ok     bool `json:"ok"`
	Status int  `json:"status"`
}

// type _SResponse struct {
// 	Ok        bool   `json:"ok"`
// 	Status    int    `json:"status"`
// 	Sessionid string `json:"sessionid"`
// }

// type UsersResponse struct {
// 	_SResponse

// 	Users []User `json:"users"`
// }

// type SigninResponse = UserResponse

// type EventsResponse struct {
// 	_SResponse

// 	Events []Event `json:"events"`
// }

type UserResponse struct {
	_Response

	User User `json:"user"`
}

type SessionResponse struct {
	_Response

	Session Session `json:"session"`
	User    User    `json:"user"`
}
