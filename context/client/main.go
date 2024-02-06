package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("request iniciada")
	defer log.Println("request finalizada")
	select {
	case <-time.After(5 * time.Second):
		log.Printf("request processada com sucesso")
	case <-ctx.Done():
		log.Println("request cancelada pelo cliente")
		http.Error(w, "requeste cancela pelo cliente", http.StatusRequestTimeout)
	}

}
