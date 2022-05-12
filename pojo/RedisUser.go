package pojo

//GET
func FindAllUsersr() []User {
	var users []User //聲明模型結構體類型變user slice
	// database.Rdbconnect.HGet(&users) //select * from users;
	return users
}

//POST
func CreateUserr(user User) User {
	// database.Rdbconnect.HSet(&user)
	return user
}
