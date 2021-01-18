package objects

type GetAllBrokerIds struct{
	Id *string `json:"id"`
}

type TradingIdList struct{
	Status bool `json:"status"`
	TradingIds []map[string]interface{} `json:"tradingIds"`
}