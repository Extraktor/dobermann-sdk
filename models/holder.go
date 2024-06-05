package models

type Holder struct {
	Name          string `json:"name" bson:"name"`
	Email         string `json:"email" bson:"email"`
	Document      string `json:"cpfCnpj" bson:"cpfCnpj"`
	PostalCode    string `json:"postalCode" bson:"postalCode"`
	AddressNumber string `json:"addressNumber" bson:"addressNumber"`
	Phone         string `json:"phone" bson:"phone"`
}
