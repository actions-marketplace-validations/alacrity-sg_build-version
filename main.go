package main

import (
	"flag"

	"github.com/alacrity-sg/build-version/processor"
)

func main() {
	repoPtr := flag.String("repo-path", ".", "Local folder path to git repository")
	tokenPtr := flag.String("token", "", "Token to use to communicate with Git APIs.")
	flag.Parse()
	input := processor.ProcessorInput{
		RepoPath: repoPtr,
		Token:    tokenPtr,
	}
	processor.ProcessSemver(&input)
}
