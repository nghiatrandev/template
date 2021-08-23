package dataservice

import "gorm.io/gorm"

type User struct {
	ID       uint64 `gorm:"id"`
	UserName string `gorm:"user_name"`
	FullName string `gorm:"fule_name"`
	Email    string `gorm:"email"`
	Password string `gorm:"password"`
	UsernameFilter string `gorm:"username_filter"`
	FullnameFilter string `gorm:"username_state"`
}

type UserDTO struct {
	UserName string `json:"user_name"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User)FromDTO(dto *UserDTO) {
	u.UserName = dto.UserName
	u.FullName = dto.FullName
	u.Email = dto.Email
	u.Password = dto.Password
}

func (u *User) Create (db *gorm.DB) error {
	return db.Create(&u).Error
}

func GetList (db *gorm.DB, username, fullname string) error {

	record := db.Model(User{})
	where := ""
	if username != "" {
		record = record.Where( "username ilike '%'"+username+"'%'")
	}

	if fullname != ""





	return db.
}