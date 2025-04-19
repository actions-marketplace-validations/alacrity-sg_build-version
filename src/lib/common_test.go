package lib

import "testing"

func TestGetIncrementTypeEmptyOffline(t *testing.T) {
	result, err := GetIncrementType("", false)
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
		result, err := GetIncrementType(incrementType, false)
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
		result, err := GetIncrementType(incrementType, false)
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
		result, err := GetIncrementType(incrementType, false)
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
		result, err := GetIncrementType(incrementType, false)
		if result != nil && err == nil {
			t.Fail()
		}
	}
}
