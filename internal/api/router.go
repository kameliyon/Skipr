package api

import (
	"net/http"
)

func SetupRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/player", PlayerHandler)
    mux.HandleFunc("/players", PlayersHandler)
	mux.HandleFunc("/lineup", LineupHandler)

	return mux
}
