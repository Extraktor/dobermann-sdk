package models

type Client struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Document string `json:"cpfCnpj" bson:"cpfCnpj"`
}
