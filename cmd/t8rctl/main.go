package main

import (
	"log"

	"github.com/ylallemant/t8rctl/pkg/cli"
)

func main() {
	if err := cli.Command().Execute(); err != nil {
		log.Fatalf("error during execution: %v", err)
	}
}
