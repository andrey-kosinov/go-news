package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	db "github.com/andrey.kosinov/go-news/database"
	"github.com/andrey.kosinov/go-news/auth"
	// "fmt"
)

// Вывод списка новостей для обычных пользователей
func Index(c *gin.Context) {
	news := []db.News{}
	db.DB.Order("id desc").Find(&news)

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"news": news,
	})
}

// Вывод одной конкретной новости по ID
func News(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	news := db.News{};
	db.DB.First(&news, id)

	if (news.ID == 0) {
		c.String(404, "Запись не найдена")
		return
	}

	c.HTML(http.StatusOK, "news.tmpl", gin.H{
		"news": news,
	})

	news = db.News{};
}

// Вывод формы для ввода логина и пароля
func LoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{})
}

// Метод попытки аутентификации
// Является только оберткой, подготавливающей данные для
// методов CheckLoginAndPassword и Autheticate из пакета auth, который и делает всю магию
func Login(c *gin.Context) {

	login := c.PostForm("login")
	password := c.PostForm("password")

	user, ok := auth.CheckLoginAndPassword(login, password)
	if (ok) {
		auth.Authenticate(&user, c)
	} else {
		c.String(503, "Not authorized")
	}

}
