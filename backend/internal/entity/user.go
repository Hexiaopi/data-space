package entity

type User struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	PassWord   string `json:"password"`
	Avatar     string `json:"avatar"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	State      uint8  `json:"state"`
	Roles      []Role `json:"roles"`
}
