package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	configPath string
	outputPath string
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyser les logs depuis un fichier de configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Analyze lancé avec config =", configPath, " output =", outputPath)
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "config.json", "Chemin du fichier de configuration JSON")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "report.json", "Chemin du fichier de sortie JSON")

}
