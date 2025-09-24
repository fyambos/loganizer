package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "loganizer",
	Short: "loganizer est un outil pour analyser des logs.",
	Long:  `Un outil cli pour analyser des logs en parallele et exporter les résultats.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur: %v\n", err)
		os.Exit(1)
	}
}

func init() {

}
