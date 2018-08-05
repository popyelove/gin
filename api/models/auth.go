package models
import (
	db "Gin/common/databases/mysql"
)
type Auth struct {
	ID int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (auth *Auth)CheckAuth() bool {
	db.SqlDB.QueryRow("SELECT id FROM auth WHERE username=? and password=?",auth.Username,auth.Password).Scan(&auth.ID)
	if auth.ID > 0 {
		return true
	}
	return false
}
