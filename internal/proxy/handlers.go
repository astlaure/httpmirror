package proxy

import (
	"fmt"
	"github.com/astlaure/httpmirror/internal/messages"
	"io"
	"net/http"
	"sync"
)

func handleProxy(w http.ResponseWriter, r *http.Request) {
	service := r.PathValue("service")
	path := r.PathValue("path")

	destination, err := GetServiceFromRoute(service)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// create client
	client := http.Client{}
	requests, err := CreateRequests(destination, path, r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	requestGroup := sync.WaitGroup{}
	var activeResponse *http.Response
	var previewResponse *http.Response

	requestGroup.Add(1)
	go func() {
		defer requestGroup.Done()
		activeResponse, _ = client.Do((*requests)[0])
	}()

	requestGroup.Add(1)
	go func() {
		defer requestGroup.Done()
		previewResponse, _ = client.Do((*requests)[1])
	}()

	// call both
	requestGroup.Wait()

	// save results
	defer previewResponse.Body.Close()
	previewBodyBytes, err := io.ReadAll(previewResponse.Body)

	if err != nil {
		return
	}

	// return response
	defer activeResponse.Body.Close()
	bodyBytes, err := io.ReadAll(activeResponse.Body)

	if err != nil {
		return
	}

	proxyRequest := messages.ProxyRequest{
		Service:  destination.Name,
		Tracking: (*requests)[0].Header.Get("X-Shadow-Tracking"),
	}

	activeMessage := messages.ProxyMessage{
		Status:   uint(activeResponse.StatusCode),
		Path:     activeResponse.Request.URL.Path,
		Protocol: activeResponse.Request.Proto,
		Headers:  StringifyHeaders(activeResponse),
		Body:     string(bodyBytes),
	}

	previewMessage := messages.ProxyMessage{
		Status:   uint(previewResponse.StatusCode),
		Path:     previewResponse.Request.URL.Path,
		Protocol: previewResponse.Request.Proto,
		Headers:  StringifyHeaders(previewResponse),
		Body:     string(previewBodyBytes),
	}

	messages.CreateProxyRequest(proxyRequest, activeMessage, previewMessage)

	fmt.Println(activeResponse.StatusCode)
	fmt.Println(string(bodyBytes))

	fmt.Println(previewResponse.StatusCode)
	fmt.Println(string(previewBodyBytes))

	copyResponseHeaders(activeResponse, w)
	w.WriteHeader(activeResponse.StatusCode)
	w.Write(bodyBytes)
}
