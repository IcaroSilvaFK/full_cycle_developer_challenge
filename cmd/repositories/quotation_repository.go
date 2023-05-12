package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/IcaroSilvaFK/full_cycle_goexpert_challenge/cmd/db"
	"github.com/google/uuid"
)

type Quotation struct {
	ID        string
	Code      string `json:"code"`
	Codein    string `json:"codein"`
	Name      string `json:"name"`
	High      string `json:"high"`
	Low       string `json:"low"`
	VarBid    string `json:"varBid"`
	PctChange string `json:"pctChange"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	Timestamp string `json:"timestamp"`
}

type QuotationInterface interface {
	Create(Quotation) (string, error)
}

func NewQuotationRepository() QuotationInterface {
	return &Quotation{}
}

func (*Quotation) Create(q Quotation) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)

	defer cancel()

	q.ID = uuid.New().String()
	db := db.NewDatabaseConnection()

	stmt, err := db.PrepareContext(ctx, "INSERT INTO quotations (id,code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp) VALUES ($1, $2, $3, $4, $5, $6, $7,$8, $9, $10, $11)")

	if err != nil {
		return "", err
	}
	fmt.Println(q)
	_, err = stmt.Exec(q.ID, q.Code, q.Codein, q.Name, q.High, q.Low, q.VarBid, q.PctChange, q.Bid, q.Ask, q.Timestamp)

	if err != nil {
		return "", err
	}

	return q.Bid, nil
}
