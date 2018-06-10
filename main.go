package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var command = &cobra.Command{
		Use:   "hela",
		Short: "Hela automation tool for service health checking",
		run: func(cmd *cobra.Command, args []string) {

		},
	}

	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
