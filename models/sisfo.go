package models

type Sisfo struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password []byte `json:"-"`
	Email    string `json:"email"`
}
