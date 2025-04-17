package processor

import (
	"os"

	"github.com/alacrity-sg/build-version/lib"
)

type ProcessorInput struct {
	RepoPath *string
	Token    *string
}

func ProcessSemver(input *ProcessorInput) (*string, error) {
	_, githubEnv := os.LookupEnv("GITHUB_CI")
	if githubEnv == true {
		token := os.Getenv("GITHUB_TOKEN")
		if *input.Token == "" {
			token = *input.Token
		}
		commitHash, err := lib.GetLastCommit(*input.RepoPath)
		if err != nil {
			return nil, err
		}

	}
}
