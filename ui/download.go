package ui

import (
	"fmt"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type versionEntry structf {
	label string
	url   string
}

var ccVersions = []versionEntry{
	{"4.0.0 (Recommended)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_4_0_0_1680_capcutpc_0_creatortool.exe"},
	{"5.4.0 (Beta6)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_4_0_1991_beta6_capcutpc_beta_creatortool.exe"},
	{"5.4.0 (Beta5)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_4_0_1988_beta5_capcutpc_beta_creatortool.exe"},
	{"5.4.0 (Beta4)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_4_0_1982_beta4_capcutpc_beta_creatortool.exe"},
	{"5.4.0 (Beta3)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_4_0_1979_beta3_capcutpc_beta_creatortool.exe"},
	{"5.4.0 (Beta2)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_4_0_1978_beta2_capcutpc_beta_creatortool.exe"},
	{"5.4.0 (Beta1)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_4_0_1976_beta1_capcutpc_beta_creatortool.exe"},
	{"5.3.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_3_0_1964_capcutpc_0_creatortool.exe"},
	{"5.3.0 (Test2)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_3_0_1961_capcutpc_0_creatortool.exe"},
	{"5.3.0 (Test1)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_3_0_1957_capcutpc_0_creatortool.exe"},
	{"5.3.0 (Beta5)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_3_0_1962_beta5_capcutpc_beta_creatortool.exe"},
	{"5.3.0 (Beta4)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_3_0_1956_beta4_capcutpc_beta_creatortool.exe"},
	{"5.3.0 (Beta3)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_3_0_1952_beta3_capcutpc_beta_creatortool.exe"},
	{"5.3.0 (Beta2)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_3_0_1949_beta2_capcutpc_beta_creatortool.exe"},
	{"5.3.0 (Test1 Beta2)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_3_0_1947_beta2_capcutpc_beta_creatortool.exe"},
	{"5.3.0 (Beta1)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_3_0_1942_beta1_capcutpc_beta_creatortool.exe"},
	{"5.3.0 (Test1 Beta1)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_3_0_1941_beta1_capcutpc_beta_creatortool.exe"},
	{"5.2.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_2_0_1950_capcutpc_0_creatortool.exe"},
	{"5.2.0 (Test3)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_2_0_1946_capcutpc_0_creatortool.exe"},
	{"5.2.0 (Test2)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_2_0_1940_capcutpc_0_creatortool.exe"},
	{"5.2.0 (Test1)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_2_0_1939_capcutpc_0_creatortool.exe"},
	{"5.2.0 (Beta8)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_2_0_1945_beta8_capcutpc_beta_creatortool.exe"},
	{"5.2.0 (Beta7)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_2_0_1937_beta7_capcutpc_beta_creatortool.exe"},
	{"5.2.0 (Beta6)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_2_0_1934_beta6_capcutpc_beta_creatortool.exe"},
	{"5.2.0 (Beta5)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_2_0_1933_beta5_capcutpc_beta_creatortool.exe"},
	{"5.2.0 (Beta4)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_2_0_1929_beta4_capcutpc_beta_creatortool.exe"},
	{"5.2.0 (Beta3)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_2_0_1928_beta3_capcutpc_beta_creatortool.exe"},
	{"5.2.0 (Beta2)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_2_0_1925_beta2_capcutpc_beta_creatortool.exe"},
	{"5.2.0 (Beta1)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_2_0_1923_beta1_capcutpc_beta_creatortool.exe"},
	{"5.1.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_1_0_1926_capcutpc_0_creatortool.exe"},
	{"5.1.0 (Test2)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_1_0_1922_capcutpc_0_creatortool.exe"},
	{"5.1.0 (Test1)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_1_0_1919_capcutpc_0_creatortool.exe"},
	{"5.1.0 (Beta7)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_1_0_1924_beta7_capcutpc_beta_creatortool.exe"},
	{"5.1.0 (Beta6)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_1_0_1920_beta6_capcutpc_beta_creatortool.exe"},
	{"5.1.0 (Beta5)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_1_0_1918_beta5_capcutpc_beta_creatortool.exe"},
	{"5.1.0 (Beta4)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_1_0_1916_beta4_capcutpc_beta_creatortool.exe"},
	{"5.1.0 (Beta3)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_1_0_1913_beta3_capcutpc_beta_creatortool.exe"},
	{"5.1.0 (Beta2)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_1_0_1910_beta2_capcutpc_beta_creatortool.exe"},
	{"5.1.0 (Beta1)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_1_0_1907_beta1_capcutpc_beta_creatortool.exe"},
	{"5.0.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_0_0_1908_capcutpc_0_creatortool.exe"},
	{"5.0.0 (Latest v2)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_0_0_1903_capcutpc_0_creatortool.exe"},
	{"5.0.0 (Test1)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_0_0_1899_capcutpc_0_creatortool.exe"},
	{"5.0.0 (Beta6)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_0_0_1906_beta6_capcutpc_beta_creatortool.exe"},
	{"5.0.0 (Beta5)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_0_0_1905_beta5_capcutpc_beta_creatortool.exe"},
	{"5.0.0 (Beta4)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_0_0_1902_beta4_capcutpc_beta_creatortool.exe"},
	{"5.0.0 (Beta3)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_0_0_1901_beta3_capcutpc_beta_creatortool.exe"},
	{"5.0.0 (Beta2)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_0_0_1898_beta2_capcutpc_beta_creatortool.exe"},
	{"5.0.0 (Beta1)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_5_0_0_1897_beta1_capcutpc_beta_creatortool.exe"},
	{"4.7.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_4_7_0_1869_capcutpc_0_creatortool.exe"},
	{"4.6.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_4_6_0_1842_capcutpc_0_creatortool.exe"},
	{"4.5.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_4_5_0_1815_capcutpc_0_creatortool.exe"},
	{"4.4.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_4_4_0_1783_capcutpc_0_creatortool.exe"},
	{"4.3.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_4_3_0_1754_capcutpc_0_creatortool.exe"},
	{"4.2.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_4_2_0_1728_capcutpc_0_creatortool.exe"},
	{"4.1.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_4_1_0_1706_capcutpc_0_creatortool.exe"},
	{"3.9.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_3_9_0_1663_capcutpc_0_creatortool.exe"},
	{"3.8.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_3_8_0_1638_capcutpc_0_creatortool.exe"},
	{"3.7.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_3_7_0_1622_capcutpc_0_creatortool.exe"},
	{"3.6.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_3_6_0_1596_capcutpc_0_creatortool.exe"},
	{"3.5.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_3_5_0_1578_capcutpc_0_creatortool.exe"},
	{"3.4.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_3_4_0_1559_capcutpc_0_creatortool.exe"},
	{"3.3.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_3_3_0_1535_capcutpc_0_creatortool.exe"},
	{"3.2.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_3_2_0_1516_capcutpc_0_creatortool.exe"},
	{"3.1.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_3_1_0_1497_capcutpc_0_creatortool.exe"},
	{"3.0.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_3_0_0_1478_capcutpc_0_creatortool.exe"},
	{"2.9.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_2_9_0_1457_capcutpc_0_creatortool.exe"},
	{"2.8.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_2_8_0_1441_capcutpc_0_creatortool.exe"},
	{"2.7.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_2_7_0_1435_capcutpc_0_creatortool.exe"},
	{"2.6.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_2_6_0_1269_capcutpc_0.exe"},
	{"2.5.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_2_5_0_1222_capcutpc_0.exe"},
	{"2.4.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_2_4_0_1186_capcutpc_0.exe"},
	{"2.3.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_2_3_0_1158_capcutpc_0.exe"},
	{"2.2.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_2_2_0_1112_capcutpc_0.exe"},
	{"2.1.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_2_1_0_1038_capcutpc_0.exe"},
	{"2.0.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_2_0_0_822_capcutpc_0.exe"},
	{"1.9.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_1_9_0_699_capcutpc_0.exe"},
	{"1.8.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_1_8_0_633_capcutpc_0.exe"},
	{"1.7.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_1_7_0_587_capcutpc_0.exe"},
	{"1.6.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_1_6_0_519_capcutpc_0.exe"},
	{"1.5.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_1_5_0_433_capcutpc_0.exe"},
	{"1.4.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_1_4_0_334_capcutpc_0.exe"},
	{"1.3.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_1_3_0_262_capcutpc_0.exe"},
	{"1.2.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_1_2_0_213_capcutpc_0.exe"},
	{"1.1.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_1_1_0_155_capcutpc_0.exe"},
	{"1.0.0 (Latest)", "https://lf16-capcut.faceulv.com/obj/capcutpc-packages-us/packages/CapCut_1_0_0_44_capcutpc_0.exe"},
}

type downloadModel struct {
	appsDir    string
	cursor     int
	start      int
	showDetail bool
	msg        string
}

func NewDownload(appsDir string) tea.Model {
	return downloadModel{appsDir: appsDir}
}

func (m downloadModel) Init() tea.Cmd {
	return nil
}

const listHeight = 18

func (m downloadModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.msg != "" {
			if msg.String() == "esc" {
				return NewMenu(m.appsDir), nil
			}
			return m, nil
		}
		if m.showDetail {
			switch msg.String() {
			case "1":
				go openBrowser(ccVersions[m.cursor].url)
				m.msg = fmt.Sprintf("Opening in browser...\n\n  If it didn't open, copy this URL:\n\n  %s", ccVersions[m.cursor].url)
				return m, nil
			case "2":
				m.msg = fmt.Sprintf("URL:\n\n  %s", ccVersions[m.cursor].url)
				return m, nil
			case "esc":
				m.showDetail = false
				m.msg = ""
				return m, nil
			}
		}
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
				if m.cursor < m.start {
					m.start = m.cursor
				}
			}
		case "down", "j":
			if m.cursor < len(ccVersions)-1 {
				m.cursor++
				if m.cursor >= m.start+listHeight {
					m.start = m.cursor - listHeight + 1
				}
			}
		case "pgup":
			m.cursor -= listHeight
			if m.cursor < 0 {
				m.cursor = 0
			}
			m.start = m.cursor
		case "pgdown":
			m.cursor += listHeight
			if m.cursor >= len(ccVersions) {
				m.cursor = len(ccVersions) - 1
			}
			if m.cursor >= m.start+listHeight {
				m.start = m.cursor - listHeight + 1
			}
		case "enter":
			m.showDetail = true
			m.msg = ""
		case "esc":
			return NewMenu(m.appsDir), nil
		}
	}
	return m, nil
}

func (m downloadModel) View() string {
	if m.msg != "" {
		return "\n  DOWNLOAD SUPPORTED VERSION\n\n" + m.msg + "\n\n  Press esc to go back...\n"
	}

	if m.showDetail {
		s := "\n  DOWNLOAD SUPPORTED VERSION\n\n"
		s += fmt.Sprintf("  Version: %s\n\n", ccVersions[m.cursor].label)
		s += "  [1] Open in browser\n"
		s += "  [2] Show URL only\n"
		s += "  [esc] Back\n"
		return s
	}

	s := "\n  DOWNLOAD SUPPORTED VERSION\n\n"
	end := m.start + listHeight
	if end > len(ccVersions) {
		end = len(ccVersions)
	}

	for i := m.start; i < end; i++ {
		cursor := "  "
		if i == m.cursor {
			cursor = ">>"
		}
		s += fmt.Sprintf("  %s %s\n", cursor, ccVersions[i].label)
	}

	s += fmt.Sprintf("\n  (%d/%d)", m.cursor+1, len(ccVersions))
	s += "\n\n  [enter] Select  [esc] Back\n\n"
	s += "  How to downgrade:\n"
	s += "  Goto \"C:\\Users\\Your PC Name\\AppData\\Local\\CapCut\\Apps\"\n"
	s += "  And delete all of the folders in there.\n"
	return s
}

func openBrowser(url string) {
	cmd := exec.Command("cmd", "/c", "start", strings.ReplaceAll(url, "&", "^&"))
	cmd.Run()
}
