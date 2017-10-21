package main

import (
	"log"

	"github.com/MakeCents-NYC/artland/pkg/cmd"
)

func main() {
	al := cmd.NewCmdArtLand()
	if err := al.Execute(); err != nil {
		log.Fatal(err)
	}
}
