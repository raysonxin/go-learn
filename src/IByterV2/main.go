// IByter project IByter.go
package main

import (
	"fmt"
	"net/http"
	"swk/inihelper"
)

var sessionManager *SessionManager
var dbContext DbContext

func main() {
	conf := inihelper.SetConfig("./conf.ini")
	dbIp := conf.GetValue("DataBase", "Ip")
	fmt.Println("初始化数据库", dbIp)
	dbContext = NewDbContext(dbIp)
	fmt.Println("数据库初始化完毕")
	sessionManager = &SessionManager{make(map[string]*Session, 0), "swk"}
	handle()
	err := http.ListenAndServe(":8008", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func handle() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("Content"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("Script"))))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("Uploads"))))

	home := HomeController{}
	http.HandleFunc("/", home.Index)
	http.HandleFunc("/login", home.Login)
	http.HandleFunc("/logout", home.Logout)

	profile := ProfileController{}
	http.HandleFunc("/profile", profile.Index)

    bookService:=BookService{}
	http.HandleFunc("/books", bookService.Books)
	http.HandleFunc("/books/add", bookService.Add)
    http.HandleFunc("/books/delete", bookService.Delete)
}
