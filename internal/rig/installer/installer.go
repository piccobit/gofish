package installer

import (
	"os"
	"path/filepath"

	"github.com/tinned-fish/gofish/internal/rig"
)

// Installer provides an interface for installing client rigs.
type Installer interface {
	// Install adds a rig to a path
	Install() error
	// Path is the directory of the installed rig.
	Path() string
	// Update updates a rig.
	Update() error
}

// Install installs a rig.
func Install(i Installer) error {
	basePath := filepath.Dir(i.Path())
	if _, pathErr := os.Stat(basePath); os.IsNotExist(pathErr) {
		if err := os.MkdirAll(basePath, 0755); err != nil {
			return err
		}
	}

	if _, pathErr := os.Stat(i.Path()); !os.IsNotExist(pathErr) {
		return i.Update()
	}

	return i.Install()
}

// Update updates a rig.
func Update(i Installer) error {
	if _, pathErr := os.Stat(i.Path()); os.IsNotExist(pathErr) {
		return rig.ErrDoesNotExist
	}

	return i.Update()
}

// FindSource determines the correct Installer for the given source.
func FindSource(location string) (Installer, error) {
	installer, err := existingVCSRepo(location)
	if err != nil && err.Error() == "Cannot detect VCS" {
		return installer, rig.ErrMissingSource
	}
	return installer, err
}

// New determines and returns the correct Installer for the given source
func New(source, name, version string) (Installer, error) {
	if isLocalReference(source) {
		return NewLocalInstaller(source, name)
	}

	return NewVCSInstaller(source, name, version)
}

// isLocalReference checks if the source exists on the filesystem.
func isLocalReference(source string) bool {
	_, err := os.Stat(source)
	return err == nil
}

// isRig checks if the directory contains a "Food" directory.
func isRig(dirname string) bool {
	_, err := os.Stat(filepath.Join(dirname, "Food"))
	return err == nil
}
