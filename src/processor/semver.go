package processor

import (
	"errors"
	"os"

	"github.com/Masterminds/semver"
	"github.com/alacrity-sg/build-version/src/generator"
	"github.com/alacrity-sg/build-version/src/git"
	"github.com/alacrity-sg/build-version/src/lib"
)

type ProcessorInput struct {
	RepoPath       string
	Token          string
	OutputFilePath string
	IncrementType  string
	OfflineMode    bool
}

func ProcessSemver(input *ProcessorInput) (*string, error) {
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
			incrementType, err := lib.GetIncrementType(input.IncrementType, input.OfflineMode)
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
