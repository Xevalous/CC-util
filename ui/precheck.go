package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"cc-util/util"

	tea "github.com/charmbracelet/bubbletea"
)

type checkResult struct {
	name   string
	status string // "OK", "WARN", "FAIL"
}

type precheckModel struct {
	checks  []checkResult
	failed  bool
	ready   bool
	appsDir string
}

type tickMsg time.Time

func runPrechecks() []checkResult {
	var checks []checkResult

	// OS check
	if runtime.GOOS == "windows" {
		checks = append(checks, checkResult{"OS: Windows", "OK"})
	} else {
		checks = append(checks, checkResult{"OS: " + runtime.GOOS, "FAIL"})
	}

	// CapCut installed
	localAppData := os.Getenv("LOCALAPPDATA")
	capcutPath := filepath.Join(localAppData, "CapCut")
	if _, err := os.Stat(capcutPath); err == nil {
		checks = append(checks, checkResult{"CapCut installed", "OK"})
	} else {
		checks = append(checks, checkResult{"CapCut installed", "FAIL"})
	}

	// CapCut running
	if util.IsCapCutRunning() {
		checks = append(checks, checkResult{"CapCut is running", "WARN"})
	} else {
		checks = append(checks, checkResult{"CapCut is not running", "OK"})
	}

	return checks
}

func NewPrecheck() tea.Model {
	checks := runPrechecks()
	failed := false
	for _, c := range checks {
		if c.status == "FAIL" {
			failed = true
			break
		}
	}

	appsDir, _ := util.ResolveAppsDir()

	return precheckModel{checks: checks, failed: failed, appsDir: appsDir}
}

func (m precheckModel) Init() tea.Cmd {
	if m.failed {
		return nil
	}
	return tea.Tick(2*time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m precheckModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		if m.failed {
			return m, tea.Quit
		}
	case tickMsg:
		if !m.failed {
			return NewMenu(m.appsDir), nil
		}
	}
	return m, nil
}

func (m precheckModel) View() string {
	s := "\n  CC UTIL\n\n  Checking requirements...\n\n"

	for _, c := range m.checks {
		var tag string
		switch c.status {
		case "OK":
			tag = "[OK]"
		case "WARN":
			tag = "[WARN]"
		case "FAIL":
			tag = "[FAIL]"
		}
		s += fmt.Sprintf("  %s  %s\n", tag, c.name)
	}

	s += "\n"

	if m.failed {
		s += "  Press any key to exit...\n"
	}

	return s
}


