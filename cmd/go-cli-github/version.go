package main

import (
	"encoding/json"
	"fmt"
)

// VersionCmd represents the `version` command.
type VersionCmd struct{}

// Run the Version command.
func (*VersionCmd) Run() error {
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
			goVersion,
		})
	if err != nil {
		return err
	}
	_, err = fmt.Println(string(v))
	return err
}
