package principal

import (
	"nut/src/component/navbar"
	"strconv"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m Model) View() tea.View {

	navbar := navbar.TabView(m.ActiveTab, m.maxWidth)

	pEvents := sPanelEvents(m.maxHeight, m.maxWidth).Render(m.EventTerminal)

	pEnergy := sPanelEnergy(m.maxHeight).Render(strconv.Itoa(m.Energy) + "%")
	pTemperature := sPanelTemperature(m.maxHeight).Render(strconv.Itoa(m.Temperature) + "C")
	pOxygen := sPanelOxygen(m.maxHeight).Render(strconv.Itoa(m.Oxygen) + "%")

	panels := lipgloss.JoinHorizontal(lipgloss.Top, pEvents, pOxygen, pTemperature, pEnergy)

	content := lipgloss.JoinVertical(lipgloss.Top, navbar, " ", panels)

	var v tea.View

	v.SetContent(content)
	v.AltScreen = true
	return v
}
