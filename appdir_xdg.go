//go:build !darwin && !windows
// +build !darwin,!windows

package appdir

import (
	"os"
	"path/filepath"
	"strconv"
)

type dirs struct {
	name string
}

func (d *dirs) UserConfig() string {
	baseDir := os.Getenv("XDG_CONFIG_HOME")
	if baseDir == "" {
		baseDir = filepath.Join(os.Getenv("HOME"), ".config")
	}

	return filepath.Join(baseDir, d.name)
}

func (d *dirs) UserCache() string {
	baseDir := os.Getenv("XDG_CACHE_HOME")
	if baseDir == "" {
		baseDir = filepath.Join(os.Getenv("HOME"), ".cache")
	}

	return filepath.Join(baseDir, d.name)
}

func (d *dirs) UserData() string {
	baseDir := os.Getenv("XDG_DATA_HOME")
	if baseDir == "" {
		baseDir = filepath.Join(os.Getenv("HOME"), ".local", "share")
	}

	return filepath.Join(baseDir, d.name)
}

func (d *dirs) UserLogs() string {
	return filepath.Join(d.UserCache(), "logs")
}

func (d *dirs) UserRun() string {
	if baseDir := os.Getenv("XDG_RUNTIME_DIR"); baseDir != "" {
		return filepath.Join(baseDir, d.name)
	}
	uid := os.Getuid()
	if uid >= 0 {
		return filepath.Join("/run", "user", strconv.Itoa(uid), d.name)
	}
	return filepath.Join(os.TempDir(), d.name)
}

func (d *dirs) SystemConfig() string {
	return filepath.Join("/etc", d.name)
}

func (d *dirs) SystemData() string {
	return filepath.Join("/var/lib", d.name)
}

func (d *dirs) SystemLogs() string {
	return filepath.Join("/var/log", d.name)
}

func (d *dirs) SystemRun() string {
	return filepath.Join("/var/run", d.name)
}
