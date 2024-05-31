package test

import (
	"testing"

	"github.com/mnabil1718/go-restful-api/sample"
)

func TestSampleDependencyInjector(t *testing.T) {
	_, err := sample.InitializeService()
	if err != nil {
		panic(err.Error())
	}
}
