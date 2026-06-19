package util

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func IsCapCutRunning() bool {
	out, err := exec.Command("tasklist", "/FI", "IMAGENAME eq CapCut.exe").Output()
	if err != nil {
		return false
	}
	return strings.Contains(string(out), "CapCut.exe")
}

func KillCapCut() {
	exec.Command("taskkill", "/F", "/IM", "CapCut.exe").Run()
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		if !IsCapCutRunning() {
			return
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func ResolveAppsDir() (string, error) {
	localAppData := os.Getenv("LOCALAPPDATA")
	if localAppData == "" {
		return "", os.ErrNotExist
	}
	dir := filepath.Join(localAppData, "CapCut", "Apps")
	if _, err := os.Stat(dir); err == nil {
		return dir, nil
	}
	return "", os.ErrNotExist
}

func ListVersions(appsDir string) []string {
	entries, err := os.ReadDir(appsDir)
	if err != nil {
		return nil
	}
	var versions []string
	for _, e := range entries {
		if e.IsDir() && isVersionDir(e.Name()) {
			versions = append(versions, e.Name())
		}
	}
	return versions
}

func isVersionDir(name string) bool {
	if len(name) == 0 {
		return false
	}
	for _, c := range name {
		if c != '.' && (c < '0' || c > '9') {
			return false
		}
	}
	return name != "." && name != ".."
}
