package model

type User struct {
	ID       uint   `json:"id" form:"id" gorm:"primary_key;autoIncrement"`
	Email    string `json:"email" form:"email" gorm:"type:varchar(255)"`
	Password string `json:"password" form:"password" gorm:"type:varchar(255)"`
	Role     string `json:"role" form:"role" gorm:"type:varchar(255)"`
}
