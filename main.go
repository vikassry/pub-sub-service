package main

import (
	"fmt"
	"net/http"
	"pub-sub-service/router"
)

func main() {
	fmt.Println("Hello world!")
	ginRouter := router.SetupRouter()
	_ = http.ListenAndServe(":8080", ginRouter)
}
