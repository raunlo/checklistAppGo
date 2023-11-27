package main

import (
	"com.rlohmus.checklist/internal/infra"
	"go.uber.org/config"
	"os"
	"path/filepath"
)

func main() {
	applicationConfig := getApplicationConfig("./cmd/application.yaml")
	application := infra.Init(applicationConfig)
	if err := application.StartApplication(); err != nil {
		panic(err)
	}
}

func getApplicationConfig(configPath string) infra.ApplicationConfiguration {
	var applicationConfiguration infra.ApplicationConfiguration
	if filePath, err := filepath.Abs(configPath); err != nil {
		panic(err)
	} else if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else if provider, err := config.NewYAML(config.Expand(os.LookupEnv), config.Source(file)); err != nil {
		panic(err)
	} else if err := provider.Get("").Populate(&applicationConfiguration); err != nil {
		panic(err)
	} else {
		return applicationConfiguration
	}
}
