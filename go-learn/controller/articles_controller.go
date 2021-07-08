package controller

import (
	"go-learn/global"
	"go-learn/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetArticle controller
func GetArticle(c *gin.Context) {
	var article model.Article

	global.DB.First(&article, c.Param("id"))
	global.DB.Preload("User").Find(&article)

	c.JSON(http.StatusOK, gin.H{"article": article})
}

// CreateArticle controller
func CreateArticle(c *gin.Context) {
	var a model.Article
	c.BindJSON(&a)

	article := model.Article{
		Title:   a.Title,
		Content: a.Content,
		UserID:  a.UserID,
	}
	result := global.DB.Create(&article)

	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"message": "创建失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"": article})
	}
}
