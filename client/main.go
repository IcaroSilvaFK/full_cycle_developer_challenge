package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Bid string `json:"bid"`
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)

	if err != nil {
		log.Fatal(err)
	}

	r, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	f, err := os.Open("cotacao.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	b, err := io.ReadAll(r.Body)

	fmt.Println(string(b))
	if err != nil {
		log.Fatal(err)
	}

	var response Response

	if err = json.Unmarshal(b, &response); err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)

}
