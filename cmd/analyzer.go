package cmd

import (
	"fmt"

	"github.com/fyambos/loganizer/internal/analyzer"
	"github.com/fyambos/loganizer/internal/config"
	"github.com/fyambos/loganizer/internal/reporter"

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

		// charger config
		logs, err := config.LoadConfig(configPath)
		if err != nil {
			fmt.Println("Erreur lecture config:", err)
			return
		}
		fmt.Println("Config chargée avec", len(logs), "entrées")

		results := analyzer.AnalyzeLogs(logs) //lancer analyse

		for _, r := range results {
			fmt.Printf("[%s] %s (%s): %s\n", r.Status, r.LogID, r.FilePath, r.Message) // affichage resultats
			if r.Status == "FAILED" {
				fmt.Println("  Détails erreur:", r.ErrorDetails) // détails erreur si echec
			}
		}

		// exporter en json
		if err := reporter.WriteReport(outputPath, results); err != nil {
			fmt.Println("Erreur écriture rapport:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "config.json", "Chemin du fichier de configuration JSON")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "report.json", "Chemin du fichier de sortie JSON")

}
