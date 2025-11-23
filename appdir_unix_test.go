//go:build unix && !darwin && !windows
// +build unix,!darwin,!windows

package appdir

import (
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"testing"
)

func setHome(t *testing.T, path string) func() {
	t.Helper()
	orig := os.Getenv("HOME")
	if err := os.Setenv("HOME", path); err != nil {
		t.Fatalf("setenv HOME: %v", err)
	}
	return func() {
		os.Setenv("HOME", orig)
	}
}

func TestUserConfig(t *testing.T) {
	restore := setHome(t, "/home/user")
	defer restore()
	d := dirs{name: "app"}
	expected := "/home/user/.config/app"
	if have := d.UserConfig(); have != expected {
		t.Fatalf("expected %v, found %v", expected, have)
	}
}

func TestUserCache(t *testing.T) {
	restore := setHome(t, "/home/user")
	defer restore()
	d := dirs{name: "app"}
	expected := "/home/user/.cache/app"
	if have := d.UserCache(); have != expected {
		t.Fatalf("expected %v, found %v", expected, have)
	}
}

func TestUserLogs(t *testing.T) {
	restore := setHome(t, "/home/user")
	defer restore()
	d := dirs{name: "app"}
	expected := "/home/user/.cache/app/logs"
	if have := d.UserLogs(); have != expected {
		t.Fatalf("expected %v, found %v", expected, have)
	}
}

func TestUserData(t *testing.T) {
	restore := setHome(t, "/home/user")
	defer restore()
	d := dirs{name: "app"}
	expected := "/home/user/.local/share/app"
	if have := d.UserData(); have != expected {
		t.Fatalf("expected %v, found %v", expected, have)
	}
}

func TestUserRun(t *testing.T) {
	d := dirs{name: "app"}
	expected := filepath.Join("/run", "user", strconv.Itoa(os.Getuid()), "app")
	if have := d.UserRun(); have != expected {
		t.Fatalf("expected %v, found %v", expected, have)
	}
}

func TestUserConfigWithUID(t *testing.T) {
	current, err := user.Current()
	if err != nil {
		t.Skipf("user.Current failed: %v", err)
	}
	d := dirs{name: "app"}
	expected := filepath.Join(current.HomeDir, ".config", "app")
	if have := d.UserConfig(current.Uid); have != expected {
		t.Fatalf("expected %v, found %v", expected, have)
	}
}

func TestUserRunWithUID(t *testing.T) {
	current, err := user.Current()
	if err != nil {
		t.Skipf("user.Current failed: %v", err)
	}
	d := dirs{name: "app"}
	expected := filepath.Join("/run", "user", current.Uid, "app")
	if have := d.UserRun(current.Uid); have != expected {
		t.Fatalf("expected %v, found %v", expected, have)
	}
}

func TestSystemConfig(t *testing.T) {
	d := dirs{name: "app"}
	if have := d.SystemConfig(); have != "/etc/app" {
		t.Fatalf("unexpected path %v", have)
	}
}

func TestSystemData(t *testing.T) {
	d := dirs{name: "app"}
	if have := d.SystemData(); have != "/var/lib/app" {
		t.Fatalf("unexpected path %v", have)
	}
}

func TestSystemLogs(t *testing.T) {
	d := dirs{name: "app"}
	if have := d.SystemLogs(); have != "/var/log/app" {
		t.Fatalf("unexpected path %v", have)
	}
}

func TestSystemRun(t *testing.T) {
	d := dirs{name: "app"}
	if have := d.SystemRun(); have != "/var/run/app" {
		t.Fatalf("unexpected path %v", have)
	}
}
