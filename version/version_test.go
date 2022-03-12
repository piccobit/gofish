package version

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	var expected string

	expected = "Version: dev\nBuild by: goreleaser\n"
	if String() != expected {
		t.Errorf("Expected .String() to return '%s', got '%s'", expected, String())
	}

	version = "0.16.3"
	expected = fmt.Sprintf("Version: %s\nBuild by: goreleaser\n", version)
	if String() != "Version: 0.1.0\nBuild by: goreleaser\n" {
		t.Errorf("Expected .String() to return '%s', got '%s'", expected, String())
	}

	date = "2022-03-11T08:08:58Z"
	expected = fmt.Sprintf("Version: %s\nDate: %s\nBuild by: goreleaser\n", version, date)
	if String() != expected {
		t.Errorf("Expected .String() to return '%s', got '%s'", expected, String())
	}

	commit = "e4c6c4bb686874ce32908159d27b2402ba558950"
	expected = fmt.Sprintf("Version: %s\nDate: %s\nCommit: %s\nBuild by: goreleaser\n", version, date, commit)
	if String() != expected {
		t.Errorf("Expected .String() to return '%s', got '%s'", expected, String())
	}

	buildBy = "HDS"
	expected = fmt.Sprintf("Version: %s\nDate: %s\nCommit: %s\nBuild by: %s\n", version, date, commit, buildBy)
	if String() != expected {
		t.Errorf("Expected .String() to return '%s', got '%s'", expected, String())
	}
}
