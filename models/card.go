package models

type Card struct {
	HolderName  string `json:"holderName" bson:"holderName"`
	Number      string `json:"number" bson:"number"`
	ExpiryMonth string `json:"expiryMonth" bson:"expiryMonth"`
	ExpiryYear  string `json:"expiryYear" bson:"expiryYear"`
	Ccv         string `json:"ccv" bson:"ccv"`
}
