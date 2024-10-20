package proxy

import (
	"github.com/astlaure/httpmirror/internal/messages"
	"net/http"
)

func CreateProxy() *http.Server {
	messages.InitTable()

	proxyServeMux := http.NewServeMux()

	proxyServeMux.HandleFunc("/{service}", handleProxy)
	proxyServeMux.HandleFunc("/{service}/{path...}", handleProxy)

	proxyServer := http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: proxyServeMux,
	}

	return &proxyServer
}
