package appdir

import (
	"os"
	"path/filepath"
)

type dirs struct {
	name string
}

func (d *dirs) UserConfig(uid ...string) string {
	if home := homeDirFor(uid); home != "" {
		return filepath.Join(home, "Library", "Application Support", d.name)
	}
	return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", d.name)
}

func (d *dirs) UserCache(uid ...string) string {
	if home := homeDirFor(uid); home != "" {
		return filepath.Join(home, "Library", "Caches", d.name)
	}
	return filepath.Join(os.Getenv("HOME"), "Library", "Caches", d.name)
}

func (d *dirs) UserLogs(uid ...string) string {
	if home := homeDirFor(uid); home != "" {
		return filepath.Join(home, "Library", "Logs", d.name)
	}
	return filepath.Join(os.Getenv("HOME"), "Library", "Logs", d.name)
}

func (d *dirs) UserData(uid ...string) string {
	if home := homeDirFor(uid); home != "" {
		return filepath.Join(home, "Library", "Application Support", d.name)
	}
	return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", d.name)
}

func (d *dirs) UserRun(uid ...string) string {
	if home := homeDirFor(uid); home != "" {
		return filepath.Join(home, "."+d.name, "run")
	}
	return filepath.Join(os.Getenv("HOME"), "."+d.name, "run")
}

func (d *dirs) SystemConfig() string {
	return filepath.Join("/Library", "Application Support", d.name)
}

func (d *dirs) SystemData() string {
	return filepath.Join("/Library", "Application Support", d.name)
}

func (d *dirs) SystemLogs() string {
	return filepath.Join("/Library", "Logs", d.name)
}

func (d *dirs) SystemRun() string {
	return filepath.Join("/var", "run", d.name)
}
