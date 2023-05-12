package main

import (
	"fmt"
	"net/http"

	"github.com/IcaroSilvaFK/full_cycle_goexpert_challenge/cmd/controllers"
)

func main() {

	http.HandleFunc("/cotacao", controllers.QuotationController)

	fmt.Println("🚀Listening on port localhost:8080")

	http.ListenAndServe(":8080", nil)
}
