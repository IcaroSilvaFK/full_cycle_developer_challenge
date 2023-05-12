package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/IcaroSilvaFK/full_cycle_goexpert_challenge/cmd/repositories"
	"github.com/IcaroSilvaFK/full_cycle_goexpert_challenge/cmd/services"
)

type QuotationResponse struct {
	USDBRL repositories.Quotation
}

func QuotationController(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"message":"Method not allowed"}`))
		return
	}

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)

	svc := services.NewQuotationService(ctx)

	defer cancel()

	b, err := svc.CreateNewQuotation()

	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":` + err.Error() + `}`))
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"bid":` + b + `}`))
}
