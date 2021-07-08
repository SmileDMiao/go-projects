package controller

import (
	"go-learn/global"
	"go-learn/model"
	"go-learn/service"
	"go-learn/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// LoginParams struct
type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login action
func Login(c *gin.Context) {
	var login LoginParams
	var user model.User

	c.ShouldBindJSON(login)

	login.Password = utils.MD5V([]byte(login.Password))
	result := global.DB.Where("username = ? AND password = ?", login.Username, login.Password).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"message": "用户名不存在或者密码错误"})
	} else {
		tokenNext(c, user)
	}
}

// 登录以后签发jwt
func tokenNext(c *gin.Context, user model.User) {
	j := service.NewJWT()
	claims := service.CustomClaims{
		ID:         user.ID,
		Username:   user.Username,
		NickName:   user.NickName,
		BufferTime: global.CONFIG.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + global.CONFIG.JWT.ExpiresTime,
			Issuer:    "malzahar",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "获取token失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

// Register controller
func Register(c *gin.Context) {
	var u model.User
	c.BindJSON(&u)

	user := model.User{
		Username: u.Username,
		NickName: u.NickName,
		Password: utils.MD5V([]byte(u.Password)),
	}
	result := global.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"message": "创建失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

// GetArticles controller
func GetArticles(c *gin.Context) {
	var user model.User
	var articles []model.Article

	global.DB.First(&user, c.Param("id"))
	global.DB.Model(&user).Association("Articles").Find(&articles)

	c.JSON(http.StatusOK, gin.H{"articles": articles})
}
