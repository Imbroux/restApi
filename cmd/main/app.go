package main

import "restApi/internal/user"

func main() {
	h := &user.Handler{}
	e := h.Register()
	e.Logger.Fatal(e.Start(":8080"))
}
