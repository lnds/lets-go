package main

import (
	"fmt"
	"net/http"
)

func main(){
  mux := http.NewServeMux()
  mux.HandleFunc("GET /hola", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "¡Hola Web!")
  })
  mux.HandleFunc("GET /saluda/{nombre}", func(w http.ResponseWriter, r *http.Request) {
    nombre := r.PathValue("nombre")
    fmt.Fprintf(w, "¡Hola %s!", nombre)
  })
  http.ListenAndServe(":8000", mux)
}

