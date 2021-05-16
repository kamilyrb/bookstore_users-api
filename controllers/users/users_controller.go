package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kamilyrb/bookstore_users-api/domain/users"
	"github.com/kamilyrb/bookstore_users-api/services"
	"github.com/kamilyrb/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
	}
	fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")

}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
