package models

type User struct {
	Model
	Name string `gorm:"unique_index" json:"name"`
	Email string `gorm:"unique_index" json:"email"`
	Pwd string `json:"-"`
	Avatar string `json:"avatar"`
	Role int  `json:"role"`// 0代表管理员，1代表正常用户
}

func QueryUserByEmailAndPwd(email,pwd string) (user User,err error) {
	return user,db.Where("email = ? and pwd = ?",email,pwd).Take(&user).Error
}

func QueryUserByName(name string )(user User,err error){
	return user,db.Where("name = ? ",name).Take(&user).Error
}

func QueryUserByEmail(email string )(user User,err error){
	return user,db.Where("email = ? ",email).Take(&user).Error
}

func SaveUser(user *User) error {
	return db.Save(user).Error
}