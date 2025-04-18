package processor

import (
	"errors"
	"os"

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
			generatedVersion, err := lib.GetGeneratedVersion(*rcTag)
			lib.CheckIfError(err)
			finalVersion := generatedVersion.BuildReleaseVersion()
			_, err = semver.NewVersion(finalVersion)
			lib.CheckIfError(err)
			return &finalVersion, nil
		} else {
			releaseTag, err := lib.GetLatestReleaseTag(*input.RepoPath)
			lib.CheckIfError(err)
			generatedVersion, err := lib.GetGeneratedVersion(*releaseTag)
			lib.CheckIfError(err)
			err = generatedVersion.IncrementPatch()
			lib.CheckIfError(err)
			finalVersion := generatedVersion.BuildReleaseCandidateVersion(jobRunId)
			_, err = semver.NewVersion(finalVersion)
			lib.CheckIfError(err)
			return &finalVersion, nil
		}
	}
	_, gitlabEnv := os.LookupEnv("GITLAB_CI")
	if gitlabEnv {
		return nil, errors.New("Non GitHub implementation is not supported right now.")
	}

	return nil, errors.New("Non GitHub implementation is not supported right now.")
}
