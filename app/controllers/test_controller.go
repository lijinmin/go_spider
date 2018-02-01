package controllers
import "net/http"
import "io"
import "html/template"
import "fmt"
import "os"
import "log"
import "path/filepath"
  func SayHello(w http.ResponseWriter, r *http.Request) {
  	// io.WriteString(w,"hello world")
      wd,_ := os.Getwd()
      s1,_ := template.ParseFiles(filepath.Join(wd,"app/views/layouts/test_layout.html"),filepath.Join(wd,"app/views/test/sayhello.html"))
      s1.ExecuteTemplate(w, "test_layout", nil)
      s1.Execute(w, nil)
  }

func SayBye(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w,"good bye")
	// s1 := template.New("some template") //创建一个模板
	// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	wd, _ := os.Getwd()
	log.Print("Work directory:", wd)
	fmt.Println(r.URL.Path)
	// s1, _ := template.ParseFiles("../views/test/header.tmpl", "../views/test/content.tmpl", "../views/test/footer.tmpl")
    	s1,_ := template.ParseFiles(filepath.Join(wd,"app/views/test/header.tmpl"),filepath.Join(wd,"app/views/test/content.tmpl"),filepath.Join(wd,"app/views/test/footer.tmpl"))
    	// s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
   	s1.ExecuteTemplate(w, "content", nil)
   	// fmt.Println(r.URL.Path)
    	// s1.ExecuteTemplate(os.Stdout, "footer", nil)
    	// fmt.Println()
    	s1.Execute(w, nil)
  }

  func Search(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "这是从后台发送的数据")
  }
