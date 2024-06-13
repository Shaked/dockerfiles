package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sethvargo/go-envconfig"

	"github.com/Shaked/dockerfiles/go/internal/util"
)

func main() {
	var c struct {
		ProjectID string `env:"PROJECT_ID"`
	}

	if err := envconfig.Process(context.TODO(), &c); err != nil {
		log.Fatalf("Failed to process env vars: %v", err)
	}

	fmt.Printf("ProjectID: %s\n", c.ProjectID)
	fmt.Println(util.Hello())
}
