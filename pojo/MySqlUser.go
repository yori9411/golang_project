package pojo

import (
	"golangAPI/database"
	"log"
)

//GET
func FindAllUsers() []User {
	var users []User                //聲明模型結構體類型變user slice
	database.DBconnect.Find(&users) //select * from users;
	return users
}

func FindByUserId(userId string) User {
	var user User
	database.DBconnect.Where("id=?", userId).First(&user) //select * from user where userId="" limit =1;
	return user
}

//POST
func CreateUser(user User) User {
	database.DBconnect.Create(&user)
	return user
}

//Delete
func DeleteUser(userId string) bool {
	user := User{}
	result := database.DBconnect.Where("id=?", userId).Delete(&user)
	log.Println(result)
	if result.RowsAffected == 0 {
		return false
	}
	return true

}

//Update
func UpdateUser(userId string, user User) User {
	database.DBconnect.Model(&user).Where("id=?", userId).Updates(user)
	return user
}
