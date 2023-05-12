package services

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/IcaroSilvaFK/full_cycle_goexpert_challenge/cmd/repositories"
)

type QuotationResponse struct {
	USDBRL repositories.Quotation
}

type QuotationService struct {
	context context.Context
}

type QuotationInterface interface {
	CreateNewQuotation() (string, error)
}

func NewQuotationService(context context.Context) QuotationInterface {
	return &QuotationService{
		context: context,
	}
}

func (q *QuotationService) CreateNewQuotation() (string, error) {

	rq, err := http.NewRequestWithContext(q.context, http.MethodGet, "https://economia.awesomeapi.com.br/last/USD-BRL", nil)

	if err != nil {

		return "", err
	}

	rq.Header.Set("Cache-Control", "max-age=604800")
	rs, err := http.DefaultClient.Do(rq)

	if err != nil {

		return "", err
	}

	defer rs.Body.Close()

	bt, err := io.ReadAll(rs.Body)

	if err != nil {

		return "", err
	}

	var quotationResponse QuotationResponse

	if err = json.Unmarshal(bt, &quotationResponse); err != nil {
		return "", err
	}

	b, err := repositories.NewQuotationRepository().Create(quotationResponse.USDBRL)

	if err != nil {
		return "", err
	}

	return b, nil

}
