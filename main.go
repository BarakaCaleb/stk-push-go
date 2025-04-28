package main

import (
	"fmt"
	"log"
	"net/http"
	"stkpush/internal/config"
	"stkpush/internal/handlers"
)

func main() {
	config.LoadEnv()

	http.HandleFunc("/callback", handlers.CallbackHandler)
	http.HandleFunc("/pay", handlers.PayHandler)

	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
