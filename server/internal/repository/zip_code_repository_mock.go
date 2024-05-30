package repository

import (
	"github.com/stretchr/testify/mock"
	"github.com/victorbrugnolo/golang-temp-zipcode/internal/entity"
)

type ZipCodeRepositoryMock struct {
	mock.Mock
}

func (m *ZipCodeRepositoryMock) GetZipcodeData(zipcode string) (*entity.ZipcodeDataResponse, *entity.ErrorResponse) {
	args := m.Called(zipcode)

	if args.Get(1) != nil {
		return nil, args.Get(1).(*entity.ErrorResponse)
	}

	return args.Get(0).(*entity.ZipcodeDataResponse), nil
}
