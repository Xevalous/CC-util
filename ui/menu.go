package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type menuModel struct {
	cursor  int
	appsDir string
}

func NewMenu(appsDir string) tea.Model {
	return menuModel{appsDir: appsDir}
}

func (m menuModel) Init() tea.Cmd {
	return nil
}

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "1":
			return NewPatch(m.appsDir), nil
		case "2":
			return NewLock(m.appsDir), nil
		case "3":
			return NewDownload(m.appsDir), nil
		case "4", "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m menuModel) View() string {
	s := "\n          CC UTIL\n\n"
	s += "  [1] Patch\n"
	s += "  [2] Lock Version\n"
	s += "  [3] Download Supported Version\n"
	s += "  [4] Exit\n\n"
	return s
}
