package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// ParseBody reads and unmarshals the body of the HTTP request into the provided struct
func ParseBody(r *http.Request, x interface{}) error {
	// ensure the request body is closed after reading
	defer r.Body.Close()

	// read the entire body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading body: %v", err)
		return err
	}

	// unmarshall the JSON body into the provided struct
	err = json.Unmarshal(body, x)
	if err != nil {
		log.Printf("error unmarshalling JSON: %v", err)
		return err
	}

	// return nil if everything was successful
	return nil
}
