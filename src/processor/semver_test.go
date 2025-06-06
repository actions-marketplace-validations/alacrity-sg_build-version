package processor

import (
	"fmt"
	"testing"

	"github.com/alacrity-sg/build-version/src/bvtest"
	"github.com/alacrity-sg/build-version/src/lib"
	"github.com/go-git/go-git/v5"
)

func TestProcessSemverMainWithTagDefaults(t *testing.T) {
	t.Setenv("GITHUB_ACTIONS", "true")
	t.Setenv("GITHUB_REF_NAME", "main")
	dir, r, _, commit := bvtest.SetupRepo("main", t)
	tagOptions := &git.CreateTagOptions{
		Message: "Commit",
	}
	_, err := r.CreateTag("v1.0.0", commit.Hash, tagOptions)
	lib.CheckIfError(err)
	_, err = r.CreateTag("v1.0.1-rc.123", commit.Hash, tagOptions)
	lib.CheckIfError(err)
	input := ProcessorInput{
		RepoPath:    dir,
		OfflineMode: true,
	}
	version, err := input.ProcessSemver()
	lib.CheckIfError(err)
	expectedVersion := "1.0.1"
	if *version != expectedVersion {
		t.Errorf("Expected %s but received %s", *version, expectedVersion)
		t.Fail()
	}
}

func TestProcessSemverNonMainWithDefaults(t *testing.T) {
	t.Setenv("GITHUB_ACTIONS", "true")
	t.Setenv("GITHUB_REF_NAME", "feature/123")
	t.Setenv("GITHUB_RUN_ID", "abc")
	dir, _, _, _ := bvtest.SetupRepo("main", t)
	input := ProcessorInput{
		RepoPath:    dir,
		OfflineMode: true,
	}
	version, err := input.ProcessSemver()
	lib.CheckIfError(err)
	expectedVersion := "0.0.1-rc.abc"
	if *version != expectedVersion {
		t.Error(fmt.Sprintf("Expected %s but received %s", *version, expectedVersion))
		t.Fail()
	}
}

func TestProcessSemverNonMainWithExistingTagDefaults(t *testing.T) {
	t.Setenv("GITHUB_ACTIONS", "true")
	t.Setenv("GITHUB_REF_NAME", "feature/123")
	t.Setenv("GITHUB_RUN_ID", "abc")
	dir, r, _, commit := bvtest.SetupRepo("main", t)
	tagOptions := &git.CreateTagOptions{
		Message: "Commit",
	}
	_, err := r.CreateTag("v1.0.0", commit.Hash, tagOptions)
	lib.CheckIfError(err)
	input := ProcessorInput{
		RepoPath:    dir,
		OfflineMode: true,
	}
	version, err := input.ProcessSemver()
	lib.CheckIfError(err)
	expectedVersion := "1.0.1-rc.abc"
	if *version != expectedVersion {
		t.Error(fmt.Sprintf("Expected %s but received %s", *version, expectedVersion))
		t.Fail()
	}
}

func TestProcessSemverNonMainWithExistingTagPatch(t *testing.T) {
	t.Setenv("GITHUB_ACTIONS", "true")
	t.Setenv("GITHUB_REF_NAME", "feature/123")
	t.Setenv("GITHUB_RUN_ID", "abc")
	dir, r, _, commit := bvtest.SetupRepo("main", t)
	tagOptions := &git.CreateTagOptions{
		Message: "Commit",
	}
	_, err := r.CreateTag("v1.0.0", commit.Hash, tagOptions)
	lib.CheckIfError(err)
	input := ProcessorInput{
		RepoPath:      dir,
		IncrementType: "patch",
		OfflineMode:   true,
	}
	version, err := input.ProcessSemver()
	lib.CheckIfError(err)
	expectedVersion := "1.0.1-rc.abc"
	if *version != expectedVersion {
		t.Error(fmt.Sprintf("Expected %s but received %s", *version, expectedVersion))
		t.Fail()
	}
}

func TestProcessSemverNonMainWithExistingTagMinor(t *testing.T) {
	t.Setenv("GITHUB_ACTIONS", "true")
	t.Setenv("GITHUB_REF_NAME", "feature/123")
	t.Setenv("GITHUB_RUN_ID", "abc")
	dir, r, _, commit := bvtest.SetupRepo("main", t)
	tagOptions := &git.CreateTagOptions{
		Message: "Commit",
	}
	_, err := r.CreateTag("v1.0.0", commit.Hash, tagOptions)
	lib.CheckIfError(err)
	input := ProcessorInput{
		RepoPath:      dir,
		IncrementType: "minor",
		OfflineMode:   true,
	}
	version, err := input.ProcessSemver()
	lib.CheckIfError(err)
	expectedVersion := "1.1.0-rc.abc"
	if *version != expectedVersion {
		t.Error(fmt.Sprintf("Expected %s but received %s", *version, expectedVersion))
		t.Fail()
	}
}

func TestProcessSemverNonMainWithExistingTagMajor(t *testing.T) {
	t.Setenv("GITHUB_ACTIONS", "true")
	t.Setenv("GITHUB_REF_NAME", "feature/123")
	t.Setenv("GITHUB_RUN_ID", "abc")
	dir, r, _, commit := bvtest.SetupRepo("main", t)
	tagOptions := &git.CreateTagOptions{
		Message: "Commit",
	}
	_, err := r.CreateTag("v1.0.0", commit.Hash, tagOptions)
	lib.CheckIfError(err)
	input := ProcessorInput{
		RepoPath:      dir,
		IncrementType: "major",
		OfflineMode:   true,
	}
	version, err := input.ProcessSemver()
	lib.CheckIfError(err)
	expectedVersion := "2.0.0-rc.abc"
	if *version != expectedVersion {
		t.Error(fmt.Sprintf("Expected %s but received %s", *version, expectedVersion))
		t.Fail()
	}
}

func TestGetIncrementTypeEmptyOffline(t *testing.T) {
	input := ProcessorInput{
		OfflineMode:   true,
		IncrementType: "",
	}
	result, err := input.parseIncrementType()
	if err != nil {
		t.Fail()
	}
	if *result != "patch" {
		t.Fail()
	}
}

func TestGetIncrementTypePatchOffline(t *testing.T) {
	incrementTypes := []string{"Patch", "patch"}
	for _, incrementType := range incrementTypes {
		input := ProcessorInput{
			OfflineMode:   true,
			IncrementType: incrementType,
		}
		result, err := input.parseIncrementType()
		if err != nil {
			t.Fail()
		}
		if *result != "patch" {
			t.Fail()
		}
	}
}

func TestGetIncrementTypeMinorOffline(t *testing.T) {
	incrementTypes := []string{"Minor", "minor"}
	for _, incrementType := range incrementTypes {
		input := ProcessorInput{
			OfflineMode:   true,
			IncrementType: incrementType,
		}
		result, err := input.parseIncrementType()
		if err != nil {
			t.Fail()
		}
		if *result != "minor" {
			t.Fail()
		}
	}
}

func TestGetIncrementTypeMajorOffline(t *testing.T) {
	incrementTypes := []string{"Major", "major"}
	for _, incrementType := range incrementTypes {
		input := ProcessorInput{
			OfflineMode:   true,
			IncrementType: incrementType,
		}
		result, err := input.parseIncrementType()
		if err != nil {
			t.Fail()
		}
		if *result != "major" {
			t.Fail()
		}
	}
}

func TestGetIncrementTypeInvalidValueffline(t *testing.T) {
	incrementTypes := []string{"asdasda"}
	for _, incrementType := range incrementTypes {
		input := ProcessorInput{
			OfflineMode:   true,
			IncrementType: incrementType,
		}
		result, err := input.parseIncrementType()
		if result != nil && err == nil {
			t.Fail()
		}
	}
}
