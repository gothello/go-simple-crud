package main

import (
	"fmt"
	"net/http"

	"github.com/gothello/go-web-studies/routes"
)

func main() {
	routes.LoadRoutes()
	fmt.Println("Serve ON 3000")
	http.ListenAndServe(":3000", nil)
}
