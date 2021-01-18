package objects

type ErrorResponse struct{
	Status bool `json:"status"`
	Msg string `json:"msg"`
}

type SuccessResponse struct{
	Status bool `json:"status"`
	Msg string `json:"msg"`
}