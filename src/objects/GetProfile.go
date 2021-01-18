package objects

type ProfileResp struct{
	Status bool `json:"status"`
	Profile map[string]interface{} `json:"profile"`
}
