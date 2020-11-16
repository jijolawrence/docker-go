package main

import (
	"net/http"

	"github.com/jijolawrence/docker-go/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3001", nil)
}
