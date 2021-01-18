package objects

type KYCStruct struct{
	ApproveContact *bool `json:"approveContact"`
	KycComplete    *bool `json:"kycComplete"`
}

type BrokerStruct struct{
	AccountBlocked *bool `json:"account_blocked"`
	UpdateProfile  *bool `json:"update_profile"`
	FreezeAccount *bool `json:"freeze_account"`
	CreateNewTradingId *bool `json:"create_new_trading_id"`
	ResetBrokerPassword *bool `json:"reset_broker_password"`
}

type ExchangePushStruct struct{
	Messages *bool `json:"messages"`
}

type PermissionsStruct struct{
	CreateAdmin *bool `json:"createAdmin"`
	KYC *KYCStruct `json:"kyc"`
	Broker *BrokerStruct `json:"brokers"`
	ExchangePush *ExchangePushStruct `json:"exchange_push"`
}

type Register struct{
	Email *string `json:"email"`
	Name *string `json:"name"`
	SuperAdmin *bool `json:"superAdmin"`
	Permissions *PermissionsStruct
}