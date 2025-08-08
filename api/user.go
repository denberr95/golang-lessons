package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = []User{
	{
		ID:      "1",
		Name:    "Mario",
		Surname: "Rossi",
	},
	{
		ID:      "2",
		Name:    "Luigi",
		Surname: "Verdi",
	},
}

func registerUserRoutes(r *gin.Engine) {
	group := r.Group("v1/users")
	group.GET("", listUsers)
	group.POST("", createUser)
	group.GET("/:id", getUser)
}

func createUser(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func listUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, responseUserNotFound(id))
}

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type UserNotFound struct {
	Msg string `json:"msg"`
	ID  string `json:"id"`
}

func responseUserNotFound(id string) UserNotFound {
	return UserNotFound{
		Msg: "User not found",
		ID:  id,
	}
}
