package domain

// User represent entity of the user on the DB
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Passhash string `json:"passhash"`
	Email    string `json:"email"`
	FullName string `json:"fullname"`
}
