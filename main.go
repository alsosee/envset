// GitHub Action that sets environment variable.
// Can be used in cases when value may contain colons or double quotes.
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("::error::%v", err)
		os.Exit(1)
	}
}

func run() error {
	name := os.Getenv("INPUT_NAME")
	value := os.Getenv("INPUT_VALUE")

	if name == "" {
		return fmt.Errorf("name is required")
	}

	if value == "" {
		return fmt.Errorf("value is required")
	}

	return writeEnv(name, value)
}

func writeEnv(name, value string) error {
	githubEnv := formatOutput(name, value)
	if githubEnv == "" {
		return nil
	}

	path := os.Getenv("GITHUB_ENV")

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf(
			"failed to open result file %q: %v. "+
				"If you are using self-hosted runners "+
				"make sure they are updated to version 2.297.0 or greater",
			path,
			err,
		)
	}
	defer f.Close()

	if _, err = f.WriteString(githubEnv); err != nil {
		return fmt.Errorf("failed to write result to file %q: %v", path, err)
	}

	return nil
}

func formatOutput(name, value string) string {
	if value == "" {
		return ""
	}

	// if value contains new line, use multiline format
	if bytes.ContainsRune([]byte(value), '\n') {
		return fmt.Sprintf("%s<<ENV\n%s\nENV", name, value)
	}

	return fmt.Sprintf("%s=%s", name, value)
}
