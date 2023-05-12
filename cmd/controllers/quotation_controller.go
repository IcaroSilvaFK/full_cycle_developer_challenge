package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/IcaroSilvaFK/full_cycle_goexpert_challenge/cmd/repositories"
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
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)

	defer cancel()

	rq, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/last/USD-BRL", nil)

	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf(`{"message":"Error on prepare request","error": %v}`, err)
		w.Write([]byte(msg))
		return
	}

	rq.Header.Set("Cache-Control", "max-age=604800")
	rs, err := http.DefaultClient.Do(rq)

	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf(`{"message":"Error on request economia api","error": %v}`, err)
		w.Write([]byte(msg))
		return
	}

	defer rs.Body.Close()

	bt, err := io.ReadAll(rs.Body)

	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf(`{"message":"Error on read body","error": %v}`, err)
		w.Write([]byte(msg))
		return
	}

	var quotationResponse QuotationResponse

	if err = json.Unmarshal(bt, &quotationResponse); err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf(`{"message":"Error on unmarshal","error": %v}`, err)
		w.Write([]byte(msg))
		return
	}

	err = repositories.NewQuotationRepository().Create(quotationResponse.USDBRL)

	if err != nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf(`{"message":"Error on create quotation","error": %v}`, err)
		w.Write([]byte(msg))
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(quotationResponse.USDBRL)

}
