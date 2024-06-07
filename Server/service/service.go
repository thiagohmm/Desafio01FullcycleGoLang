package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/thiagohmm/Desafio01FullcycleGoLang/db"
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

func SaveUsdbrl(com *sql.DB, usdbrl *Usdbrl) (int64, error) {
	fmt.Println(usdbrl)
	result, err := com.Exec("INSERT INTO usdbrl (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, createDate) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		usdbrl.Code, usdbrl.Codein, usdbrl.Name, usdbrl.High, usdbrl.Low, usdbrl.VarBid, usdbrl.PctChange, usdbrl.Bid, usdbrl.Ask, usdbrl.Timestamp, usdbrl.CreateDate)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get rows affected: %w", err)
	}
	return rowsAffected, nil
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

	con := db.DB
	rowsR, err := SaveUsdbrl(con, &result.Usdbrl)

	fmt.Print(rowsR)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &result.Usdbrl, nil

}
