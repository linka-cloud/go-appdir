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

func (d *dirs) UserConfig(uid ...string) string {
	base := unixHome(uid...)
	if base == "" {
		return filepath.Join(os.TempDir(), d.name, "config")
	}
	return filepath.Join(base, ".config", d.name)
}

func (d *dirs) UserCache(uid ...string) string {
	base := unixHome(uid...)
	if base == "" {
		return filepath.Join(os.TempDir(), d.name, "cache")
	}
	return filepath.Join(base, ".cache", d.name)
}

func (d *dirs) UserData(uid ...string) string {
	base := unixHome(uid...)
	if base == "" {
		return filepath.Join(os.TempDir(), d.name, "data")
	}
	return filepath.Join(base, ".local", "share", d.name)
}

func (d *dirs) UserLogs(uid ...string) string {
	return filepath.Join(d.UserCache(uid...), "logs")
}

func (d *dirs) UserRun(uid ...string) string {
	if uidVal, ok := providedUID(uid); ok {
		return filepath.Join("/run", "user", uidVal, d.name)
	}
	current := os.Getuid()
	if current >= 0 {
		return filepath.Join("/run", "user", strconv.Itoa(current), d.name)
	}
	return filepath.Join(os.TempDir(), d.name)
}

func unixHome(uid ...string) string {
	if home := homeDirFor(uid); home != "" {
		return home
	}
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if home, err := os.UserHomeDir(); err == nil {
		return home
	}
	return ""
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
