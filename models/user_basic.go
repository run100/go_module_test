package models

import (
	"fmt"
	"github.com/run100/go_module_test/utils"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Avatar        string //头像
	Identity      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)

	//for _, row := range data {
	//	fmt.Printf("row %v\n", row)
	//}

	return data
}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func FindUserByNameAndPwd(name string, password string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ? and pass_word=?", name, password).First(&user)

	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.MD5Encode(str)
	utils.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp)
	return user
}
