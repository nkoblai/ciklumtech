package main

import (
	"net/http"

	"github.com/ciklumtech/router"
)

func main() {
	if err := http.ListenAndServe(":8080", router.New()); err != nil {
		panic(err)
	}
}
