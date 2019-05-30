package model

type User struct {
	Phone    string `json:"phone" bson:"Phone"`
	Password string `json:"password" bson:"Password"`
}

func (*User) GetCollectionName() string {
	return "User"
}
