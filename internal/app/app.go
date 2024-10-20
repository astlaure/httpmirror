package app

import (
	"net/http"
)

func StartApp() *http.Server {
	appServeMux := http.NewServeMux()
	appServeMux.HandleFunc("/", handleIndex)

	appServer := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: appServeMux,
	}

	return &appServer
}
