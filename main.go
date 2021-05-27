package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_hello/pkg/db"
	"go_hello/pkg/redis"
	"net/http"
	"time"
)

func getUserInfo(c *gin.Context)  {
	token := c.Query("token")
	jsonStr, err := redis.Client.Get(token).Result()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	var user db.User
	if err := json.Unmarshal([]byte(jsonStr), &user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = db.Client.Model(&db.User{}).Where("id=?", user.Id).Find(&user).Error
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("%s, %s", user.Name, user.Sex))
}

func login(c *gin.Context)  {
	name := c.Query("name")
	password :=c.Query("password")
	var user db.User
	err := db.Client.Model(&db.User{}).Where("name=? and password=?", name, password).
		Find(&user).Error
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	hash := md5.Sum([]byte(fmt.Sprintf("%s:%s:%d", name, password, time.Now().Unix())))
	token := hex.EncodeToString(hash[:])
	jsonByte, err := json.Marshal(user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	_, err = redis.Client.Set(token, string(jsonByte), 1*time.Hour).Result()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, token)
}

func main() {
	r := gin.Default()
	r.GET("/get_user_info", getUserInfo)
	r.GET("/login", login)
	r.Run(":80") // listen and serve on 0.0.0.0:8080
}
