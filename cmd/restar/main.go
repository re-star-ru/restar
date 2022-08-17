package main

import (
	"net/http"
	"restar/pkg/user"
)

func main() {
	println("Hello Restar")

	uc := user.NewUserUsecase()
	uh := user.NewHttpHandler(uc)

	m := http.NewServeMux()
	m.HandleFunc("/user/create", uh.Create)

	println("listen at :9090")
	http.ListenAndServe(":9090", m)
}
