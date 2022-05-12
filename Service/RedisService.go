package service

// import (
// 	"golangAPI/pojo"
// 	_ "log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// //Get User
// // @Summary      Get All User
// // @Description  get ID, Name,Password,Email and Birthday of all users in Redis
// // @Tags         Redis
// // @Produce      json
// // @Success      200  {array}   pojo.User
// // @Router /users/ [get]
// func FindAllUsers(c *gin.Context) {
// 	users := pojo.FindAllUsers()
// 	c.JSON(http.StatusOK, users)
// }

// //Post User
// // @Summary      Post User by Id
// // @Description  post ID, Name,Password,Email and Birthday by  user's id in Redis
// // @Tags         Redis
// // @Produce      json
// // @Param        Id   query int false "id"
// // @Success      200  {array}   pojo.User
// // @Router /users/search [post]
// func PostUser(c *gin.Context) {
// 	user := pojo.User{}
// 	err := c.BindJSON(&user)
// 	if err != nil {
// 		c.JSON(http.StatusNotAcceptable, "Error:"+err.Error())
// 		return
// 	}
// 	newUser := pojo.CreateUser(user)
// 	c.JSON(http.StatusOK, newUser)
// }
