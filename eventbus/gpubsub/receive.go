package gpubsub

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
)

func ReadEventData[T any](r *http.Request) (T, error) {
	type event struct {
		Message struct {
			Base64 string `json:"data"`
		} `json:"message"`
	}

	var m event
	var data T

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {

		return data, err
	}

	decoded, err := base64.StdEncoding.DecodeString(m.Message.Base64)

	if err != nil {
		return data, err
	}

	err = json.Unmarshal(decoded, &data)

	return data, err
}
