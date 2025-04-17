package lib

import "github.com/go-git/go-git/v5"

func GetLastCommit(repoPath string) (*string, error) {
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, err
	}
	ref, err := r.Head()

	if err != nil {
		return nil, err
	}

	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})

	if err != nil {
		return nil, err
	}
	commit, err := cIter.Next()
	hashString := commit.Hash.String()
	return &hashString, nil
}
