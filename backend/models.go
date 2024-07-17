package backend

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserContact struct {
	Email  string `json:"email" bson:"email"`
	Number int    `json:"number" bson:"number"`
}
type UserSecurity struct {
	Password string `json:"password" bson:"password"`
	Totp     string `json:"totp" bson:"totp"`
}
type User struct {
	Name     string             `json:"name" bson:"name"`
	Username string             `json:"username" bson:"username"`
	Role     string             `json:"role" bson:"role"`
	Contact  UserContact        `json:"contact" bson:"contact"`
	Tags     []string           `json:"tags" bson:"tags"`
	Bio      string             `json:"bio" bson:"bio"`
	Security UserSecurity       `json:"security" bson:"security"`
	Socials  []string           `json:"socials" bson:"socials"`
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
}
