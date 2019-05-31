package model

type User struct {
	Phone    string `json:"phone" bson:"Phone" form:"phone"`
	Password string `json:"password" bson:"Password" form:"password"`
}

func (*User) GetCollectionName() string {
	return "User"
}
