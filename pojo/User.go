package pojo

import (
	"time"
)

//定義結構體（User-->Table 表名默認為結構體名稱的複數users）
type User struct {
	Id       int       `json:"UserId"`
	Name     string    `json:"UserName"`
	Password string    `json:"UserPassword"`
	Email    string    `json:"UserEmail"`
	Birthday time.Time `json:"UserBirthday"`
}
