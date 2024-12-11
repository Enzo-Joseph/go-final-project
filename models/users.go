package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int
	Name     string
	Username string
	Password string
	Role     string
}

func (User) TableName() string {
	return "users"
}

func GetStudents() ([]User, error) {
	var users []User

	res := DB.Table("users").
		Select("users.id, users.name, users.username, users.role").
		Where("users.role=?", "student").
		Find(&users)

	return users, res.Error
}

func GetLecturers() ([]User, error) {
	var users []User

	res := DB.Table("users").
		Select("users.id, users.name, users.username, users.role").
		Where("users.role=?", "lecturer").
		Find(&users)

	return users, res.Error
}

func GetLecturerByID(id string) (User, error) {
	var user User

	res := DB.Table("users").
		Select("users.id, users.name, users.username, users.role").
		Where("users.id=?", id).
		Find(&user)

	return user, res.Error
}

func CheckAuth(username string, password string) (User, error) {
	var user User

	res := DB.Table("users").
		Select("users.id, users.name, users.username, users.role").
		Where("users.username=?", username).
		Where("users.password=?", password).
		Find(&user)

	return user, res.Error
}

func CheckUsername(username string) (User, error) {
	var user User

	res := DB.Table("users").
		Select("users.id").
		Where("users.username=?", username).
		Find(&user)

	return user, res.Error
}

func CreateUser(name string, username string, password string, role string) error {
	user := User{
		Name:     name,
		Username: username,
		Password: password,
		Role:     role,
	}

	return DB.Create(&user).Error
}
