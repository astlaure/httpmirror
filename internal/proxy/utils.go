package proxy

import (
	"net/http"
)

func copyHeaders(req *http.Request, r *http.Request) {
	for key, header := range r.Header {
		for _, value := range header {
			req.Header.Add(key, value)
		}
	}
}

func copyResponseHeaders(res *http.Response, w http.ResponseWriter) {
	for key, header := range res.Header {
		for _, value := range header {
			w.Header().Set(key, value)
		}
	}
}

func StringifyHeaders(res *http.Response) string {
	headers := ""

	for key, header := range res.Header {
		for _, value := range header {
			headers += key + ": " + value + "\r\n"
		}
	}

	return headers
}
