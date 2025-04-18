package lib

import "testing"

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
	version, err := GetGeneratedVersion("1.0.0")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if *version.Major != "1" {
		t.Log("Major is not 1")
		t.Fail()
	}
	if *version.Minor != "0" {
		t.Log("Minor is not 0")
		t.Fail()
	}
	if *version.Patch != "0" {
		t.Log("Patch is not 0")
		t.Fail()
	}
}
