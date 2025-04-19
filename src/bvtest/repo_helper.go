package bvtest

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func SetupRepo(branch string, t *testing.T) (string, *git.Repository, *git.Worktree, *object.Commit) {
	dir := t.TempDir()
	r, err := git.PlainInit(dir, false)
	err = r.CreateBranch(&config.Branch{Name: branch})
	w, err := r.Worktree()
	CheckIfError(err)
	fileName := filepath.Join(dir, "test.txt")
	err = os.WriteFile(fileName, []byte("test"), 0666)

	CheckIfError(err)
	_, err = w.Add("test.txt")
	CheckIfError(err)
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
