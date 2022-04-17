package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	UID      string `json:"uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"passwd"`
}

var users = []User{
	{UID: "06660", Name: "admin", Email: "admin@net.com", Password: "admin"},
	{UID: "06661", Name: "user", Email: "user@net.com", Password: "user"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUserByName(name string) (*User, error) {
	for i, user := range users {
		if user.Name == name {
			return &users[i], nil
		}
	}

	return nil, errors.New("User(s) not found")
}

func getOneTodo(c *gin.Context) {
	name := c.Param("id")

	user, err := getUserByName(name)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error(), "error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, user)
}

func postUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)

	c.IndentedJSON(http.StatusOK, newUser)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/", getUsers)
	r.GET("/byuid/:id", getOneTodo)
	r.POST("/", postUser)

	r.Run(":8080")
}
