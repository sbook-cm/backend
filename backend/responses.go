package backend

type _Response struct {
	Ok     bool `json:"ok"`
	Status int  `json:"status"`
}

type _SResponse struct {
	Ok        bool   `json:"ok"`
	Status    int    `json:"status"`
	Sessionid string `json:"sessionid"`
}

type UsersResponse struct {
	_SResponse

	Users []User `json:"users"`
}
type UserResponse struct {
	_SResponse

	User User `json:"user"`
}

type SigninResponse = UserResponse

type EventsResponse struct {
	_SResponse

	Events []Event `json:"events"`
}
