package main

import (
	"WEB4/myapp"
	"net/http"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
