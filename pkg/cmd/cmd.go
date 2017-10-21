package cmd

import (
	"log"

	"github.com/MakeCents-NYC/artland/pkg/artland"
	"github.com/spf13/cobra"
)

const (
	// TODO: Finish these
	cliExplain = `TODO`
	cliLong    = `TODO`
	cliShort   = `TODO`
	cliUse     = `TODO`
)

// NewCmdArtLand initializes the Cobra stuff for out program
func NewCmdArtLand() *cobra.Command {
	command := &cobra.Command{
		Use:   cliUse,
		Short: cliShort,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.RunE(cmd, args); err != nil {
				log.Fatal(err)
			}
		},
		RunE: run,
	}

	// TODO: CLI args go here

	return command
}

// run is our "main". It creates a gateway and runs it. In the future it would
// pass in any needed info from Cobra
func run(cmd *cobra.Command, args []string) error {
	al := artland.NewArtLand()

	return al.Serve()
}
