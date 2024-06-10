package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/valyala/fastjson" // Replace "your-package-path" with the actual package path
)

func saveToFile(dolar string) error {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(dolar)
	if err != nil {
		return err
	}
	return nil
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)

	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("timeout to write in database")
		}
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	if err != nil {
		log.Fatal(err)
	}
	dolar := "Dólar: { " + v.Get("bid").String() + " }"
	err = saveToFile(dolar)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Dolar", v.Get("bid").String())
	if err != nil {
		log.Fatal(err)
	}

}
