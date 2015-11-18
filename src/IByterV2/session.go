package main

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"time"
    "gopkg.in/mgo.v2/bson"
)

type SessionManager struct {
	Sessions map[string]*Session
	AppName  string //名称
}

func (s *SessionManager) CreateSession() *Session {
	b := make([]byte, 32)
	io.ReadFull(rand.Reader, b)
	sid := base64.URLEncoding.EncodeToString(b)
	session := &Session{sid, make(map[string]string, 0)}
	session.TimeOut()
	s.Sessions[sid] = session
	return session
}

func (s *SessionManager) RemoveSession(id string) {
	delete(sessionManager.Sessions, id)
}

func (s *SessionManager) GetUser(r *http.Request) User {
	cookie, err := r.Cookie(s.AppName)
	user := User{}
	user.IsOnline=false
	if err == nil {
		value, _ := url.QueryUnescape(cookie.Value)
		session, ok := s.Sessions[value]
		if ok {
			user.Nickname = session.Values["UserName"]
			user.Header = session.Values["Header"]
			user.UserId = bson.ObjectIdHex(session.Values["UserId"]) 
			user.IsOnline=true
		}
	}
	return user
}

type Session struct {
	Id     string
	Values map[string]string
}

func (s *Session) Add(key string, value string) {

}

func (s *Session) Remove(key string) {

}

func (s *Session) TimeOut() {
	time.AfterFunc(30*time.Minute, func() {
		delete(sessionManager.Sessions, s.Id)
	})
}
