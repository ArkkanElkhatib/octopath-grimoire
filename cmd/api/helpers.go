package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type envelope map[string]interface{}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func (app *application) readQueryString(qs url.Values, key string, defaultValue string) string {
	value := qs.Get(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func (app *application) readQueryStringAsInt(qs url.Values, key string, defaultValue int) int {
	value := qs.Get(key)

	if value == "" {
		return defaultValue
	}

	valueInt, err := strconv.Atoi(value)
	// TODO: Add a way to handle this error rather than utilizng default value
	if err != nil {
		return defaultValue
	}

	return valueInt
}
