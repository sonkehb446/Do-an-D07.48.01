package sessions

import (
	"Hybrid-blog/constant"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"

	"Hybrid-blog/helpers/utility"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var IDClient = "client"
var NameClient = "undefined"

type InfoUserSession struct {
	Email  string
	UserID string
	Avatar string
	Name   string
	Role   string
}

var infoUserSession = &InfoUserSession{}

func NewCookies(c *gin.Context, userid string, name string, email string, url string, role string) error {
	if email == "" {
		IDClient = utility.RandStringBytes(5)
	} else {
		id := strings.Split(email, "@")
		IDClient = id[0]
	}

	if name == "" {
		name = NameClient
	}

	session, err := store.Get(c.Request, IDClient)
	if err != nil {
		return err
	}
	// Tạo Thời Gian mà cookie sử dụng và cusstom
	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
		// MaxAge:   5,
		HttpOnly: true,
	}
	// Set name session value.
	session.Values[constant.SessionAvatar] = url
	session.Values[constant.SessionUserID] = userid
	session.Values[constant.SessionEmail] = email
	session.Values[constant.SessionIDClient] = IDClient
	session.Values[constant.SessionName] = name
	session.Values[constant.SessionRole] = role
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		return err
	} else {
		infoUserSession.Avatar = url
		infoUserSession.UserID = userid
		infoUserSession.Email = email
		infoUserSession.Name = name
		infoUserSession.Role = role
	}
	return nil
}

func ClearSession(c *gin.Context) {
	session, err := store.Get(c.Request, IDClient)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: false,
	}
	deleteValue(session)
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(session)
}

func deleteValue(session *sessions.Session) {
	delete(session.Values, constant.SessionName)
	delete(session.Values, constant.SessionEmail)
	delete(session.Values, constant.SessionRole)
	delete(session.Values, constant.SessionIDClient)
	delete(session.Values, constant.SessionUserID)
	infoUserSession.UserID = ""
	infoUserSession.Name = ""
	infoUserSession.Avatar = ""
	infoUserSession.Email = ""
	infoUserSession.Role = ""
}

func GetSession() *sessions.CookieStore {
	return store
}

func GetNameclient() string {
	return IDClient
}

func GetInfoUserSession() (*InfoUserSession, error) {
	error := errors.New("NULL InfoUserSession")
	if infoUserSession.UserID != "" {
		return infoUserSession, nil
	}
	return nil, error
}

func UpdateInfoUserSession(url, email, name, role, userid string, c *gin.Context) (*InfoUserSession, error) {
	session, _ := GetSession().Get(c.Request, IDClient)
	session.Values[constant.SessionAvatar] = url
	session.Values[constant.SessionUserID] = userid
	session.Values[constant.SessionEmail] = email
	session.Values[constant.SessionIDClient] = IDClient
	session.Values[constant.SessionName] = name
	session.Values[constant.SessionRole] = role
	error := errors.New("NULL InfoUserSession")
	err := session.Save(c.Request, c.Writer)
	if err != nil {
		return nil, err
	}
	if infoUserSession.UserID != "" {
		infoUserSession.Avatar = url
		infoUserSession.UserID = userid
		infoUserSession.Email = email
		infoUserSession.Name = name
		infoUserSession.Role = role
		return nil, nil
	}
	return nil, error
}
