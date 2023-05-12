package repositories

import (
	"context"
	"time"

	"github.com/IcaroSilvaFK/full_cycle_goexpert_challenge/cmd/db"
	"github.com/google/uuid"
)

type Quotation struct {
	id        string
	code      string
	codein    string
	name      string
	high      string
	low       string
	varBid    string
	pctChange string
	Bid       string `json:"bid"`
	ask       string
	timestamp string
}

type QuotationInterface interface {
	Create(Quotation) error
}

func NewQuotationRepository() QuotationInterface {
	return &Quotation{}
}

func (*Quotation) Create(q Quotation) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)

	defer cancel()

	q.id = uuid.New().String()
	db := db.NewDatabaseConnection()

	stmt, err := db.PrepareContext(ctx, "INSERT INTO quotations (id,code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp) VALUES ($1, $2, $3, $4, $5, $6, $7,$8, $9, $10, $11)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(q.id, q.code, q.codein, q.name, q.high, q.low, q.varBid, q.pctChange, q.Bid, q.ask, q.timestamp)

	if err != nil {
		return err
	}

	return nil
}
