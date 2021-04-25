package auth

import (
	"time"
	"math/rand"
	"crypto/md5"
	"fmt"
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	db "github.com/andrey.kosinov/go-news/database"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Проверка логина и пароля
// Пароль в базе хранится в MD5
func CheckLoginAndPassword(login string, password string) (db.User, bool) {
	md5_password := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	user := db.User{};
	db.DB.Where("login LIKE ?",login).First(&user)

	if (user.Password == md5_password) {
		return user, true
	}

	return db.User{}, false
}

// Аутентификация пользователя путём создания сессии в базе данных
// и выставления Cookie с именем сессии
func Authenticate(user *db.User, c *gin.Context) string {
	uniq_str := user.Login+user.Password+"-"+RandStringRunes(50)
	sess_name := fmt.Sprintf("%x", md5.Sum([]byte(uniq_str)))
	ttl := int64(3600)

	// Удаляем все другие сессии пользователя
	db.DB.Where("user_id = ?", user.ID).Delete(db.Session{})

	session := db.Session{Name: sess_name, UserId: user.ID, TimeToKill: time.Now().Unix()+ttl}
	db.DB.Create(&session)

	c.SetCookie("session", sess_name, 3600, "/", c.Request.Host, false, true)
	c.Redirect(http.StatusFound , "/admin")

	return sess_name;
}

// Проверка существования сессии полученой от браузера из Cookie
func CheckSession(session_name string) (*db.Session ,bool) {
	session := db.Session{}
	result := db.DB.Preload("User").Where("name LIKE ?", session_name).First(&session)

	// Если сессия не найдена возвращаем false в булевой переменной ok
	if (errors.Is(result.Error, gorm.ErrRecordNotFound)) {
		return &db.Session{}, false
	}

	// Проверяем устарела ли сессия
	if (session.TimeToKill < time.Now().Unix()) {
		// Удаляем все другие сессии пользователя
		db.DB.Where("user_id = ?", session.User.ID).Delete(db.Session{})
		return &db.Session{}, false
	}

	return &session, true;
}

// Метод логаута путём удаления пользовательской сессии
func Logout(session *db.Session) {
	db.DB.Where("user_id = ?", session.UserId).Delete(db.Session{})
}

// Функция генерации случайно строки для внутренних нужд пакета
// Но она глобально доступна, так как может понадобится и за
// пределами пакета
func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}