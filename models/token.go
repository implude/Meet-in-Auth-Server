package models

type Token struct {
	TokenID uint64 `json:"id"`
	Expired bool   `json:"expire"`
}
