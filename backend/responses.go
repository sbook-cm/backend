package backend

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

type SigninResponse struct {
	UserResponse

	Sessionid string `json:"sessionid"`
}

type EventsResponse struct {
	_Response

	Events []Event `json:"events"`
}
