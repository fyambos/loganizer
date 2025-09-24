package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

type LogTarget struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

type Result struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

var (
	configPath string
	outputPath string
)

func runAnalyze(configPath, outputPath string) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Erreur lecture config:", err)
		return
	}
	var logs []LogTarget
	if err := json.Unmarshal(data, &logs); err != nil {
		fmt.Println("Erreur parsing JSON:", err)
		return
	}
	fmt.Println("Config chargée avec", len(logs), "entrées")

	// inspiré d'un exemple goroutine + waitgroup
	results := make([]Result, 0, len(logs)) // slice pour stocker les résultats
	out := make(chan Result)                // canal pour recevoir les résultats

	var wg sync.WaitGroup // waitgroup pour attendre les goroutines

	// lancer une goroutine par log
	for _, t := range logs {
		wg.Add(1)
		go func(target LogTarget) {
			defer wg.Done()
			// ici on envoie un résultat simulé dans le canal
			out <- Result{
				LogID:    target.ID,
				FilePath: target.Path,
				Status:   "OK",
				Message:  "Analyse simulée.",
			}
		}(t)
	}

	// fermeture du canal quand toutes les goroutines ont fini
	go func() {
		wg.Wait()
		close(out)
	}()

	// boucle de réception des résultats
	for r := range out {
		results = append(results, r)
		fmt.Println("Résultat pour", r.LogID, ":", r.Status)
	}
}

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyser les logs depuis un fichier de configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Analyze lancé avec config =", configPath, " output =", outputPath)
		runAnalyze(configPath, outputPath)
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "config.json", "Chemin du fichier de configuration JSON")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "report.json", "Chemin du fichier de sortie JSON")

}
