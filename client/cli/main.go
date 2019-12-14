package main

import (
	"fmt"
	"github.com/colin-axner/sangam/client/cli/cmd"
	"os"
)

func main() {
	rootCmd := cmd.RootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
