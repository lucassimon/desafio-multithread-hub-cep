package providers

import (
	"github.com/lucassimon/desafio-multithread-hub-cep/internal/dto"
)

type SearchCepInterface interface {
	Search(cep string) (dto.CepOutput, error)
}
