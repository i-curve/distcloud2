package model

import "gorm.io/gorm"

// User 用户数据表
type User struct {
	gorm.Model
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Privilege int    `json:"role"`
}

// RegisterUser 是注册用户函数
func RegisterUser(user *User) (bool, error) {
	err := db.Model(&User{}).Create(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// DeleteUser 是删除用户函数
func DeleteUser(id int) (bool, error) {
	err := db.Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// LoginUser 用户登录
func LoginUser(email, password string) int {
	var _user User
	db.Model(&User{}).Where("email = ?", email).Find(&_user)
	if password != _user.Password {
		return -1
	}
	return int(_user.ID)
}

// UpdateUser 是更新用户函数
func UpdateUser(user User) (bool, error) {
	err := db.Model(&user).Updates(user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// MsgUser 获取用户信息
func MsgUser(id int, username, email string) (User, bool) {
	var user User
	var err error
	if id != 0 {
		err = db.Model(&User{}).Where("id = ?", id).First(&user).Error
	} else if username != "" {
		err = db.Model(&User{}).Where("username = ?", username).First(&user).Error
	} else if email != "" {
		err = db.Model(&User{}).Where("email = ?", email).First(&user).Error
	} else {
		return user, false
	}

	if err != nil || user.ID <= 0 {
		return user, false
	}
	return user, true
}

// IsExist 验证用户否存在函数
func IsExist(email string) bool {
	var user User
	err := db.Model(&user).Where("email=?", email).First(&user).Error
	if err == nil && user.ID > 0 {
		return true
	}
	return false
}

// AddUser 是增加用户函数
func AddUser(email, password string) (bool, error) {
	_user := User{
		Email:     email,
		Password:  password,
		Privilege: 2,
	}
	err := db.Model(&User{}).Create(&_user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetUsers 获取全部用户信息
func GetUsers() ([]User, bool) {
	var users []User
	result := db.Select("id", "username", "email", "privilege").Find(&users)
	if result.Error != nil {
		return nil, false
	}
	return users, true
}
