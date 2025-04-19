package generator

import (
	"strings"
	"testing"
)

func TestGeneratedVersionPatch(t *testing.T) {
	major := "1"
	minor := "0"
	patch := "0"
	version := &GeneratedVersion{
		Major: &major,
		Minor: &minor,
		Patch: &patch,
	}
	err := version.IncrementPatch()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if *version.Patch != "1" {
		t.Log(*version.Patch)
		t.Fail()
	}
}

func TestGeneratedVersionMinor(t *testing.T) {
	major := "1"
	minor := "2"
	patch := "3"
	version := &GeneratedVersion{
		Major: &major,
		Minor: &minor,
		Patch: &patch,
	}
	err := version.IncrementMinor()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if *version.Minor != "3" {
		t.Log(*version.Minor)
		t.Fail()
	}
}

func TestGeneratedVersionMajor(t *testing.T) {
	major := "1"
	minor := "2"
	patch := "3"
	version := &GeneratedVersion{
		Major: &major,
		Minor: &minor,
		Patch: &patch,
	}
	err := version.IncrementMajor()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if *version.Major != "2" {
		t.Log(*version.Major)
		t.Fail()
	}
}

func TestGetGeneratedVersion(t *testing.T) {
	expectedVersion := "1.0.0"
	version, err := GetGeneratedVersion(expectedVersion)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	versionSplits := strings.Split(expectedVersion, ".")
	if *version.Major != versionSplits[0] {
		t.Error("Major is not matching")
		t.Fail()
	}
	if *version.Minor != versionSplits[1] {
		t.Error("Minor is not matching")
		t.Fail()
	}
	if *version.Patch != versionSplits[2] {
		t.Error("Patch is not matching")
		t.Fail()
	}
}
