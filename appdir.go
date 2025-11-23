// Package appdir provides helpers to locate application directories such as config and cache.
package appdir

// Dirs exposes helpers to locate application directory paths.
type Dirs interface {
	// UserConfig returns the user-specific config directory. When uid is
	// provided, it is used to resolve another user's home directory.
	UserConfig(uid ...string) string
	// UserCache returns the user-specific cache directory. When uid is provided,
	// it is used to resolve another user's home directory.
	UserCache(uid ...string) string
	// UserLogs returns the user-specific logs directory. When uid is provided, it
	// is used to resolve another user's home directory.
	UserLogs(uid ...string) string
	// UserData returns the user-specific data directory. When uid is provided, it
	// is used to resolve another user's home directory.
	UserData(uid ...string) string
	// UserRun returns the user-specific runtime directory. When uid is provided,
	// it is used to resolve another user's runtime directory.
	UserRun(uid ...string) string
	// SystemConfig returns the system-wide config directory.
	SystemConfig() string
	// SystemData returns the system-wide data directory.
	SystemData() string
	// SystemLogs returns the system-wide logs directory.
	SystemLogs() string
	// SystemRun returns the system-wide runtime directory.
	SystemRun() string
}

// New creates a new Dirs implementation for the provided application name.
func New(name string) Dirs {
	return &dirs{name: name}
}
