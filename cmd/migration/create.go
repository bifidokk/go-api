package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/bifidokk/go-api/internal/entity"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(
		&entity.Note{},
		&entity.User{},
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}

	io.WriteString(os.Stdout, stmts)
}
