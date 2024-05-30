package repository

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/victorbrugnolo/golang-temp-zipcode/internal/entity"
)

type ZipCodeRepository struct{}

func NewZipCodeRepository() *ZipCodeRepository {
	return &ZipCodeRepository{}
}

func (r *ZipCodeRepository) GetZipcodeData(zipcode string) (*entity.ZipcodeDataResponse, *entity.ErrorResponse) {
	url := "http://viacep.com.br/ws/" + zipcode + "/json"
	resp, err := http.Get(url)

	if err != nil {
		return nil, entity.BuildErrorResponse(http.StatusInternalServerError, err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, entity.BuildErrorResponse(http.StatusInternalServerError, err)
	}

	var viaCepResponse entity.ZipcodeDataResponse
	err = json.Unmarshal(body, &viaCepResponse)

	if err != nil {
		return nil, entity.BuildErrorResponse(http.StatusInternalServerError, err)

	}

	if viaCepResponse.Error {
		return nil, entity.BuildErrorResponse(http.StatusNotFound, errors.New("can not find zipcode"))
	}

	return &viaCepResponse, nil
}
