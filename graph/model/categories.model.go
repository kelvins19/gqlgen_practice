package model

type Categories struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description *string     `json:"description"`
	Products    []*Products `json:"products"`
}
