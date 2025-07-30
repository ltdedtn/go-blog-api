package main

import (
	"fmt"
	"go-blog-api/routes"
	"log"
	"net/http"
)

func main() {
    r := routes.RegisterRoutes()
    fmt.Println("Server running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
