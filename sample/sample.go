package sample

import "errors"

type SampleRepository struct {
	Error bool
}

func NewSampleRepository(isError bool) *SampleRepository {
	return &SampleRepository{isError}
}

type SampleService struct {
	*SampleRepository
}

func NewSampleService(repository *SampleRepository) (*SampleService, error) {
	if repository.Error {
		return nil, errors.New("failed creating service")
	}

	return &SampleService{repository}, nil
}
