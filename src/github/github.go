package github

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

func GetPullRequestLabelsWithCommitHash(repo string, commitSha string, token string) ([]*github.Label, error) {
	client, err := GetClient(token)
	if err != nil {
		return nil, err
	}
	repoSplits := strings.Split(repo, "/")
	repoOwner := repoSplits[0]
	repoName := repoSplits[1]
	pr, _, err := client.PullRequests.ListPullRequestsWithCommit(context.Background(), repoOwner, repoName, commitSha, &github.ListOptions{})
	if err != nil {
		return nil, err
	}

	if len(pr) == 0 {
		return nil, errors.New("No PR found matching commit found")
	}
	return pr[0].Labels, nil
}

func ValidatePermissions(repo string, token string) error {
	client, err := GetClient(token)
	if err != nil {
		return err
	}
	repoSplits := strings.Split(repo, "/")
	_, _, err = client.PullRequests.List(context.Background(), repoSplits[0], repoSplits[1], &github.PullRequestListOptions{
		ListOptions: github.ListOptions{
			PerPage: 1,
		},
	})
	if err != nil {
		return errors.New("Unable to access PullRequests. Please ensure you have assigned 'pull-requests: read; permission to the token")
	}
	_, _, err = client.Repositories.ListCommits(context.Background(), repoSplits[0], repoSplits[1], &github.CommitsListOptions{
		ListOptions: github.ListOptions{
			PerPage: 1,
		},
	})
	if err != nil {
		return errors.New("Unable to access Commits. Please ensure you have assigned 'contents: read; permission to the token")
	}
	return nil
}
