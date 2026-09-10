package navbar

import "charm.land/lipgloss/v2"

func navElementTemplate() lipgloss.Style {
	n := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		Padding(1, 2).
		Width(15).
		Align(lipgloss.Center)
	return n
}

func sNavBar(widthMax int) lipgloss.Style {
	n := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(widthMax).
		Height(5).
		Padding(0).
		Margin(0)
	return n
}

func sActiveTab() lipgloss.Style {
	n := navElementTemplate().
		Background(lipgloss.Color("#FFFFFF")).
		Foreground(lipgloss.Black)
	return n
}

func sInactiveTab() lipgloss.Style {
	n := navElementTemplate()
	return n
}
