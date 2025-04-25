package processor

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/alacrity-sg/build-version/src/generator"
	"github.com/alacrity-sg/build-version/src/git"
	"github.com/alacrity-sg/build-version/src/github"
	"github.com/alacrity-sg/build-version/src/lib"
)

type ProcessorInput struct {
	RepoPath       string
	Token          string
	OutputFilePath string
	IncrementType  string
	OfflineMode    bool
}

func (input *ProcessorInput) ProcessSemver() (*string, error) {
	_, githubEnv := os.LookupEnv("GITHUB_ACTIONS")
	if githubEnv {
		refName := os.Getenv("GITHUB_REF_NAME")
		jobRunId := os.Getenv("GITHUB_RUN_ID")
		if refName == "main" {
			// Process RC to become release
			rcTag, err := git.GetLatestRCTag(input.RepoPath)
			lib.CheckIfError(err)
			generatedVersion, err := generator.GetGeneratedVersion(*rcTag)
			lib.CheckIfError(err)
			finalVersion := generatedVersion.BuildReleaseVersion()
			_, err = semver.NewVersion(finalVersion)
			lib.CheckIfError(err)
			return &finalVersion, nil
		} else {
			releaseTag, err := git.GetLatestReleaseTag(input.RepoPath)
			lib.CheckIfError(err)
			generatedVersion, err := generator.GetGeneratedVersion(*releaseTag)
			lib.CheckIfError(err)
			incrementType, err := input.parseIncrementType()
			lib.CheckIfError(err)
			if *incrementType == "major" {
				err = generatedVersion.IncrementMajor()
				lib.CheckIfError(err)
			} else if *incrementType == "minor" {
				err = generatedVersion.IncrementMinor()
				lib.CheckIfError(err)
			} else {
				err = generatedVersion.IncrementPatch()
				lib.CheckIfError(err)
			}
			finalVersion := generatedVersion.BuildReleaseCandidateVersion(jobRunId)
			_, err = semver.NewVersion(finalVersion)
			lib.CheckIfError(err)
			return &finalVersion, nil
		}
	} else {
		return nil, errors.New("Non GitHub implementation is not supported right now.")
	}
}

func (input *ProcessorInput) parseIncrementType() (*string, error) {
	defaultIncrement := "patch"
	if input.IncrementType != "" {
		lowercaseIncrementType := strings.ToLower(input.IncrementType)
		if lowercaseIncrementType == "major" || lowercaseIncrementType == "minor" || lowercaseIncrementType == "patch" {
			defaultIncrement = lowercaseIncrementType
		} else {
			return nil, errors.New(fmt.Sprintf("Expected IncrementType to be 'major', 'minor' or 'patch' but received '%s'", lowercaseIncrementType))
		}
	}
	if input.OfflineMode {
		return &defaultIncrement, nil
	}

	repo := os.Getenv("GITHUB_REPOSITORY")
	refName := os.Getenv("GITHUB_REF_NAME")
	refNameSplits := strings.Split(refName, "/")
	if len(refNameSplits) != 2 {
		return &defaultIncrement, nil
	}
	prId, err := strconv.Atoi(refNameSplits[0])
	if err != nil {
		return &defaultIncrement, nil
	}
	labels, err := github.GetLabelsFromPullRequest(repo, prId, input.Token)
	if err != nil {
		return nil, err
	}
	for _, label := range labels {
		if label == "major" {
			defaultIncrement = label
			break
		}
		if label == "minor" && defaultIncrement == "patch" {
			defaultIncrement = label
		}
		if label == "patch" && defaultIncrement == "patch" {
			defaultIncrement = label
		}
	}
	return &defaultIncrement, nil
}
