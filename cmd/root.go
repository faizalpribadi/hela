package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "root-command",
	Short: "root command",
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal("cannot parser command line")
		os.Exit(1)
	}
}
