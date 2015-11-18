package main

import (
	"html/template"
	"net/http"
	"gopkg.in/mgo.v2/bson"
)

type ProfileController struct {
}

var funcMaps = template.FuncMap{
	"getUser":          GetUser,
	"getSelfBooks":     GetSelfBooks,
	"getBook":          GetBook,
	"objectIdToString": ObjectIdToString,
}

func GetBook(r *http.Request) Book {
	bookId := r.FormValue("bookId")
	book := Book{}
	if bookId==""{
		return book
	}
	dbContext.Books.Find(bson.M{"_id": bson.ObjectIdHex(bookId)}).One(&book)
	return book
}

func ObjectIdToString(id bson.ObjectId) string {
	return id.Hex()
}

func (h *ProfileController) Index(w http.ResponseWriter, r *http.Request) {
	t := template.New("Layout.html")
	t.Funcs(funcMaps)
	t.ParseFiles("Views/Shared/Layout.html", "Views/Shared/Navigate.html",
		"Views/Profile/Index.html", "Views/Profile/Sidebar.html", "Views/Profile/Main.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, r)
}
