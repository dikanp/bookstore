package users

import (
	// "encoding/json"
	"fmt"
	"strconv"
	// "io/ioutil"
	"net/http"

	"github.com/dikanp/domain/users"
	"github.com/dikanp/services"
	"github.com/dikanp/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	//Handle Error
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	//handle json error
	// 	fmt.Println(err.Error())
	// 	return
	// }
	
	if err := c.ShouldBindJSON(&user); err!= nil {
		// fmt.Println(err.Error())

		restErr := errors.NewBadRequestError("invalid json body")

		// restErr := errors.RestErr{
		// 	Message: "invalid json body",
		// 	Status: http.StatusBadRequest,
		// 	Error: "bad_request",
		// }

		// c.AbortWithStatusJSON(400, gin.H{"status": false, "message": err.Error()})
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//error create
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
	// c.String(http.StatusOK, "pong")
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 0, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId);
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

// func GetAllUser(c *gin.Context) {
// 	users, getErr := services.GetAllUser();
// 	if getErr != nil {
// 		c.JSON(getErr.Status, getErr)
// 		return
// 	}
// 	c.JSON(http.StatusOK, users)
// }

func SearchUser(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func Tes(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
