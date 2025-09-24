package reporter

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fyambos/loganizer/internal/analyzer"
)

func WriteReport(path string, results []analyzer.Result) error {
	rep, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(path, rep, 0644); err != nil {
		return err
	}
	fmt.Println("Rapport écrit dans", path)
	return nil
}
