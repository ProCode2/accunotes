package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

type RequestBody io.ReadCloser

func BodyAsJson[T any](b RequestBody, filler *T) error {
	decoder := json.NewDecoder(b)

	err := decoder.Decode(&filler)

	return err
}

func WriteJson[T any](w http.ResponseWriter, payload T) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(payload)
}
