package objects

type UpdateProfile struct{
	Name *string `json:"name"`
	Designation *string `json:"designation"`
	DOB *string `json:"DOB"`
	NickName *string `json:"nickName"`
}

