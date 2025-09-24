package config

import (
	"encoding/json"
	"os"
)

type LogTarget struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

func LoadConfig(path string) ([]LogTarget, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var logs []LogTarget
	if err := json.Unmarshal(data, &logs); err != nil {
		return nil, err
	}
	return logs, nil
}
