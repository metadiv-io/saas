package types

type Consumption struct {
	SubscriptionUUID string  `json:"subscription_uuid"`
	UserUUID         string  `json:"user_uuid"`
	Credit           float64 `json:"credit"`
}
