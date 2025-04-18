package processor

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/alacrity-sg/build-version/lib"
)

type ProcessorInput struct {
	RepoPath       *string
	Token          *string
	OutputFilePath *string
}

func ProcessSemver(input *ProcessorInput) (*string, error) {
	_, githubEnv := os.LookupEnv("GITHUB_CI")
	if githubEnv {
		// token := os.Getenv("GITHUB_TOKEN")
		// repo := os.Getenv("GITHUB_REPOSITORY")
		refName := os.Getenv("GITHUB_REF_NAME")
		jobRunId := os.Getenv("GITHUB_RUN_ID")
		if refName == "main" {
			// Process RC to become release
			rcTag, err := lib.GetLatestRCTag(*input.RepoPath)
			lib.CheckIfError(err)
			rcTagSplits := strings.Split(*rcTag, "-")
			major := rcTagSplits[0]
			minor := rcTagSplits[1]
			patch := rcTagSplits[2]
			newVersion := fmt.Sprintf("%s.%s.%s", major, minor, patch)
			_, err = semver.NewVersion(newVersion)
			lib.CheckIfError(err)
			return &newVersion, nil
		} else {
			releaseTag, err := lib.GetLatestReleaseTag(*input.RepoPath)
			lib.CheckIfError(err)
			newVersion := fmt.Sprintf("%s-rc.%s", *releaseTag, jobRunId)
			_, err = semver.NewVersion(newVersion)
			lib.CheckIfError(err)
			return &newVersion, nil
		}
	}
	_, gitlabEnv := os.LookupEnv("GITLAB_CI")
	if gitlabEnv {
		return nil, errors.New("Non GitHub implementation is not supported right now.")
	}

	return nil, errors.New("Non GitHub implementation is not supported right now.")
}
