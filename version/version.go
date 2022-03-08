package version

import (
	"fmt"
)

// version is the current version of GoFish.
var version = "dev"

// date is the extra build time data.
var date = ""

// commit is the Git commit SHA.
var commit = ""

// buildBy is the information about the used builder.
var buildBy = "goreleaser"

// String represents the version information as a well-formatted string.
func String() string {
    ver := ""

    if len(version) != 0 {
        ver += fmt.Sprintf("Version: %s\n", version)
    }

    if len(date) != 0 {
        ver += fmt.Sprintf("Date: %s\n", date)
    }

    if len(commit) != 0 {
        ver += fmt.Sprintf("Commit: %s\n", commit)
    }

    if len(buildBy) != 0 {
        ver += fmt.Sprintf("Build by: %s\n", buildBy)
    }

	return ver
}
