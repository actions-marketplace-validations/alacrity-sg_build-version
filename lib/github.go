package lib

import (
	"context"
	"errors"
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

func GetLabelsFromPullRequest(repo string, prId *int64, token string) ([]string, error) {
	client, err := GetClient(token)
	if err != nil {
		return nil, err
	}
	repoSplits := strings.Split(repo, "/")
	repoOwner := repoSplits[0]
	repoName := repoSplits[1]
	pr, _, err := client.PullRequests.Get(context.Background(), repoOwner, repoName, int(*prId))
	githubLabels := pr.Labels
	var labels []string
	for _, value := range githubLabels {
		labels = append(labels, value.GetName())
	}
	return labels, nil
}

func GetPullRequestIdWithCommitHash(repo string, commitSha *string, token string) (*int64, error) {
	client, err := GetClient(token)
	if err != nil {
		return nil, err
	}
	repoSplits := strings.Split(repo, "/")
	repoOwner := repoSplits[0]
	repoName := repoSplits[1]
	pr, _, err := client.PullRequests.ListPullRequestsWithCommit(context.Background(), repoOwner, repoName, *commitSha, &github.ListOptions{})
	if err != nil {
		return nil, err
	}

	if len(pr) == 0 {
		return nil, errors.New("No PR found matching commit found")
	}
	return pr[0].ID, nil
}
