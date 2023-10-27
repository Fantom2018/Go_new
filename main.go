package main

import (
	"flag"
	"fmt"
	"net/http"
)

type User struct{}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	var user = User{}
	_ = user
	fmt.Fprint(w, "gfhgfhgjkfdgh")

}

func main() {
	listenAddr := flag.String("lictenaddr", ":4999", "traaad")
	flag.Parse()

	http.HandleFunc("/user", handleGetUser)

	http.ListenAndServe(*listenAddr, nil)

}
