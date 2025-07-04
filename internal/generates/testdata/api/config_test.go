package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg, err := ReadConfig("input.yml")
	if err != nil {
		t.Fatal(fmt.Errorf("reading config: %w", err))
	}

	fmt.Println(cfg.Gateways.Public.Services.Objectives.Endpoints.Create)

	for ep, details := range cfg.Gateways.Public.Services.Objectives.Endpoints.Fields() {
		fmt.Printf("%s => %s\n", ep, details)
	}
}
