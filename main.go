package main

import (
	//  "fmt" // пакет для форматированного ввода вывода
	"fmt"
	"html/template"
	"net/http" // пакет для поддержки HTTP протокола

	//    "strings" // пакет для работы с  UTF-8 строками
	"log" // пакет для логирования
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler404(w, r, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-type", "text/html")
	t, _ := template.ParseFiles("template/index.html")
	t.Execute(w, map[string]interface{}{"Title": "New wallet"})
}

func errorHandler404(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		t, _ := template.ParseFiles("template/404.html")
		t.Execute(w, "error")
	}
}

func main() {
	http.HandleFunc("/", index) // установим роутер
	fmt.Println("Server is listening on localhost:9000")
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	err := http.ListenAndServe(":9000", nil) // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
