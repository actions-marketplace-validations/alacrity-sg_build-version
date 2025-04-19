package git

import (
	"testing"

	"github.com/alacrity-sg/build-version/src/lib"
	bv_test "github.com/alacrity-sg/build-version/test"
	"github.com/go-git/go-git/v5"
)

func TestGetLatestReleaseTagSingle(t *testing.T) {
	dir, r, _, commit := bv_test.SetupRepo("main", t)
	expectedTag := "v1.0.0"
	_, err := r.CreateTag(expectedTag, commit.Hash, &git.CreateTagOptions{
		Message: expectedTag,
	})
	lib.CheckIfError(err)
	tag, err := GetLatestReleaseTag(dir)
	lib.CheckIfError(err)
	if *tag != expectedTag[1:] {
		t.Fail()
	}
}

func TestGetLatestReleaseTagMultiple(t *testing.T) {
	dir, r, _, commit := bv_test.SetupRepo("main", t)
	unexpectedTag := "v1.0.0"
	_, err := r.CreateTag(unexpectedTag, commit.Hash, &git.CreateTagOptions{
		Message: unexpectedTag,
	})
	lib.CheckIfError(err)
	expectedTag := "v1.0.1"
	_, err = r.CreateTag(expectedTag, commit.Hash, &git.CreateTagOptions{
		Message: expectedTag,
	})
	lib.CheckIfError(err)
	tag, err := GetLatestReleaseTag(dir)
	lib.CheckIfError(err)
	if *tag != expectedTag[1:] {
		t.Fail()
	}
}

func TestGetLatestReleaseTagMultipleWithRC(t *testing.T) {
	dir, r, _, commit := bv_test.SetupRepo("main", t)
	tagOptions := &git.CreateTagOptions{
		Message: "Commit",
	}
	_, err := r.CreateTag("v1.0.0", commit.Hash, tagOptions)
	lib.CheckIfError(err)
	expectedTag := "v1.0.1"
	_, err = r.CreateTag(expectedTag, commit.Hash, tagOptions)
	lib.CheckIfError(err)
	_, err = r.CreateTag("v1.0.0-rc.1234", commit.Hash, tagOptions)
	lib.CheckIfError(err)
	tag, err := GetLatestReleaseTag(dir)
	lib.CheckIfError(err)
	if *tag != expectedTag[1:] {
		t.Fail()
	}
}

func TestGetRCTagSingle(t *testing.T) {
	dir, r, _, commit := bv_test.SetupRepo("main", t)
	tagOptions := &git.CreateTagOptions{
		Message: "Commit",
	}
	expectedTag := "v1.0.0-rc.12345as"
	_, err := r.CreateTag(expectedTag, commit.Hash, tagOptions)
	lib.CheckIfError(err)
	tag, err := GetLatestRCTag(dir)
	lib.CheckIfError(err)
	if *tag != expectedTag[1:] {
		t.Fail()
	}
}

func TestGetRCTagMultiple(t *testing.T) {
	dir, r, _, commit := bv_test.SetupRepo("main", t)
	tagOptions := &git.CreateTagOptions{
		Message: "Commit",
	}
	_, err := r.CreateTag("v1.0.0-rc.0000", commit.Hash, tagOptions)
	lib.CheckIfError(err)
	expectedTag := "v1.0.0-rc.12345as"
	_, err = r.CreateTag(expectedTag, commit.Hash, tagOptions)
	lib.CheckIfError(err)
	tag, err := GetLatestRCTag(dir)
	lib.CheckIfError(err)
	if *tag != expectedTag[1:] {
		t.Fail()
	}
}

func TestGetRCTagMultipleWithRelease(t *testing.T) {
	dir, r, _, commit := bv_test.SetupRepo("main", t)
	tagOptions := &git.CreateTagOptions{
		Message: "Commit",
	}
	_, err := r.CreateTag("v1.0.0-rc.0000", commit.Hash, tagOptions)
	lib.CheckIfError(err)
	expectedTag := "v1.0.0-rc.12345as"
	_, err = r.CreateTag(expectedTag, commit.Hash, tagOptions)
	lib.CheckIfError(err)
	_, err = r.CreateTag("v1.0.1", commit.Hash, tagOptions)
	tag, err := GetLatestRCTag(dir)
	lib.CheckIfError(err)
	if *tag != expectedTag[1:] {
		t.Fail()
	}
}
