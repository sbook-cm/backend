package backend

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
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
	ID       primitive.ObjectID `json:"id" bson:"_id"`
}

type Session struct {
	Userid    string                 `json:"userid" bson:"userid"`
	Sessionid string                 `json:"sessionid" bson:"sessionid"`
	Time      time.Time              `json:"time" bson:"time"`
	Data      map[string]interface{} `json:"data" bson:"data"`
	Flashes   []FlashMessage         `json:"flashes" bson:"flashes"`
	ID        primitive.ObjectID     `json:"id" bson:"_id"`
}

// {
// 	"name": "ken-morel",
// 	"username": "ken-morel",
// 	"role": "dev",
// 	"contact": {
// 		"email": "engonken8@gmail.com",
// 		"number": 237676573888,
// 	},
// 	"tags": ["dev"],
// 	"bio": "A great dev",
// 	"security": {
// 		"password": "amemimy",
// 		"totp": "",
// 	},
// 	"socials": ["x/com/ken-morel"],
// 	"friends": [],
// 	"following": [],
// }
// {
// 	"userid": "66aecf5d327f1f0012e17b11",
// 	"sessionid": "ken-morel86e7909f",
// 	"data": {},
// 	"flashes": [],
// }
