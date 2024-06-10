package entity

type User struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	State      uint8  `json:"state"`
}

type UserInfo struct {
	User
	Roles       []Role `json:"roles"`
	CurrentRole Role   `json:"current_role"`
}
