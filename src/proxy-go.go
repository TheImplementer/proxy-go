package main

import (
	"net/http"
	"log"
	"bytes"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(responseWriter http.ResponseWriter, request *http.Request) {
	log.Printf("Request received: %v %v", request.Method, request.RequestURI)

	httpClient := http.DefaultClient
	response, err := httpClient.Do(copyRequest(request))
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	for header, value := range response.Header {
		for _, headerValue := range value {
			responseWriter.Header().Add(header, headerValue)
		}
	}

	responseBuffer := &bytes.Buffer{}
	responseBuffer.ReadFrom(response.Body)
	responseWriter.WriteHeader(response.StatusCode)
	responseWriter.Write(responseBuffer.Bytes())
}

func copyRequest(request *http.Request) (*http.Request) {
	copy, _ := http.NewRequest(request.Method, request.RequestURI, request.Body)

	for header, value := range request.Header {
		for _, headerValue := range value {
			copy.Header.Add(header, headerValue)
		}
	}

	return copy
}
