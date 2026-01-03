package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Backend Go prêt pour ts-migrate!")
    })
    fmt.Println("Serveur Go démarré sur : http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
