package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/joho/godotenv"
)

func main() {
	envFile := ".env"

	if err := godotenv.Load(envFile); err != nil {
		fmt.Fprintf(os.Stderr, "failed to load env file: %v\n", err)
	}

	stmts, err := gormschema.New("postgres").Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
