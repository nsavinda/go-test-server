package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello,")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback default port
	}

	http.HandleFunc("/", helloHandler)

	addr := ":" + port
	fmt.Printf("Server starting at http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
