// Package appdir provides helpers to locate application directories such as config and cache.
package appdir

// Dirs exposes helpers to locate application directory paths.
type Dirs interface {
	// UserConfig returns the user-specific config directory.
	UserConfig() string
	// UserCache returns the user-specific cache directory.
	UserCache() string
	// UserLogs returns the user-specific logs directory.
	UserLogs() string
	// UserData returns the user-specific data directory.
	UserData() string
	// SystemConfig returns the system-wide config directory.
	SystemConfig() string
	// SystemData returns the system-wide data directory.
	SystemData() string
	// SystemLogs returns the system-wide logs directory.
	SystemLogs() string
}

// New creates a new Dirs implementation for the provided application name.
func New(name string) Dirs {
	return &dirs{name: name}
}
