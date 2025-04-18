package processor

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/alacrity-sg/build-version/lib"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func setupRepo(branch string, t *testing.T) (string, *git.Repository, *git.Worktree, *object.Commit) {
	dir := t.TempDir()
	r, err := git.PlainInit(dir, false)
	err = r.CreateBranch(&config.Branch{Name: branch})
	w, err := r.Worktree()
	lib.CheckIfError(err)
	fileName := filepath.Join(dir, "test.txt")
	err = os.WriteFile(fileName, []byte("test"), 0666)

	lib.CheckIfError(err)
	_, err = w.Add("test.txt")
	lib.CheckIfError(err)
	commit, err := w.Commit("base", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Test",
			Email: "test@test.com",
			When:  time.Now(),
		},
	})
	obj, err := r.CommitObject(commit)
	return dir, r, w, obj
}

func TestProcessSemverMainWithTagDefaults(t *testing.T) {
	t.Setenv("GITHUB_CI", "true")
	t.Setenv("GITHUB_REF_NAME", "main")
	dir, r, _, commit := setupRepo("main", t)
	tagOptions := &git.CreateTagOptions{
		Message: "Commit",
	}
	_, err := r.CreateTag("v1.0.0", commit.Hash, tagOptions)
	lib.CheckIfError(err)
	_, err = r.CreateTag("v1.0.1-rc.123", commit.Hash, tagOptions)
	lib.CheckIfError(err)
	version, err := ProcessSemver(&ProcessorInput{
		RepoPath: &dir,
	})
	lib.CheckIfError(err)
	expectedVersion := "1.0.1"
	if *version != expectedVersion {
		t.Error(fmt.Sprintf("Expected %s but received %s", *version, expectedVersion))
		t.Fail()
	}
}

func TestProcessSemverNonMainWithDefaults(t *testing.T) {
	t.Setenv("GITHUB_CI", "true")
	t.Setenv("GITHUB_REF_NAME", "feature/123")
	t.Setenv("GITHUB_RUN_ID", "abc")
	dir, _, _, _ := setupRepo("main", t)
	version, err := ProcessSemver(&ProcessorInput{
		RepoPath: &dir,
	})
	lib.CheckIfError(err)
	expectedVersion := "0.0.1-rc.abc"
	if *version != expectedVersion {
		t.Error(fmt.Sprintf("Expected %s but received %s", *version, expectedVersion))
		t.Fail()
	}
}

func TestProcessSemverNonMainWithExistingTagDefaults(t *testing.T) {
	t.Setenv("GITHUB_CI", "true")
	t.Setenv("GITHUB_REF_NAME", "feature/123")
	t.Setenv("GITHUB_RUN_ID", "abc")
	dir, r, _, commit := setupRepo("main", t)
	tagOptions := &git.CreateTagOptions{
		Message: "Commit",
	}
	_, err := r.CreateTag("v1.0.0", commit.Hash, tagOptions)
	lib.CheckIfError(err)
	version, err := ProcessSemver(&ProcessorInput{
		RepoPath: &dir,
	})
	lib.CheckIfError(err)
	expectedVersion := "1.0.1-rc.abc"
	if *version != expectedVersion {
		t.Error(fmt.Sprintf("Expected %s but received %s", *version, expectedVersion))
		t.Fail()
	}
}
