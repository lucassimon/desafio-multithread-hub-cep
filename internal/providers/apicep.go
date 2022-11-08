package providers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/lucassimon/desafio-multithread-hub-cep/internal/dto"
)

type ApiCep struct {
	BaseUrl string
}

type ApiCepOutput struct {
	Code     string `json:"code"`
	Address  string `json:"address"`
	District string `json:"district"`
	City     string `json:"city"`
	State    string `json:"state"`
}

func NewApiCep() *ApiCep {
	return &ApiCep{
		BaseUrl: "https://cdn.apicep.com/",
	}
}

func (c *ApiCep) Search(ch chan<- *dto.CepOutput, cep string) error {
	url := c.BaseUrl + "/file/apicep/" + cep + ".json"

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		close(ch)
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	var output ApiCepOutput
	err = json.Unmarshal(body, &output)
	if err != nil {
		return err
	}

	cep_translated := c.translate(output)
	ch <- cep_translated

	return nil
}

func (c *ApiCep) translate(output ApiCepOutput) *dto.CepOutput {
	return &dto.CepOutput{
		Code:     output.Code,
		State:    output.State,
		City:     output.City,
		District: output.District,
		Address:  output.Address,
	}
}
