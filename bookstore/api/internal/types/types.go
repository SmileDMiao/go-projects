// Code generated by goctl. DO NOT EDIT.
package types

type CreateUserReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserResp struct {
	Ok bool `json:"ok"`
}
