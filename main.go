package main

import (
	"net/http"

	"github.com/cemezgn/keyValueApp/pkg/file"
	key_value "github.com/cemezgn/keyValueApp/pkg/key-value"
)

func main() {
	mux := http.NewServeMux()

	store := file.Read()

	repo := key_value.NewRepository(store)
	service := key_value.NewService(repo)

	mux.Handle("/keys", service)
	mux.Handle("/keys/", service)

	file.Run(repo)

	http.ListenAndServe("localhost:8090", mux)
}

