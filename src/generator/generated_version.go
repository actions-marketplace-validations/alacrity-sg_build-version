package generator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type GeneratedVersion struct {
	Major string
	Minor string
	Patch string
}

func GetGeneratedVersion(version string) (*GeneratedVersion, error) {
	firstSplit := strings.Split(version, "-")
	trimmedVersion := firstSplit[0]
	secondSplit := strings.Split(trimmedVersion, ".")
	if len(secondSplit) != 3 {
		return nil, errors.New("Version string does not fit semver structure")
	}
	return &GeneratedVersion{
		Major: secondSplit[0],
		Minor: secondSplit[1],
		Patch: secondSplit[2],
	}, nil
}

func (r *GeneratedVersion) IncrementPatch() error {
	intPatch, err := strconv.Atoi(r.Patch)
	if err != nil {
		return err
	}
	intPatch = intPatch + 1
	strPatch := strconv.Itoa(intPatch)
	r.Patch = strPatch
	return nil
}

func (r *GeneratedVersion) IncrementMinor() error {
	intMinor, err := strconv.Atoi(r.Minor)
	if err != nil {
		return err
	}
	intMinor = intMinor + 1
	strMinor := strconv.Itoa(intMinor)
	r.Minor = strMinor
	r.Patch = "0"
	return nil
}

func (r *GeneratedVersion) IncrementMajor() error {
	intMajor, err := strconv.Atoi(r.Major)
	if err != nil {
		return err
	}
	intMajor = intMajor + 1
	strMajor := strconv.Itoa(intMajor)
	r.Major = strMajor
	r.Minor = "0"
	r.Patch = "0"
	return nil
}

func (r *GeneratedVersion) BuildReleaseVersion() string {
	version := fmt.Sprintf("%s.%s.%s", r.Major, r.Minor, r.Patch)
	return version
}

func (r *GeneratedVersion) BuildReleaseCandidateVersion(uniqueId string) string {
	version := fmt.Sprintf("%s.%s.%s-rc.%s", r.Major, r.Minor, r.Patch, uniqueId)
	return version
}
