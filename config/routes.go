package config
import "net/http"
import "../app/controllers"
func Routes() {
	http.HandleFunc("/",controllers.SayHello)
	http.HandleFunc("/bye",controllers.SayBye)
	http.HandleFunc("/test/search",controllers.Search)
}