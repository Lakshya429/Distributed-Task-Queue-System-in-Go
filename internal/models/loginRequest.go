package models 

type LoginRequest struct {
	Username string `json : string`
	Password string `json : string`
}