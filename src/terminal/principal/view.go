package principal

import (
	"nut/src/component/navbar"
	"strconv"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m Model) View() tea.View {

	navbar := navbar.TabView(0, m.MaxWidth)

	pEvents := sPanelEvents(m.MaxHeight, m.MaxWidth).Render(m.EventTerminal)

	pEnergy := sPanelEnergy(m.MaxHeight).Render(strconv.Itoa(m.Energy) + "%")
	pTemperature := sPanelTemperature(m.MaxHeight).Render(strconv.Itoa(m.Temperature) + "C")
	pOxygen := sPanelOxygen(m.MaxHeight).Render(strconv.Itoa(m.Oxygen) + "%")

	panels := lipgloss.JoinHorizontal(lipgloss.Top, pEvents, pOxygen, pTemperature, pEnergy)

	content := lipgloss.JoinVertical(lipgloss.Top, navbar, " ", panels)

	var v tea.View

	v.SetContent(content)
	v.AltScreen = true
	return v
}
