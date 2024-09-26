package httpclient

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/amledigital/arcxp-circulations/internal/models"
)

type HttpClient struct {
	Method    string `json:"method"`
	URL       string `json:"url"`
	Body      []byte `json:"data"`
	AuthToken string `json:"-"`
}

func NewHttpClient(method, url, token string, data []byte) *HttpClient {
	return &HttpClient{
		Method:    method,
		URL:       url,
		Body:      data,
		AuthToken: token,
	}
}

func (h *HttpClient) FetchCirculationsByID(documentID, website string) (*models.Circulations, error) {

	client := &http.Client{}

	req, err := http.NewRequest(h.Method, h.URL, bytes.NewReader(h.Body))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+h.AuthToken)

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			return nil, err
		}

		var c = &models.Circulations{}

		err = json.Unmarshal(body, &c)

		if err != nil {
			return nil, err
		}

		return c, nil

	}

	return nil, nil

}
