package providers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/lucassimon/desafio-multithread-hub-cep/internal/dto"
)

type ViaCep struct {
	BaseUrl string
}

type ViaCepOutput struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func NewViaCep() *ViaCep {
	return &ViaCep{
		BaseUrl: "http://viacep.com.br/",
	}
}

func (c *ViaCep) Search(ch chan<- *dto.CepOutput, cep string) error {
	url := c.BaseUrl + "ws/" + cep + "/json"

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

	var output ViaCepOutput
	err = json.Unmarshal(body, &output)
	if err != nil {
		return err
	}

	cep_translated := c.translate(output)
	ch <- cep_translated

	return nil
}

func (c *ViaCep) translate(output ViaCepOutput) *dto.CepOutput {
	return &dto.CepOutput{
		Code:     output.Cep,
		State:    output.Uf,
		City:     output.Localidade,
		District: output.Bairro,
		Address:  output.Logradouro,
	}
}
