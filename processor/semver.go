package processor

import (
	"fmt"
	"os"

	"github.com/alacrity-sg/build-version/lib"
)

type ProcessorInput struct {
	RepoPath *string
	Token    *string
}

func ProcessSemver(input *ProcessorInput) (*string, error) {
	_, githubEnv := os.LookupEnv("GITHUB_CI")
	var labels []string
	if githubEnv == true {
		token := os.Getenv("GITHUB_TOKEN")
		repo := os.Getenv("GITHUB_REPOSITORY")
		if *input.Token == "" {
			token = *input.Token
		}
		commitHash, err := lib.GetLastCommit(*input.RepoPath)
		if err != nil {
			return nil, err
		}
		prId, err := lib.GetPullRequestIdWithCommitHash(repo, commitHash, token)
		if err != nil {
			return nil, err
		}
		labels, err = lib.GetLabelsFromPullRequest(repo, prId, token)
	}

	for _, value := range labels {
		fmt.Println(value)
	}
	result := ""
	return &result, nil
}
