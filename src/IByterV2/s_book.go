package main

import (
	"html/template"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type BookService struct {
}

func (h *BookService) Delete(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	dbContext.Books.RemoveId(bson.ObjectIdHex(bookId))
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *BookService) Books(w http.ResponseWriter, r *http.Request) {
	t := template.New("Layout.html")
	t.Funcs(funcMaps)
	t.ParseFiles("Views/Shared/Layout.html", "Views/Profile/Index.html",
		"Views/Profile/Books.html", "Views/Profile/SideBar.html", "Views/Shared/Navigate.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, r)
}

func (h *BookService) Add(w http.ResponseWriter, r *http.Request) {
	t := template.New("Layout.html")
	t.Funcs(funcMaps)
	if r.Method == "GET" {
		t.ParseFiles("Views/Shared/Layout.html", "Views/Profile/Index.html",
			"Views/Profile/AddBook.html", "Views/Profile/SideBar.html", "Views/Shared/Navigate.html")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t.Execute(w, r)
	} else {
		r.ParseForm()
		user := sessionManager.GetUser(r)
		title := r.Form.Get("title")
		summary := r.Form.Get("summary")
		bookId := r.Form.Get("bookId")
		book := Book{}
		book.UserId = user.UserId
		if bookId != "" {
			book.BookId = bson.ObjectIdHex(bookId)
		} else {
			book.BookId = bson.NewObjectId()
		}
		book.CreateTime = time.Now()
		book.ModifyTime = time.Now()
		book.Summary = summary
		book.Title = title
		dbContext.Books.UpsertId(book.BookId, book)
		http.Redirect(w, r, "/books", http.StatusFound)
	}
}