package service

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Usdbrl struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type Response struct {
	Usdbrl Usdbrl `json:"USDBRL"`
}

func (c Usdbrl) GetUsdbrl(ctx context.Context) (*Usdbrl, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	b := []byte(body)

	var result Response
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}

	return &result.Usdbrl, nil

}
