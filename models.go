package models

type User struct {
	Id_user  uint   `json:"id_user" gorm:"primaryKey;autoIncrement:true"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
