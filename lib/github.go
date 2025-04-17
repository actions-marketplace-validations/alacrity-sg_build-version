package lib

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/google/go-github/v71/github"
)

func GetClient(token string) (*github.Client, error) {
	if token == "" {
		return nil, errors.New("GITHUB_TOKEN env variable is not set. Please set and try again")
	}
	client := github.NewClient(nil).WithAuthToken(token)
	return client, nil
}

func GetLabelsFromPullRequest(repo string, prId string, token string) ([]*github.Label, error) {
	client, err := GetClient(token)
	if err != nil {
		return nil, err
	}
	repoSplits := strings.Split(repo, "/")
	repoOwner := repoSplits[0]
	repoName := repoSplits[1]
	prNumber, err := strconv.Atoi(prId)
	pr, _, err := client.PullRequests.Get(context.Background(), repoOwner, repoName, prNumber)
	labels := pr.Labels
	return labels, nil
}

func GetPullRequestWithCommitHash(repo string, commitSha, token string)([]*github.Label, error) {
	client, err := GetClient(token)
	if err != nil {
		return nil, err
	}
	repoSplits := strings.Split(repo, "/")
	repoOwner := repoSplits[0]
	repoName := repoSplits[1]
	repo.
}
