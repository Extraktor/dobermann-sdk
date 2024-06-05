package models

type Subscription struct {
	Customer             string  `json:"customer" bson:"customer"`
	BillingType          string  `json:"billingType" bson:"billingType"`
	Value                float64 `json:"value" bson:"value"`
	Cycle                string  `json:"cycle" bson:"cycle"`
	RemoteIp             string  `json:"remoteIp" bson:"remoteIp"`
	Description          string  `json:"description" bson:"description"`
	CreditCard           Card    `json:"creditCard" bson:"creditCard"`
	CreditCardHolderInfo Holder  `json:"creditCardHolderInfo" bson:"creditCardHolderInfo"`
}
