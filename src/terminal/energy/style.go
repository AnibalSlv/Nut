package energy

import "charm.land/lipgloss/v2"

var maxHeight int = 4

func panelTemplate() lipgloss.Style {
	p := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 3).
		AlignVertical(lipgloss.Center).
		Height(maxHeight).
		Width(50)
	return p
}

func sPanelEngine() lipgloss.Style {
	p := panelTemplate().
		Foreground(lipgloss.Color("#FFD93D")).
		BorderForeground(lipgloss.Color("#FFD93D"))
	return p
}
