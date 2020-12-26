package routes

import (
	"net/http"

	"github.com/gothello/go-web-studies/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("create", controllers.Create)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("update", controllers.Update)
	http.HandleFunc("/delete", controllers.Delete)
}
