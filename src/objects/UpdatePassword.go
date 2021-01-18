package objects

type UpdatePassword struct{
	Password *string `json:"password"`
	ConfirmPassword *string `json:"confirm_password"`
}
