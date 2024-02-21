package main

import (
	"gokit/myexample/internal/order"
	"net/http"
	"os"

	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func main() {

	// go-kit logger
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))

	// create base service
	service := order.NewService(order.NewRepository())

	// attach logger middleware
	service = order.NewLoggingMiddleware(logger, service)

	// create base router
	router := mux.NewRouter()

	// create handler
	order.MakeAddOrderHandler(router, service)

	// start server
	http.ListenAndServe(":8080", router)
}
