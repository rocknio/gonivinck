// Code generated by goctl. DO NOT EDIT.
package types

type RegisterRequest struct {
	Name     string `json:"name"`
	Gender   int64  `json:"gender"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type UserInfo struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Gender int64  `json:"gender"`
	Mobile string `json:"mobile"`
}

type LoginRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
}
