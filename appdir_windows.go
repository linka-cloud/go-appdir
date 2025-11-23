package appdir

import (
	"os"
	"path/filepath"
	"sync"

	"golang.org/x/sys/windows"
)

type dirs struct {
	name string
}

var initOnce sync.Once
var localAppData string
var roamingAppData string
var programData string

func initFolders() {
	var err error
	localAppData, err = windows.KnownFolderPath(windows.FOLDERID_LocalAppData, 0)
	if err != nil {
		localAppData = os.Getenv("LOCALAPPDATA")
	}
	roamingAppData, err = windows.KnownFolderPath(windows.FOLDERID_RoamingAppData, 0)
	if err != nil {
		roamingAppData = os.Getenv("APPDATA")
	}
	programData, err = windows.KnownFolderPath(windows.FOLDERID_ProgramData, 0)
	if err != nil {
		programData = os.Getenv("PROGRAMDATA")
	}
}

func (d *dirs) UserConfig(uid ...string) string {
	if home := homeDirFor(uid); home != "" {
		return filepath.Join(home, "AppData", "Roaming", d.name)
	}
	initOnce.Do(initFolders)
	return filepath.Join(roamingAppData, d.name)
}

func (d *dirs) UserCache(uid ...string) string {
	if home := homeDirFor(uid); home != "" {
		return filepath.Join(home, "AppData", "Local", d.name)
	}
	initOnce.Do(initFolders)
	return filepath.Join(localAppData, d.name)
}

func (d *dirs) UserLogs(uid ...string) string {
	return d.UserCache(uid...)
}

func (d *dirs) UserData(uid ...string) string {
	return d.UserCache(uid...)
}

func (d *dirs) UserRun(uid ...string) string {
	return filepath.Join(d.UserCache(uid...), "Run")
}

func (d *dirs) SystemConfig() string {
	initOnce.Do(initFolders)
	return systemProgramPath(programData, d.name)
}

func (d *dirs) SystemData() string {
	initOnce.Do(initFolders)
	return systemProgramPath(programData, d.name)
}

func (d *dirs) SystemLogs() string {
	initOnce.Do(initFolders)
	return systemProgramPath(programData, d.name)
}

func (d *dirs) SystemRun() string {
	initOnce.Do(initFolders)
	return systemProgramPath(programData, d.name)
}

func systemProgramPath(base, name string) string {
	if base == "" {
		return ""
	}
	return filepath.Join(base, name)
}
