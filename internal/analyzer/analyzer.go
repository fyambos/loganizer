package analyzer

import (
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/fyambos/loganizer/internal/config"
)

type Result struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

func AnalyzeLogs(logs []config.LogTarget) []Result {
	results := make([]Result, 0, len(logs))
	out := make(chan Result)
	var wg sync.WaitGroup

	for _, t := range logs {
		wg.Add(1)
		go func(tt config.LogTarget) {
			defer wg.Done()
			out <- analyzeOne(tt)
		}(t)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for r := range out {
		results = append(results, r)
	}
	return results
}

func analyzeOne(t config.LogTarget) Result {
	if _, err := os.Stat(t.Path); err != nil {
		return Result{
			LogID:        t.ID,
			FilePath:     t.Path,
			Status:       "FAILED",
			Message:      "Fichier introuvable.",
			ErrorDetails: err.Error(),
		}
	}
	delay := rand.Intn(151) + 50 // 50–200 ms
	time.Sleep(time.Duration(delay) * time.Millisecond)

	return Result{
		LogID:        t.ID,
		FilePath:     t.Path,
		Status:       "OK",
		Message:      "Analyse terminée avec succès.",
		ErrorDetails: "",
	}
}
