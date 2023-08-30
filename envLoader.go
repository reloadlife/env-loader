package envLoader

import (
	"fmt"
	"os"
	"strings"
)

var (
	WorkingDirectory string
	Path             = "/.env"
)

func getEnvFile() string {
	if WorkingDirectory == "" {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		WorkingDirectory = wd
	}

	envFile := WorkingDirectory + Path
	fmt.Printf("[envLoader] -> Loading from %s", envFile)
	_, err := os.Stat(envFile)
	if err != nil {
		panic(err)
	}
	fileContent, err := os.ReadFile(envFile)
	if err != nil {
		panic(err)
	}

	return string(fileContent)
}

func Load() {
	fileContent := getEnvFile()
	environments := strings.Split(fileContent, "\n")
	for _, env := range environments {
		if env != "" {
			e := strings.Split(env, "=")
			if err := os.Setenv(e[0], e[1]); err != nil {
				panic(err)
			}
		}
	}
}
