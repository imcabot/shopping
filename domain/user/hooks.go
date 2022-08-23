package user

import (
	"gorm.io/gorm"
	"shoping/utils/hash"
)

//保存用户之前回调，如果密码没有被加密加密密码和salt
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Salt == "" {
		//为salt创建一个随机字符串
		salt := hash.CreateSalt()
		//创建hash加密密码
		hashPassword, err := hash.Hashpassword(u.Password + salt)
		if err != nil {
			return err
		}
		u.Password = hashPassword
		u.Salt = salt
	}
	return
}
