package main

import (
	"io"
	"fmt"
	"log"
	"net/http"
	"html/template"
	// "os"
	"./config"
)

func main() {
	http.Handle("/app/assets/", http.StripPrefix("/app/assets/", http.FileServer(http.Dir("app/assets"))))
	http.HandleFunc("/bye1", sayBye)
	config.Routes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w,"hello world")
	// s1 := template.New("some template") //创建一个模板
	// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	s1, _ := template.ParseFiles("./app/views/test/header.tmpl", "./app/views/test/content.tmpl", "./app/views/test/footer.tmpl")
    	// s1.ExecuteTemplate(os.Stdout, "header", nil)
	// fmt.Println()
   	s1.ExecuteTemplate(w, "content", nil)
   	fmt.Println(r.URL.Path)               
    	// s1.ExecuteTemplate(os.Stdout, "footer", nil)
    	// fmt.Println()
    	s1.Execute(w, nil)
  }

  func sayHello(w http.ResponseWriter, r *http.Request) {
  	io.WriteString(w,"hello world")
  }