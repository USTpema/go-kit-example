package order

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeAddOrderHandler(r *mux.Router, s Service) {

	// server options are like middleware for requests/response
	opts := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeErrorResponse),
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(beforeRequest),
	}

	// create server
	addOrderHandler := httptransport.NewServer(
		// basic.AuthMiddleware("admin", "user", "Sample Realm")(makeAddOrderEndpoint(s)),
		makeAddOrderEndpoint(s),
		decodeAddOrderRequest,
		encodeResponse,
		opts...,
	)

	// attach endpoints
	r.Handle("/api/order/add", addOrderHandler).Methods(http.MethodPost)
}

func decodeAddOrderRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		CustomerName string  `json:"customer_name"`
		Address      string  `json:"address"`
		Amount       float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return addOrderRequest{body.CustomerName, body.Address, body.Amount}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err, ok := response.(error); ok && err != nil {
		encodeErrorResponse(ctx, err, w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func encodeErrorResponse(ctx context.Context, err error, w http.ResponseWriter) {
	errResp := struct {
		Timestamp time.Time `json:"timestamp"`
		Path      string    `json:"path"`
		Status    int       `json:"status"`
		Msg       string    `json:"message"`
	}{
		Timestamp: time.Now(),
		Path:      ctx.Value(ctxKey("path")).(string),
		Status:    http.StatusBadRequest,
		Msg:       err.Error(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(errResp)
}

type ctxKey string

func beforeRequest(ctx context.Context, r *http.Request) context.Context {
	return context.WithValue(ctx, ctxKey("path"), r.URL.RequestURI())
}
