package main

import (
	"encoding/json/v2"
	"runtime"
)

// These variables are set by GoReleaser during the build.
var (
	commit      string
	date        string
	projectName string
	version     string
)

// formatVersionJSON formats version information as a JSON string.
func formatVersionJSON() (string, error) {
	v, err := json.Marshal(
		struct {
			ProjectName string
			Version     string
			Commit      string
			BuildDate   string
			GoVersion   string
		}{
			projectName,
			version,
			commit,
			date,
			runtime.Version(),
		})
	if err != nil {
		return "", err
	}
	return string(v), nil
}
