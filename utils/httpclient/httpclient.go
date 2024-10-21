package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/amledigital/arcxp-circulations/internal/models"
)

type HttpClient struct {
	Client    *http.Client `json:"-"`
	Method    string       `json:"method"`
	URL       string       `json:"url"`
	Body      []byte       `json:"data"`
	AuthToken string       `json:"-"`
}

func NewHttpClient(method, url, token string, data []byte) *HttpClient {
	return &HttpClient{
		Client:    &http.Client{},
		Method:    method,
		URL:       url,
		Body:      data,
		AuthToken: token,
	}
}

func (h *HttpClient) FetchArticlesBySectionID(from ...int) (contentElements *models.ContentApiFetchResult, more int, err error) {

	if len(from) > 0 {
		var val int

		val = from[0]
		h.URL = strings.Join([]string{h.URL, fmt.Sprintf("from=%d", val)}, "&")
	}
	client := &http.Client{}

	req, err := http.NewRequest(h.Method, h.URL, bytes.NewReader(h.Body))

	if err != nil {
		return nil, 0, err
	}

	req.Header.Add("Authorization", "Bearer "+h.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return nil, 0, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
			return nil, 0, err
		}

		var c = &models.ContentApiFetchResult{}

		err = json.Unmarshal(body, &c)

		if err != nil {
			return nil, 0, err
		}

		return c, c.Next, nil

	}

	return nil, 0, nil

}

func (h *HttpClient) FetchCirculationsByID(documentID, website string) (*models.Circulations, error) {

	client := &http.Client{}

	req, err := http.NewRequest(h.Method, h.URL, bytes.NewReader(h.Body))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+h.AuthToken)

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
			return nil, err
		}

		var c = &models.Circulations{}

		err = json.Unmarshal(body, &c)

		if err != nil {
			log.Fatalln(err)
			return nil, err
		}

		return c, nil

	}

	return nil, nil

}

func (h *HttpClient) CirculateADocument() (*models.Circulation, error) {

	req, err := http.NewRequest(h.Method, h.URL, bytes.NewReader(h.Body))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+h.AuthToken)

	resp, err := h.Client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.Body != nil {

		defer resp.Body.Close()

		var opResult *models.Circulation

		if err = json.NewDecoder(resp.Body).Decode(&opResult); err != nil {
			return nil, err
		}

		return opResult, nil

	}

	return nil, nil

}
