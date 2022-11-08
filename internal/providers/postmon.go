package providers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/lucassimon/desafio-multithread-hub-cep/internal/dto"
)

type Postmon struct {
	BaseUrl string
}

type PostmonOutput struct {
	Bairro     string `json:"bairro"`
	Cidade     string `json:"cidade"`
	Logradouro string `json:"logradouro"`
	Cep        string `json:"cep"`
	Estado     string `json:"estado"`
}

func NewPostmonCep() *Postmon {
	return &Postmon{
		BaseUrl: "https://api.postmon.com.br",
	}
}

func (c *Postmon) Search(ch chan<- *dto.CepOutput, cep string) error {
	url := c.BaseUrl + "/v1/cep/" + cep

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

	var output PostmonOutput
	err = json.Unmarshal(body, &output)
	if err != nil {
		return err
	}

	cep_translated := c.translate(output)
	ch <- cep_translated

	return nil
}

func (c *Postmon) translate(output PostmonOutput) *dto.CepOutput {
	return &dto.CepOutput{
		Code:     output.Cep,
		State:    output.Estado,
		City:     output.Cidade,
		District: output.Bairro,
		Address:  output.Logradouro,
	}
}
