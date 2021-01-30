package user_service

import (
	"clouddist/model"
	"clouddist/pkg/util"
)

// User 服务用于调用底层gorm方法
type User struct {
	ID        int
	Username  string
	Password  string
	Email     string
	Privilege int
}

// Register 用户注册接口
func (a *User) Register() (bool, error) {
	password := util.EncodeMD5(a.Password)
	_user := model.User{
		Username:  a.Username,
		Password:  password,
		Privilege: 2,
	}
	return model.RegisterUser(&_user)
}

// DeleteUser 用户删除接口
func (a *User) DeleteUser() (bool, error) {
	status, err := model.DeleteUser(a.ID)
	return status, err
}

// UpdateUser 用户信息更新接口
func (a *User) UpdateUser() (bool, error) {
	password := ""
	if a.Password != "" {
		password = util.EncodeMD5(a.Password)
	}
	_user := model.User{
		Username: a.Username,
		Password: password,
		Email:    a.Email,
	}
	_user.ID = uint(a.ID)
	_, err := model.UpdateUser(_user)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetMsg 获取全部信息
func (a *User) GetMsg() bool {

	user, status := model.MsgUser(a.ID, a.Username, a.Email)

	if !status {
		return false
	}
	a.ID = int(user.ID)
	a.Username = user.Username
	a.Email = user.Email
	a.Password = user.Password
	a.Privilege = user.Privilege
	return true
}

// Check 验证用户身份
func (a *User) Check(password string) bool {
	password = util.EncodeMD5(password)
	return a.Password == password
}
