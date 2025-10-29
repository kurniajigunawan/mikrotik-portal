package model

type ResetSessionRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
