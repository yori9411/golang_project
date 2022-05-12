package service

import (
	"golangAPI/pojo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Get User
// @Summary      Get All User
// @Description  get ID, Name,Password,Email and Birthday of all users in MySQL
// @Tags         MySQL
// @Produce      json
// @Success      200  {array}   pojo.User
// @Router /users/ [get]
func FindAllUsers(c *gin.Context) {
	users := pojo.FindAllUsers()
	c.JSON(http.StatusOK, users)
}

//Get User by Id
// @Summary      Get User by Id
// @Description  get ID, Name,Password,Email and Birthday by  user's id in MySQL
// @Tags         MySQL
// @Produce      json
// @Param        Id   query int false "id"
// @Success      200  {array}   pojo.User
// @Router /users/search [get]
func FindByUserId(c *gin.Context) {
	user := pojo.FindByUserId(c.Param("id"))
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	log.Println("User -> ", user)
	c.JSON(http.StatusOK, user)
}

//Post User
// @Summary      Post User by Id
// @Description  post ID, Name,Password,Email and Birthday by  user's id in MySQL
// @Tags         MySQL
// @Produce      json
// @Param        Id   query int false "id"
// @Success      200  {array}   pojo.User
// @Router /users/search [post]
func PostUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error:"+err.Error())
		return
	}
	newUser := pojo.CreateUser(user)
	c.JSON(http.StatusOK, newUser)
}

//Delte User
// @Summary      Delete User by Id
// @Description  delete ID, Name,Password,Email and Birthday by  user's id in MySQL
// @Tags         MySQL
// @Produce      json
// @Param        Id   query int false "id"
// @Success      200  {array}   pojo.User
// @Router /users/search [delete]
func DeleteUser(c *gin.Context) {
	user := pojo.DeleteUser(c.Param("id"))
	if !user {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, "Successfully")
}

//Update User
func PutUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
		return
	}
	user = pojo.UpdateUser(c.Param("id"), user)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, user)
}
