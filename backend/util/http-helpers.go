package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

// ReadJsonAndValidate reads JSON from the request body and validates it.
func ReadJsonAndValidate(w http.ResponseWriter, r *http.Request, data any) error {
	if err := readJsonFromBody(w, r, data); err != nil {
		return err
	}

	if err := ValidateRequest(data); err != nil {
		return err
	}

	return nil
}

// readJsonFromBody reads JSON from the request body
func readJsonFromBody(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 10 << 20 // one megabyte

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return errors.New("json improperly formatted or data types do not match")
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only have a single JSON value")
	}

	return nil
}

// RebuildRequestBody rebuilds the request body.
func RebuildRequestBody(r *http.Request, data any) error {
	var requestBodyBuffer bytes.Buffer
	encoder := json.NewEncoder(&requestBodyBuffer)

	if err := encoder.Encode(data); err != nil {
		return err
	}
	r.Body = io.NopCloser(bytes.NewReader(requestBodyBuffer.Bytes()))
	r.Header.Set("Content-Type", "application/json")

	return nil
}

// WriteJson returns a JSON response.
func WriteJson(w http.ResponseWriter, status int, data any, headers ...http.Header) {
	out, err := json.Marshal(data)
	if err != nil {
		log.Println("unable to marshal json:", err)
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		log.Println("unable to write json:", err)
	}
}

// ErrorJson returns an error in JSON format.
func ErrorJson(w http.ResponseWriter, err error) error {
	if converted, ok := err.(*ErrorModel); ok {
		WriteJson(w, converted.Status, converted)
		return nil
	} else {
		return errors.New("error is not of type error model")
	}
}
