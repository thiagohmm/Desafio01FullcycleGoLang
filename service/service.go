package service

import (
	"encoding/json"
	"io"
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

// Remove the unused method
func (c Usdbrl) GetUsdbrl() (*Usdbrl, error) {
	resp, error := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if error != nil {
		return nil, error
	}
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}
	var response Response
	error = json.Unmarshal(body, &response)
	if error != nil {
		return nil, error
	}
	return &response.Usdbrl, nil
}
