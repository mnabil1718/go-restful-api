package test

import (
	"testing"

	"github.com/mnabil1718/go-restful-api/sample"
	"github.com/stretchr/testify/assert"
)

func TestSampleDependencyInjector(t *testing.T) {
	_, err := sample.InitializeService(true)
	if err != nil {
		assert.NotNil(t, err)
		assert.Equal(t, "failed creating service", err.Error())
	}
}
