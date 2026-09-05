package principal

import "charm.land/lipgloss/v2"

func panelTemplate() lipgloss.Style {
	p := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(0, 2).
		AlignVertical(lipgloss.Center).
		Height(30)
	return p
}

func PanelEnergy() lipgloss.Style {
	p := panelTemplate().
		Foreground(lipgloss.Color("#FFD93D")).
		BorderForeground(lipgloss.Color("#FFD93D"))
	return p
}

func PanelTemperature() lipgloss.Style {
	p := panelTemplate().
		Foreground(lipgloss.Color("#FFD93D")).
		BorderForeground(lipgloss.Color("#FFD93D"))
	return p
}
