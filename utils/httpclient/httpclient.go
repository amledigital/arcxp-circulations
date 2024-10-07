package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func (h *HttpClient) FetchArticlesBySectionID(sectionID string) (contentElements []interface{}, more int, err error) {

	fmt.Println(h.URL)

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

		var c = make(map[string]any)

		err = json.Unmarshal(body, &c)

		if err != nil {
			log.Fatalln(err)
			return nil, 0, err
		}

		fmt.Println(c)

		return nil, 0, nil

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
