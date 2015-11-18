package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"

	"gopkg.in/mgo.v2/bson"
)

type HomeController struct {
}


func GetUser(r *http.Request) User {
	user := sessionManager.GetUser(r)
	return user
}

func GetSelfBooks(r *http.Request) []Book {
	user := sessionManager.GetUser(r)
	books:=make([]Book,0)
	dbContext.Books.Find(bson.M{"userId":user.UserId}).All(&books)
	return books
}

func (h *HomeController) Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}
	t := template.New("Layout.html")
	t.Funcs(funcMaps)
	t.ParseFiles("Views/Shared/Layout.html", "Views/Home/Index.html", "Views/Shared/Navigate.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, r)
}

func (h *HomeController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		s1, _ := template.ParseFiles("Views/Home/Login.html")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		s1.Execute(w, nil)
	} else {
		r.ParseForm()
		userName := r.Form.Get("username")
		password := r.Form.Get("password")
		user := User{}
		err := dbContext.Users.Find(bson.M{"nickname": userName, "password": password}).One(&user)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		session := sessionManager.CreateSession()
		session.Values["UserName"] = userName
		session.Values["Header"] = "/uploads/header.jpg"
		session.Values["UserId"] = user.UserId.Hex()
		cookie := http.Cookie{Name: sessionManager.AppName, Value: url.QueryEscape(session.Id), Path: "/", HttpOnly: true, MaxAge: int(1000 * 60 * 30)}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (h *HomeController) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		cookie, err := r.Cookie(sessionManager.AppName)
		if err == nil {
			value, _ := url.QueryUnescape(cookie.Value)
			sessionManager.RemoveSession(value)
		} else {
			fmt.Println(err)
		}
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
