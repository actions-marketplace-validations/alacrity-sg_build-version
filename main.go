package main

import (
	"flag"
	"fmt"
	"github.com/alacrity-sg/build-version/src/github"
	"os"

	"github.com/alacrity-sg/build-version/src/lib"
	"github.com/alacrity-sg/build-version/src/processor"
)

func main() {
	repoPtr := flag.String("repo-path", ".", "Local folder path to git repository")
	tokenPtr := flag.String("token", "", "Token to use to communicate with Git APIs.")
	outputFilePtr := flag.String("output-file", "./build-version.env", "Local file path to create env compliant file")
	incrementTypePtr := flag.String("increment-type", "", "Specify increment type. Accepts 'major', 'minor' or 'patch'")
	offlineModePtr := flag.Bool("offline", false, "Whether to enable offline mode that disables network communication")
	flag.Parse()
	input := processor.ProcessorInput{
		RepoPath:       *repoPtr,
		Token:          *tokenPtr,
		OutputFilePath: *outputFilePtr,
		IncrementType:  *incrementTypePtr,
		OfflineMode:    *offlineModePtr,
	}

	_, githubEnv := os.LookupEnv("GITHUB_ACTIONS")
	if input.OfflineMode {
		fmt.Println("Offline mode is enabled. Please note that all external features such as communicating with Git services will be disabled")
	} else {
		if githubEnv {
			repository := os.Getenv("GITHUB_REPOSITORY")
			err := github.ValidatePermissions(repository, input.Token)
			if err != nil {
				fmt.Println("[GitHub] Error validating permissions:", err)
				os.Exit(1)
			}
		}
	}
	version, err := input.ProcessSemver()
	if err != nil {
		panic(err)
	}
	err = lib.WriteToFile(*version, input.OutputFilePath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully generated version. Generated result file is %s", *outputFilePtr)
	fmt.Printf("Build Version for this run is %s", *version)
}
