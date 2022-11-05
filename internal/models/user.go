package models

// User holds all relevant userinformation
type User struct {
	Uid      string `json:"uid,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
