//go:build wireinject
// +build wireinject

package sample

import "github.com/google/wire"

// to generate code run: wire gen module_name/package_name
// package name should be where the injector file lives
func InitializeService(isError bool) (*SampleService, error) {
	wire.Build(NewSampleRepository, NewSampleService)
	return nil, nil
}
