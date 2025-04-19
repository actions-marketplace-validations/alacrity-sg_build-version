package processor

import (
	"errors"
	"fmt"
	"os"
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
		// token := os.Getenv("GITHUB_TOKEN")
		// repo := os.Getenv("GITHUB_REPOSITORY")
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

	refName := os.Getenv("GITHUB_REF_NAME")
	if refName != "main" {
		return &defaultIncrement, nil
	}

	repo := os.Getenv("GITHUB_REPOSITORY")
	commit, err := git.GetLastCommit(input.RepoPath)
	if err != nil {
		return nil, err
	}
	labels, err := github.GetPullRequestLabelsWithCommitHash(repo, *commit, input.Token)
	if err != nil {
		return nil, err
	}
	for _, label := range labels {
		lowerLabel := strings.ToLower(*label.Name)
		if lowerLabel == "major" {
			defaultIncrement = "major"
			break
		}
		if lowerLabel == "minor" && defaultIncrement == "patch" {
			defaultIncrement = "minor"
		}
		if lowerLabel == "patch" && defaultIncrement == "patch" {
			defaultIncrement = "patch"
		}
	}
	return &defaultIncrement, nil
}
