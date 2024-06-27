package handlers

import (
	"net/http"
	"strconv"

	"github.com/familybook-project/familybook-api-gin/src/internal/app/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsersEndpoint(c *gin.Context) {
	users, err := users.GetAllUsers()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Users not found",
				"data":    []interface{}{},
			})
			return
		} else {
			c.JSON(500, gin.H{
				"message": err,
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    users,
	})

}

func GetUserByIdEndpoint(c *gin.Context) {
	id := c.Param("id")
	// パラメータが整数でない場合の処理
	if _, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID. User ID must be an integer.",
		})
		return
	}
	user, err := users.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
			"data":    []interface{}{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    user,
	})
}
