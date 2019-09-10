package currency

import (
	"errors"
	"io"
	"net/http"
)

// Symbol is used to store and structured data from server.
type Symbol struct {
	Status bool              `json:"success"`
	Data   map[string]string `json:"symbols"`
}

// History is used to store and structured data from server.
type History struct {
	Status     bool               `json:"success"`
	Historical bool               `json:"historical"`
	Date       string             `json:"date"`
	Base       string             `json:"base"`
	Rates      map[string]float32 `json:"rates"`
}

// SetBase is used to store and structured data from server.
type SetBase struct {
	Status bool               `json:"success"`
	Date   string             `json:"date"`
	Base   string             `json:"base"`
	Rates  map[string]float32 `json:"rates"`
}

// Convert is used to store and structured data from server.
type Convert struct {
	Status bool    `json:"success"`
	Date   string  `json:"date"`
	Result float32 `json:"result"`

	Query struct {
		From   string  `json:"from"`
		To     string  `json:"to"`
		Amount float32 `json:"amount"`
	} `json:"query"`

	Info struct {
		Timestamp uint32  `json:"timestamp"`
		Rate      float32 `json:"rate"`
	} `json:"info"`
}

// FluctOrTs is used to store and structured data from server.
type FluctOrTs struct {
	Status    bool                          `json:"success"`
	StartData string                        `json:"start_date"`
	EndDate   string                        `json:"end_date"`
	Base      string                        `json:"base"`
	Rates     map[string]map[string]float32 `json:"rates"`
}

var errServer = errors.New("response error from server or subscribe plan is restricted")
var errDecode = errors.New("error while decoding data")

func (p *Parser) get() (io.ReadCloser, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("POST", p.base, nil)
	if err != nil {
		return nil, errors.New("error while connect to server")
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.New("error while server write data to client")
	}
	return response.Body, nil
}

func toString(sy ...string) (string, error) {
	var rURL string
	if len(sy) == 0 {
		return rURL, errors.New("error parsing symbols")
	}
	for i, s := range sy {
		if i == 0 {
			rURL += s
		} else {
			rURL += "," + s
		}
	}
	return rURL, nil
}
