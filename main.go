package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_hello/pkg/db"
	"net/http"
)

func getUser(c *gin.Context)  {
	id := c.Query("id")
	var user db.User
	err := db.Client.Model(&db.User{}).Where("id=?", id).Find(&user).Error
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	fmt.Println(db.Client.QueryExpr())
	c.String(http.StatusOK, fmt.Sprintf("%s, %s", user.Name, user.Sex))
}

func main() {
	r := gin.Default()
	r.GET("/get_user", getUser)
	r.Run() // listen and serve on 0.0.0.0:8080
}
