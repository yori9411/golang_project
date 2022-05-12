package database

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

//Background返回一個非空的Context。 它永遠不會被取消，沒有值，也沒有期限。
//它通常在main函式，初始化和測試時使用，並用作傳入請求的頂級上下文。
var ctx = context.Background()
var Rdbconnect *gorm.DB

func REDISDB() {
	Rdbconnect := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := Rdbconnect.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("連線redis出錯，錯誤資訊：%v", err)
	}
	fmt.Println("成功連線redis")
}
