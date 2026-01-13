package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/codercollo/websocket/internal/handlers"
)

// routes defines app routes
func routes() http.Handler {
	//Router
	mux := pat.New()

	mux.Get("/ws", http.HandlerFunc(handlers.WsEndPoint))
	mux.Get("/", http.HandlerFunc(handlers.Home))

	//Return router
	return mux
}
