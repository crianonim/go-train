package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/crianonim/go-train/handlers"
)

var colors = []string{"red", "blue", "yellow"}

type contact struct {
	Name      string
	Age       int
	FavColors []string
}

var jan = contact{"Jan", 39, colors}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello, we are ready to Go!</h1>")
	fmt.Println("root")

}

func main() {
	js, ok := json.Marshal(jan)
	fmt.Println(string(js), ok)

	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/no", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("no. started")
		time.Sleep(5 * time.Second)
		fmt.Fprint(w, "<h1>bo!</h2>")
		fmt.Println("no. finished")

	})
	http.HandleFunc("/t3", handlers.HandlerOne)
	http.HandleFunc("/jan", localHandler)
	http.ListenAndServe(":3000", nil)
}

func init() {
	fmt.Println("main init")
}
