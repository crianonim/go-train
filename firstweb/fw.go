package main

import (
	"fmt"
	"net/http"
)

func localHandler(w http.ResponseWriter, request *http.Request) {
	fmt.Println("Method", request.Method)
	coo := http.Cookie{Name: "ciacho", Value: "Jan"}
	http.SetCookie(w, &coo)
	// w.Header().Add("set-cookie", "mojekukie")
	request.ParseForm()
	for k, v := range request.Header {
		fmt.Fprintln(w, k, v)
	}
	reqVar := []byte(fmt.Sprint(request))
	w.Write(reqVar)
}

func init() {
	fmt.Println("fw init")
}
