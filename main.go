package main



import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (

  tableStyle = lipgloss.NewStyle().
    BorderStyle(lipgloss.NormalBorder()).
    BorderForeground(lipgloss.Color("240"))

  sidebarStyle = lipgloss.NewStyle().
    BorderStyle(lipgloss.NormalBorder()).
    BorderForeground(lipgloss.Color("205")).
    Width(20) // Set a fixed width for the sidebar
    
)

type model struct {
	table table.Model
  sidebar string
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("%s selected!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
    // Render the table and sidebar individually with their styles
    tableRendered := tableStyle.Render(m.table.View())
    sidebarRendered := sidebarStyle.Render(m.sidebar)

    // Use lipgloss JoinHorizontal to properly place blocks side by side
    // Adjusting to 'lipgloss.Top' aligns both components at the top
    fullView := lipgloss.JoinHorizontal(lipgloss.Top, tableRendered, "  ", sidebarRendered)

    // Return the combined view
    return fullView + "\n"
}
/*
func (m model) View() string {
    // Define the minimum width required for the table and sidebar.
    tableWidth := 80
    sidebarWidth := 20
    fullTable := m.table.View()

    // Calculate the remaining width after the table's actual width is considered.
    remainingWidth := tableWidth - lipgloss.Width(fullTable)
    if remainingWidth < 0 {
        remainingWidth = 0
    }

    // Build the full view string with the table and sidebar.
    return baseStyle.Render(
        fullTable + strings.Repeat(" ", remainingWidth) + // Ensure no negative repeat count.
        lipgloss.NewStyle().Width(sidebarWidth).Render(m.sidebar),
    ) + "\n"
}
  */

func main() { columns := []table.Column{
		{Title: "MAC Addr", Width: 20},
		{Title: "Miner Type", Width: 10},
		{Title: "Hashrate", Width: 10},
		{Title: "Status", Width: 6},
		{Title: "Mode", Width: 4},
		{Title: "Uptime", Width: 6},
		{Title: "Fleet Name", Width: 10},
	}

  rows := []table.Row{
      {"00:1A:2B:3C:4D:5E", "ASIC-S9", "13.5 TH", "Active", "Auto", "72h", "AlphaOne"},
      {"01:23:45:67:89:AB", "ASIC-S9", "14.0 TH", "Down", "Man", "15h", "BetaMax"},
      {"1A:2B:3C:4D:5E:6F", "GPU-580", "0.3 TH", "Active", "Auto", "120h", "GammaField"},
      {"00:11:22:33:44:55", "ASIC-S17", "56.0 TH", "Active", "Auto", "96h", "DeltaCrew"},
      {"AA:BB:CC:DD:EE:FF", "ASIC-S17", "53.0 TH", "Active", "Auto", "45h", "Epsilon"},
      {"11:22:33:44:55:66", "GPU-2080Ti", "0.8 TH", "Down", "Man", "22h", "ZetaWave"},
      {"6C:5A:B5:FF:AD:BC", "ASIC-S9", "13.0 TH", "Active", "Auto", "84h", "EtaStream"},
      {"DE:AD:BE:EF:00:01", "ASIC-S17", "55.5 TH", "Active", "Auto", "65h", "ThetaLine"},
      {"99:88:77:66:55:44", "GPU-580", "0.29 TH", "Down", "Man", "18h", "IotaTrack"},
      {"4F:5A:6C:7D:8E:9F", "ASIC-S17", "57.0 TH", "Active", "Auto", "110h", "KappaPlane"},
  }

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{
    table: t,
    sidebar: "Upon my head they've placed a fruitless crown and a barren sceptre in my grip.",
  }

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
