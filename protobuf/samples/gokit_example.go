package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", errors.New("empty string")
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

// For each method, we define request and response structs
type GetRequest struct {
	ID int `json:"id"`
}

type GetResponse struct {
	User model.User `json:"user"`
	Err  string     `json:"err,omitempty"`
}

type ListRequest struct {
	Name string `json:"name"`
}

type ListResponse struct {
	Users []model.User `json:"users"`
	Err   string       `json:"err,omitempty"`
}

type AddRequest struct {
	User model.User `json:"user"`
}

type AddResponse struct {
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeGetEndpoint(svc UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		v, err := svc.Get(req.ID)
		if err != nil {
			return GetResponse{v, err.Error()}, nil
		}
		return GetResponse{v, ""}, nil
	}
}

func makeListEndpoint(svc UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(ListRequest)
		v := svc.List(req.Name)
		return ListResponse{v}, nil
	}
}

// func makeAddEndpoint(svc UserService) endpoint.Endpoint ...

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	svc := UserService{}

	http.Handle("/users/{id}",
		httptransport.NewServer(makeGetEndpoint(svc), decodeGetRequest, encodeResponse),
	)

	http.Handle("/users",
		httptransport.NewServer(makeListEndpoint(svc), decodeListRequest, encodeResponse),
	)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("...")
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("...")
	}

	return GetRequest{ID: userID}, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
