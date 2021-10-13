package main

import (
	"fmt"
	"net/http"
)

func main() {
	PORT := ":8080"
	fmt.Printf("Server running at http://localhost%s", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		panic(err)
	}
}

