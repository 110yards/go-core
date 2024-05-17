package gpubsub

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
)

func ReadEventData[T any](r *http.Request) (T, error) {
	bodyBytes, err := io.ReadAll(r.Body)

	if err != nil {
		return *new(T), err
	}

	r.Body.Close()

	data, err := tryReadAsBase64[T](bodyBytes)

	if err == nil {
		return data, nil
	}

	return tryReadAsObject[T](bodyBytes)
}

func tryReadAsBase64[T any](bytes []byte) (T, error) {
	type event struct {
		Message struct {
			Data string `json:"data"`
		} `json:"message"`
	}

	var m event

	var data T
	err := json.Unmarshal(bytes, &m)

	if err != nil {
		return data, err
	}

	decoded, err := base64.StdEncoding.DecodeString(m.Message.Data)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(decoded, &data)
	return data, err
}

func tryReadAsObject[T any](bytes []byte) (T, error) {
	type event struct {
		Message struct {
			Data T `json:"data"`
		} `json:"message"`
	}

	var m event
	err := json.Unmarshal(bytes, &m)
	return m.Message.Data, err
}
