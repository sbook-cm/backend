package backend

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserContact struct {
	Email  string `json:"email"`
	Number int    `json:"number"`
}
type UserSecurity struct {
	Password string `json:"password"`
	Totp     string `json:"totp"`
}
type User struct {
	Name     string             `json:"name"`
	Username string             `json:"username"`
	Role     string             `json:"role"`
	Contact  UserContact        `json:"contact"`
	Tags     []string           `json:"tags"`
	Bio      string             `json:"bio"`
	Security UserSecurity       `json:"security"`
	Socials  []string           `json:"socials"`
	ID       primitive.ObjectID `json:"_id"`
}
