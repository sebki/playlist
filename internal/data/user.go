package data

// User holds logindata for type Person
// TODO: Maybe integrate in Person type?
type User struct {
	Uid      string `json:"uid,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
