package models

type User struct {
	UserId        int    `json:"id" db:"id"`
	UserName      string `json:"name" db:"name"`
	UserAge       int    `json:"age" db:"age"`
	UserIsRegular bool   `json:"isRegular" db:"isRegular"`
	Password      string `json:"password" db:"password"`
}
