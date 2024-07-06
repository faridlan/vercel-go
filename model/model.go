package model

type User struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

type Profile struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}
