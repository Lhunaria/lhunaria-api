package main

import (
    "fmt"
    "log"
    "net/http"
)

// helloHandler handles the "/hello" route.
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

// greetHandler handles the "/greet/{name}" route.
func greetHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Path[len("/greet/"):]
    if name == "" {
        http.Error(w, "Name is missing", http.StatusBadRequest)
        return
    }
    fmt.Fprintf(w, "Hello, %s!\n", name)
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/greet/", greetHandler)

    fmt.Println("Starting server on :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}
