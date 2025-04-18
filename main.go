package main

import (
	"flag"
	"fmt"

	"github.com/alacrity-sg/build-version/lib"
	"github.com/alacrity-sg/build-version/processor"
)

func main() {
	repoPtr := flag.String("repo-path", ".", "Local folder path to git repository")
	tokenPtr := flag.String("token", "", "Token to use to communicate with Git APIs.")
	outputFilePtr := flag.String("output-file", "./build-version.env", "Local file path to create env compliant file")
	flag.Parse()
	input := processor.ProcessorInput{
		RepoPath:       repoPtr,
		Token:          tokenPtr,
		OutputFilePath: outputFilePtr,
	}
	version, err := processor.ProcessSemver(&input)
	if err != nil {
		panic(err)
	}
	err = lib.WriteToFile(*version, *input.OutputFilePath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully generated version.")
	fmt.Printf("Build Version for this run is %s", *version)
}
