package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	db "github.com/andrey.kosinov/go-news/database"
	"github.com/andrey.kosinov/go-news/auth"
)

// Вывод списка новостей для администратора
func AdminIndex(c *gin.Context) {
	news := []db.News{}
	db.DB.Order("id desc").Find(&news)
	session := c.MustGet("session")

	c.HTML(http.StatusOK, "admin-index.tmpl", gin.H{
		"news": news,
		"session": session,
	})
}

// Удаление новости
func AdminNewsDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	db.DB.Where("id = ?", id).Delete(db.News{})

	c.Redirect(http.StatusTemporaryRedirect, "/admin")
}

// Вывод формы свойств новости
func AdminNewsForm(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	news := db.News{};
	if (id > 0) {
		db.DB.First(&news, id)
	}

	c.HTML(http.StatusOK, "admin-news-form.tmpl", gin.H{
		"news": news,
	})
}

// Сохранение новости
// Если в запросе пришёл ID новости, то меняем новость,
// а если без ID - то добавляем
func AdminNewsStore(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	title := c.PostForm("title")
	body := c.PostForm("body")

	if (id > 0) {
		db.DB.Model(&db.News{}).Where("id = ?", id).Updates(db.News{Title: title, Body: body})
	} else {
		news := db.News{Title: title, Body: body}
		db.DB.Create(&news)
		news = db.News{}
	}

	c.Redirect(http.StatusFound, "/admin")
}

// Выход из административного сектора путём убийства сессии
func AdminLogOut(c *gin.Context) {
	session := c.MustGet("session").(*db.Session)
	auth.Logout(session)
	c.Redirect(http.StatusTemporaryRedirect, "/admin")
	c.AbortWithStatus(503)
}