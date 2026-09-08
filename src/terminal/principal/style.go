package principal

import "charm.land/lipgloss/v2"

const navbarHeight int = 6

func panelTemplate(maxHeight int) lipgloss.Style {
	p := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(0, 3).
		AlignVertical(lipgloss.Center).
		Height(maxHeight - navbarHeight)
	return p
}

func sPanelEvents(maxHeight int, maxWidth int) lipgloss.Style {
	p := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#4AC109")).
		Foreground(lipgloss.Color("#4AC109")).
		Padding(1).
		Height(maxHeight - navbarHeight).
		Width(maxWidth - 30)
	return p
}

func sPanelEnergy(maxHeight int) lipgloss.Style {
	p := panelTemplate(maxHeight).
		Foreground(lipgloss.Color("#FFD93D")).
		BorderForeground(lipgloss.Color("#FFD93D"))
	return p
}

func sPanelTemperature(maxHeight int) lipgloss.Style {
	p := panelTemplate(maxHeight).
		Foreground(lipgloss.Color("#FF5200")).
		BorderForeground(lipgloss.Color("#FF5200"))
	return p
}

func sPanelOxygen(maxHeight int) lipgloss.Style {
	p := panelTemplate(maxHeight).
		Foreground(lipgloss.Color("#0094FF")).
		BorderForeground(lipgloss.Color("#0094FF"))
	return p
}
