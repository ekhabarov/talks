package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type printRequest struct {
	// request parameters
}

type printResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

func makePrintEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(printRequest)
		return printResponse{"Hello Winnipeg!", ""}, nil
	}
}

func decodePrintRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request printRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func main() {
	printHandler := httptransport.NewServer(
		makePrintEndpoint(),
		decodePrintRequest,
		encodeResponse,
	)
	http.Handle("/", printHandler)
	http.ListenAndServe(":8080", nil)
}
