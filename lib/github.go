package lib

import (
	"context"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/v71/github"
)

func GetClient() (*github.Client, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, errors.New("GITHUB_TOKEN env variable is not set. Please set and try again")
	}
	client := github.NewClient(nil).WithAuthToken()
	return client, nil
}

func GetLabelsFromPullRequest(prId string, token string) ([]*github.Label, error) {
	client, err := GetClient()
	if err != nil {
		return nil, err
	}
	repositoryName := os.Getenv("GITHUB_REPOSITORY")
	repoSplits := strings.Split(repositoryName, "/")
	owner := repoSplits[0]
	repo := repoSplits[1]
	prNumber, err := strconv.Atoi(prId)
	pr, _, err := client.PullRequests.Get(context.Background(), owner, repo, prNumber)
	labels := pr.Labels
	return labels, nil
}
