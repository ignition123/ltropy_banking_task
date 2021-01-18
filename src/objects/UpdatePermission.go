package objects

type UpdateKYCStruct struct{
	ApproveContact *bool `json:"approveContact"`
	KycComplete    *bool `json:"kycComplete"`
}

type UpdateBrokerStruct struct{
	AccountBlocked *bool `json:"account_blocked"`
	UpdateProfile  *bool `json:"update_profile"`
	FreezeAccount *bool `json:"freeze_account"`
	CreateNewTradingId *bool `json:"create_new_trading_id"`
	ResetBrokerPassword *bool `json:"reset_broker_password"`
}

type UpdateExchangePushStruct struct{
	Messages *bool `json:"messages"`
}

type UpdatePermissionsStruct struct{
	CreateAdmin *bool `json:"createAdmin"`
	KYC *UpdateKYCStruct `json:"kyc"`
	Broker *UpdateBrokerStruct `json:"brokers"`
	ExchangePush *UpdateExchangePushStruct `json:"exchange_push"`
}

type UpdatePermission struct{
	Email *string `json:"email"`
	SuperAdmin *bool `json:"superAdmin"`
	Permissions *UpdatePermissionsStruct
}