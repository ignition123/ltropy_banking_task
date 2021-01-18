package objects

type SessKYCStruct struct{
	ApproveContact *bool `json:"approveContact"`
	KycComplete    *bool `json:"kycComplete"`
}

type SessBrokerStruct struct{
	AccountBlocked *bool `json:"account_blocked"`
	UpdateProfile  *bool `json:"update_profile"`
}

type SessExchangePushStruct struct{
	Messages *bool `json:"messages"`
}

type SessPermissionsStruct struct{
	CreateAdmin *bool `json:"createAdmin"`
	KYC *KYCStruct `json:"kyc"`
	Broker *BrokerStruct `json:"brokers"`
	ExchangePush *ExchangePushStruct `json:"exchange_push"`
}

type Session struct{
	Email *string `json:"email"`
	Name *string `json:"name"`
	SuperAdmin *bool `json:"superAdmin"`
	Permissions *PermissionsStruct
}