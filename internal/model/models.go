package model

type User struct {
	ID       int    `json:"userid"`
	Name     string `json:"username"`
	Location string `json:"userlocation"`
	Age      int    `json:"userage"`
}
