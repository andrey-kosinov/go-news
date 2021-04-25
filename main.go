package main

import (
	"github.com/gin-gonic/gin"
	"github.com/andrey.kosinov/go-news/middleware"
	"github.com/andrey.kosinov/go-news/controller"
	"github.com/andrey.kosinov/go-news/database"
)

func main() {

	// Соединяемся с базой и откладываем отключение от
	// базы до момента завершения
	database.Connect()
	defer database.Disconnect()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Определяем папку с шаблонами
	router.LoadHTMLGlob("templates/*")

	// Публично доступные маршруты
	router.GET("/", controller.Index)
	router.GET("/news/:id", controller.News)
	router.GET("/login", controller.LoginForm)
	router.POST("/login", controller.Login)

	// Группа маршрутов, где необходима аутентификация
	authorized := router.Group("/admin")
	authorized.Use(middleware.AuthRequired())
	{
		authorized.GET("/",controller.AdminIndex)
		authorized.GET("/news/delete/:id",controller.AdminNewsDelete)
		authorized.GET("/news/form/:id",controller.AdminNewsForm)
		authorized.GET("/news/form",controller.AdminNewsForm)
		authorized.POST("/news/store",controller.AdminNewsStore)

		authorized.GET("/logout",controller.AdminLogOut)
	}

	// Запускаем наш сервер
	// Порт по-умолчанию :8080
	router.Run()

}
