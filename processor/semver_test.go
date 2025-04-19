package processor

import (
	"fmt"
	"testing"

	"github.com/alacrity-sg/build-version/lib"
	bv_test "github.com/alacrity-sg/build-version/test"
	"github.com/go-git/go-git/v5"
)

func TestProcessSemverMainWithTagDefaults(t *testing.T) {
	t.Setenv("GITHUB_ACTIONS", "true")
	t.Setenv("GITHUB_REF_NAME", "main")
	dir, r, _, commit := bv_test.SetupRepo("main", t)
	tagOptions := &git.CreateTagOptions{
		Message: "Commit",
	}
	_, err := r.CreateTag("v1.0.0", commit.Hash, tagOptions)
	lib.CheckIfError(err)
	_, err = r.CreateTag("v1.0.1-rc.123", commit.Hash, tagOptions)
	lib.CheckIfError(err)
	version, err := ProcessSemver(&ProcessorInput{
		RepoPath: dir,
	})
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
	dir, _, _, _ := bv_test.SetupRepo("main", t)
	version, err := ProcessSemver(&ProcessorInput{
		RepoPath: dir,
	})
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
	dir, r, _, commit := bv_test.SetupRepo("main", t)
	tagOptions := &git.CreateTagOptions{
		Message: "Commit",
	}
	_, err := r.CreateTag("v1.0.0", commit.Hash, tagOptions)
	lib.CheckIfError(err)
	version, err := ProcessSemver(&ProcessorInput{
		RepoPath: dir,
	})
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
	dir, r, _, commit := bv_test.SetupRepo("main", t)
	tagOptions := &git.CreateTagOptions{
		Message: "Commit",
	}
	_, err := r.CreateTag("v1.0.0", commit.Hash, tagOptions)
	lib.CheckIfError(err)
	version, err := ProcessSemver(&ProcessorInput{
		RepoPath:      dir,
		IncrementType: "patch",
		OfflineMode:   false,
	})
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
	dir, r, _, commit := bv_test.SetupRepo("main", t)
	tagOptions := &git.CreateTagOptions{
		Message: "Commit",
	}
	_, err := r.CreateTag("v1.0.0", commit.Hash, tagOptions)
	lib.CheckIfError(err)
	version, err := ProcessSemver(&ProcessorInput{
		RepoPath:      dir,
		IncrementType: "minor",
		OfflineMode:   false,
	})
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
	dir, r, _, commit := bv_test.SetupRepo("main", t)
	tagOptions := &git.CreateTagOptions{
		Message: "Commit",
	}
	_, err := r.CreateTag("v1.0.0", commit.Hash, tagOptions)
	lib.CheckIfError(err)
	version, err := ProcessSemver(&ProcessorInput{
		RepoPath:      dir,
		IncrementType: "major",
		OfflineMode:   false,
	})
	lib.CheckIfError(err)
	expectedVersion := "2.0.0-rc.abc"
	if *version != expectedVersion {
		t.Error(fmt.Sprintf("Expected %s but received %s", *version, expectedVersion))
		t.Fail()
	}
}
