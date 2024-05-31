// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package sample

// Injectors from injector.go:

// to generate code run: wire gen module_name/package_name
// package name should be where the injector file lives
func InitializeService() (*SampleService, error) {
	sampleRepository := NewSampleRepository()
	sampleService, err := NewSampleService(sampleRepository)
	if err != nil {
		return nil, err
	}
	return sampleService, nil
}