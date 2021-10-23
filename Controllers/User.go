package Controllers

import (
	"fmt"
	"net/http"

	"github.com/iamrahultanwar/friskco/Models"

	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) {
	var user Models.User
	userId := c.GetFloat64("userId")
	err := Models.GetCurrentUser(&user, int(userId))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func RegisterUser(c *gin.Context) {
	type RegisterUserResponse struct {
		Message string `json:"message"`
		Email   string `json:"email"`
	}
	var user Models.User
	c.BindJSON(&user)
	rows, userError := Models.FindUserByEmail(&user)
	if rows > 0 {
		c.String(http.StatusBadRequest, "User already registered")
		c.Abort()
	} else {
		err := Models.CreateUser(&user)
		if err != nil || userError != nil {
			c.String(http.StatusInternalServerError, err.Error())
			c.Abort()
		} else {
			ru := RegisterUserResponse{
				Message: "User Registered Successfully",
				Email:   user.Email,
			}
			c.JSON(http.StatusOK, ru)
		}
	}

}

func LoginUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	token, err := Models.LoginUser(&user)
	type loginRespose struct {
		Token   string `json:"token"`
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}
	if err != nil {
		res := loginRespose{
			Status:  false,
			Message: err.Error(),
		}
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, res)
	} else {
		res := loginRespose{
			Status:  true,
			Message: "Login Successful",
			Token:   token,
		}
		c.JSON(http.StatusOK, res)
	}
}
